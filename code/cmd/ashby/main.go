package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli/v2"
)

const (
	appName   = "ashby"
	envPrefix = "ASHBY_"
)

func main() {
	app := &cli.App{
		Name:     appName,
		HelpName: appName,
		Usage:    "Plot server",
		Commands: []*cli.Command{
			plotCommand,
			batchCommand,
		},
	}

	if err := app.Run(os.Args); err != nil {
		fmt.Fprintf(os.Stderr, "%+v\n", err)
		os.Exit(1)
	}
}
