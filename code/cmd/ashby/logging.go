package main

import (
	"github.com/iand/pontium/hlog"
	"github.com/urfave/cli/v2"
	"golang.org/x/exp/slog"
)

var loggingFlags = []cli.Flag{
	&cli.BoolFlag{
		Name:        "verbose",
		Aliases:     []string{"v"},
		Usage:       "Set logging level more verbose to include info level logs",
		Value:       true,
		Destination: &loggingOpts.Verbose,
	},

	&cli.BoolFlag{
		Name:        "veryverbose",
		Aliases:     []string{"vv"},
		Usage:       "Set logging level more verbose to include debug level logs",
		Destination: &loggingOpts.VeryVerbose,
	},
}

var loggingOpts struct {
	Verbose     bool
	VeryVerbose bool
}

func setupLogging() {
	logLevel := new(slog.LevelVar)
	logLevel.Set(slog.LevelWarn)
	if loggingOpts.Verbose {
		logLevel.Set(slog.LevelInfo)
	}
	if loggingOpts.VeryVerbose {
		logLevel.Set(slog.LevelDebug)
	}

	h := new(hlog.Handler)
	h = h.WithLevel(logLevel.Level())
	slog.SetDefault(slog.New(h))
}
