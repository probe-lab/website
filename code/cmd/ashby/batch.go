package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/urfave/cli/v2"
	"golang.org/x/exp/slog"
	"golang.org/x/sync/errgroup"
	"gopkg.in/yaml.v3"
)

var batchCommand = &cli.Command{
	Name:   "batch",
	Usage:  "Batch command to generate a group of plots",
	Action: Batch,
	Flags: append([]cli.Flag{
		&cli.BoolFlag{
			Name:        "compact",
			Required:    false,
			Usage:       "Emit compact json instead of pretty-printed.",
			Destination: &batchOpts.compact,
		},
		&cli.BoolFlag{
			Name:        "validate",
			Required:    false,
			Usage:       "Validate the input files without running queries.",
			Destination: &batchOpts.validate,
		},
		&cli.StringSliceFlag{
			Name:        "source",
			Aliases:     []string{"s"},
			Required:    false,
			Usage:       "Specify the url of a data source, in the format name=url. May be repeated to specify multiple sources. Postgres urls take the form 'postgres://username:password@hostname:5432/database_name'",
			Destination: &batchOpts.sources,
		},
		&cli.StringFlag{
			Name:        "in",
			Required:    true,
			Usage:       "Path of directory containing plot definitions.",
			Destination: &batchOpts.inDir,
		},
		&cli.StringFlag{
			Name:        "out",
			Required:    true,
			Usage:       "Path of directory where plots should be written.",
			Destination: &batchOpts.outDir,
		},
		&cli.StringFlag{
			Name:        "basis",
			Required:    false,
			Value:       "now",
			Usage:       "Basis time that should be passed to queries. Specify 'now' or a valid date in the past in RFC3339 or Unix timestamp format.",
			Destination: &batchOpts.basis,
		},
		&cli.BoolFlag{
			Name:        "version",
			Required:    false,
			Usage:       "Automatically version the plots by appending a relevant date suffix and write most recent plot with '-latest' suffix.",
			Destination: &batchOpts.version,
		},
		&cli.BoolFlag{
			Name:        "force",
			Required:    false,
			Usage:       "Force generation of plots even if the plot output already exists.",
			Destination: &batchOpts.force,
		},
		&cli.IntFlag{
			Name:        "concurrency",
			Required:    false,
			Usage:       "Number of concurrent goroutines to use for generating plots.",
			Destination: &batchOpts.concurrency,
			Value:       6,
		},
		&cli.StringFlag{
			Name:        "conf",
			Required:    false,
			Usage:       "Path of directory containing configuration.",
			Destination: &batchOpts.confDir,
		},
	}, loggingFlags...),
}

var batchOpts struct {
	preview     bool
	compact     bool
	sources     cli.StringSlice
	inDir       string
	outDir      string
	confDir     string
	validate    bool
	version     bool
	force       bool
	basis       string
	concurrency int
}

