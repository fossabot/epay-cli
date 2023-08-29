package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v2"

	"github.com/AH-dark/epay-cli/actions"
)

var (
	version = "dev"
	commit  string
	date    string
)

func main() {
	app := &cli.App{
		Name:        "epay-cli",
		Description: "A command line tool for epay.",
		Usage:       "epay-cli is a command line tool for epay",
		Version:     fmt.Sprintf("%s-%s", version, commit),
		Copyright:   "2022-2023 @AHdark All rights reserved.",
		Authors: []*cli.Author{
			{Name: "AHdark", Email: "ahdark@outlook.com"},
		},
		Metadata: map[string]interface{}{
			"date": date,
		},
		Commands: []*cli.Command{
			actions.MigrateCommand(),
			{
				Name:  "test",
				Usage: "test epay",
				Subcommands: []*cli.Command{
					actions.SubmitCommand(),
					actions.MapiCommand(),
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal("error: ", err)
	}
}
