package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/MetalBlueberry/go-plotly/offline"
	"github.com/urfave/cli/v2"
	"gopkg.in/yaml.v3"
)

var plotCommand = &cli.Command{
	Name:   "plot",
	Usage:  "Interactive command to generate a single plot",
	Action: Plot,
	Flags: append([]cli.Flag{
		&cli.BoolFlag{
			Name:        "preview",
			Required:    false,
			Usage:       "Preview the plot in a browser window.",
			Destination: &plotOpts.preview,
		},
		&cli.BoolFlag{
			Name:        "compact",
			Required:    false,
			Usage:       "Emit compact json instead of pretty-printed.",
			Destination: &plotOpts.compact,
		},
		&cli.BoolFlag{
			Name:        "validate",
			Required:    false,
			Usage:       "Validate the input file without running queries.",
			Destination: &plotOpts.validate,
		},
		&cli.StringSliceFlag{
			Name:        "source",
			Aliases:     []string{"s"},
			Required:    false,
			Usage:       "Specify the url of a data source, in the format name=url. May be repeated to specify multiple sources. Postgres urls take the form 'postgres://username:password@hostname:5432/database_name'",
			Destination: &plotOpts.sources,
		},
		&cli.StringFlag{
			Name:        "output",
			Aliases:     []string{"o"},
			Required:    false,
			Usage:       "Name of file JSON output should be written to. Output will be emitted to stdout by default.",
			Destination: &plotOpts.output,
		},
	}, loggingFlags...),
}

var plotOpts struct {
	preview  bool
	compact  bool
	sources  cli.StringSlice
	output   string
	validate bool
}

func Plot(cc *cli.Context) error {
	ctx := cc.Context
	setupLogging()

	cfg := &PlotConfig{
		BasisTime: time.Now().UTC(),
		Sources: map[string]DataSource{
			"demo": &DemoDataSource{},
		},
	}

	for _, sopt := range plotOpts.sources.Value() {
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

	if cc.NArg() != 1 {
		return fmt.Errorf("plot defintion must be supplied as an argument")
	}

	fname := cc.Args().Get(0)

	fcontent, err := os.ReadFile(fname)
	if err != nil {
		return fmt.Errorf("failed to read plot specification: %w", err)
	}

	var pd PlotDef
	if err := yaml.Unmarshal(fcontent, &pd); err != nil {
		return fmt.Errorf("failed to unmarshal plot definition: %w", err)
	}

	if pd.Name == "" {
		pd.Name = plotname(fname)
	}

	if err := pd.ExecuteTemplates(ctx, cfg); err != nil {
		return fmt.Errorf("failed to execute templates for plot definition: %w", err)
	}

	if plotOpts.validate {
		fmt.Println("Name: " + pd.Name)
		fmt.Println("Frequency: " + pd.Frequency)
		fmt.Println("Datasets:")
		for _, ds := range pd.Datasets {
			fmt.Println("  Name: " + ds.Name)
			fmt.Println("  Source: " + ds.Source)
			fmt.Println("  Query:")
			fmt.Println(indent(ds.Query, "      "))

		}

		return nil
	}

	fig, err := generateFig(ctx, &pd, cfg)
	if err != nil {
		return fmt.Errorf("failed to generate plot: %w", err)
	}

	var data []byte
	if plotOpts.compact {
		data, err = json.Marshal(fig)
	} else {
		data, err = json.MarshalIndent(fig, "", "  ")
	}
	if err != nil {
		return fmt.Errorf("failed to marshal to json: %w", err)
	}

	var out io.Writer = os.Stdout
	if plotOpts.output != "" {
		f, err := os.Create(plotOpts.output)
		if err != nil {
			return fmt.Errorf("failed to create output file: %w", err)
		}
		defer f.Close()
		out = f
	}

	fmt.Fprintln(out, string(data))

	if plotOpts.preview {
		offline.Show(fig)
	}
	return nil
}

type DemoDataSource struct{}

func (s *DemoDataSource) GetDataSet(_ context.Context, query string, params ...any) (DataSet, error) {
	switch query {
	case "populations":
		return &StaticDataSet{Data: map[string][]any{
			"creature": {"giraffes", "orangutans", "monkeys"},
			"month1":   {20, 14, 23},
			"month2":   {2, 18, 29},
		}}, nil
	default:
		return nil, fmt.Errorf("unknown demo dataset: %s", query)
	}
}

func indent(s string, prefix string) string {
	s = strings.ReplaceAll(s, "\n", "\n"+prefix)
	return prefix + s
}

func plotname(fname string) string {
	base := filepath.Base(fname)
	return strings.TrimSuffix(base, filepath.Ext(fname))
}