func Batch(cc *cli.Context) error {
	ctx := cc.Context
	setupLogging()

	if batchOpts.validate {
		// avoid interlacing output
		batchOpts.concurrency = 1
	}

	cfg := &PlotConfig{
		Sources: map[string]DataSource{
			"static": &StaticDataSource{},
			"demo":   &DemoDataSource{},
		},
		Colors: map[string]string{},
	}

	if batchOpts.basis == "now" {
		cfg.BasisTime = time.Now()
	} else {
		var err error
		ts, err := strconv.Atoi(batchOpts.basis)
		if err != nil {
			cfg.BasisTime, err = time.Parse(time.RFC3339, batchOpts.basis)
			if err != nil {
				return fmt.Errorf("invalid basis time: %w", err)
			}
		} else {
			cfg.BasisTime = time.Unix(int64(ts), 0)
		}

		if cfg.BasisTime.After(time.Now()) {
			return fmt.Errorf("basis time should not be in the future: %s", cfg.BasisTime.Format(time.RFC3339))
		}
	}
	cfg.BasisTime = cfg.BasisTime.UTC()
	slog.Info("plots will be generated for time " + cfg.BasisTime.Format(time.RFC3339))

	for _, sopt := range batchOpts.sources.Value() {
		name, url, ok := strings.Cut(sopt, "=")
		if !ok {
			return fmt.Errorf("source option not valid, use format 'name=url'")
		}

		if _, exists := cfg.Sources[name]; exists {
			return fmt.Errorf("duplicate source %q specified", name)
		}

		if strings.HasPrefix(url, "postgres:") {
			cfg.Sources[name] = NewPgDataSource(url)
		} else {
			return fmt.Errorf("unsupported source url: %s", url)
		}
	}

	if batchOpts.confDir != "" {
		conffs := os.DirFS(batchOpts.confDir)
		colorConfContent, err := fs.ReadFile(conffs, "colors.yaml")
		if err == nil {
			var cd ColorDoc
			if err := yaml.Unmarshal(colorConfContent, &cd); err != nil {
				return fmt.Errorf("failed to unmarshal colors.yaml: %w", err)
			}
			cfg.DefaultColor = cd.Default
			cfg.Colors = make(map[string]string, len(cd.Colors))
			for _, nc := range cd.Colors {
				cfg.Colors[nc.Name] = nc.Color
			}
		} else if !errors.Is(err, fs.ErrNotExist) {
			return fmt.Errorf("failed to read colors: %w", err)
		}
	}

	infs := os.DirFS(batchOpts.inDir)
	outfs := os.DirFS(batchOpts.outDir)
	fnames, err := fs.Glob(infs, "*.yaml")
	if err != nil {
		return fmt.Errorf("failed to read input directory: %w", err)
	}

	grp, ctx := errgroup.WithContext(ctx)
	grp.SetLimit(batchOpts.concurrency)

	for _, fname := range fnames {
		fname := fname
		grp.Go(func() error {
			fcontent, err := fs.ReadFile(infs, fname)
			if err != nil {
				return fmt.Errorf("failed to read plot definition %q: %w", fname, err)
			}

			pd, err := parsePlotDef(fname, fcontent)
			if err != nil {
				return fmt.Errorf("failed to parse plot definition %q: %w", fname, err)
			}

			outFname := pd.Name + ".json"
			outGlob := ""

			if batchOpts.version && pd.Frequency != "" {
				versionSuffix := ""
				globSuffix := ""
				switch strings.ToLower(pd.Frequency) {
				case "weekly":
					versionSuffix = "-" + cfg.BasisTime.Truncate(7*24*time.Hour).Format("2006-01-02")
					globSuffix = "-20[0-9][0-9]-[0-9][0-9]-[0-9][0-9]"
				case "daily":
					versionSuffix = "-" + cfg.BasisTime.Truncate(24*time.Hour).Format("2006-01-02")
					globSuffix = "-20[0-9][0-9]-[0-9][0-9]-[0-9][0-9]"
				case "hourly":
					versionSuffix = "-" + cfg.BasisTime.Truncate(time.Hour).Format("2006-01-02-15")
					globSuffix = "-20[0-9][0-9]-[0-9][0-9]-[0-9][0-9]-[0-9][0-9]"
				default:
					return fmt.Errorf("unsupported plot frequency: %q", pd.Frequency)
				}

				outFname = pd.Name + versionSuffix + ".json"
				outGlob = pd.Name + globSuffix + ".json"

			}

			output := filepath.Join(batchOpts.outDir, outFname)

			if err := pd.ExecuteTemplates(ctx, cfg); err != nil {
				return fmt.Errorf("failed to execute templates for plot definition: %w", err)
			}

			if batchOpts.validate {
				fmt.Println("Name: " + pd.Name)
				fmt.Println("Frequency: " + pd.Frequency)
				fmt.Println("Output: " + output)

				fmt.Println("Datasets:")
				for _, ds := range pd.Datasets {
					fmt.Println("  Name: " + ds.Name)
					fmt.Println("  Source: " + ds.Source)
					fmt.Println("  Query:")
					fmt.Println(indent(ds.Query, "      "))

				}

				return nil
			}

			if !batchOpts.force {
				info, err := os.Lstat(filepath.Join(batchOpts.inDir, fname))
				if err != nil {
					return err
				}

				exists, err := fileExistsAndIsNewerThan(filepath.Join(batchOpts.outDir, outFname), info.ModTime())
				if err != nil {
					return err
				}
				if exists {
					slog.Info("skipping plot, output already exists", "name", pd.Name)
					return nil
				}
			}

			slog.Info("generating plot", "name", pd.Name)
			// set up a monitoring loop that reports progress for long running queries
			done := make(chan struct{})
			t := time.NewTicker(time.Minute)
			go func() {
				start := time.Now()
				defer t.Stop()
				for {
					select {
					case <-t.C:
						slog.Info("still generating plot", "name", pd.Name, "elapsed", time.Since(start))
					case <-done:
						return
					}
				}
			}()
			fig, err := generateFig(ctx, pd, cfg)
			close(done) // stop the monitoring loop

			if err != nil {
				return fmt.Errorf("failed to generate plot %q: %w", pd.Name, err)
			}

			var data []byte
			if batchOpts.compact {
				data, err = json.Marshal(fig)
			} else {
				data, err = json.MarshalIndent(fig, "", "  ")
			}
			if err != nil {
				return fmt.Errorf("failed to marshal to json: %w", err)
			}

			slog.Info("writing plot output", "name", pd.Name, "filename", output)
			if err := writeOutput(output, data); err != nil {
				return fmt.Errorf("write output: %w", err)
			}

			if batchOpts.version && pd.Frequency != "" {
				// check if this file would sort after all existing files
				existing, err := fs.Glob(outfs, outGlob)
				if err != nil {
					return fmt.Errorf("failed to read existing plots from output directory: %w", err)
				}

				sort.Strings(existing)
				// is the last file the one we just wrote? if so then we assume it is latest
				if existing[len(existing)-1] == outFname {
					latest := filepath.Join(batchOpts.outDir, pd.Name+"-latest.json")
					slog.Info("writing plot output", "name", pd.Name, "filename", latest)
					if err := writeOutput(latest, data); err != nil {
						return fmt.Errorf("write latest output: %w", err)
					}
				}
			}

			return nil
		})
	}

	if err := grp.Wait(); err != nil {
		return err
	}

	return nil
}

func fileExists(fname string) (bool, error) {
	_, err := os.Lstat(fname)
	if err == nil {
		return true, nil
	}
	if !errors.Is(err, fs.ErrNotExist) {
		return true, fmt.Errorf("failed to stat file: %w", err)
	}

	return false, nil
}

func fileExistsAndIsNewerThan(fname string, ts time.Time) (bool, error) {
	info, err := os.Lstat(fname)
	if err == nil {
		return info.ModTime().After(ts), nil
	}
	if !errors.Is(err, fs.ErrNotExist) {
		return true, fmt.Errorf("failed to stat file: %w", err)
	}

	return false, nil
}

func writeOutput(fname string, data []byte) error {
	f, err := os.Create(fname)
	if err != nil {
		return fmt.Errorf("failed to create output file: %w", err)
	}
	defer f.Close()

	_, err = fmt.Fprintln(f, string(data))
	if err != nil {
		return fmt.Errorf("failed to write output file: %w", err)
	}
	return nil
}
