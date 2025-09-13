package main

import (
	"fmt"
	"os"

	"github.com/dobyte/due-cli/create"
	"github.com/dobyte/due-cli/install"
	"github.com/dobyte/due-cli/internal/version"
	"github.com/dobyte/due-cli/upgrade"
	"github.com/urfave/cli/v2"
)

func main() {
	cli.VersionFlag = &cli.BoolFlag{
		Name:               "version",
		Aliases:            []string{"v"},
		Usage:              "show the duc-cli version information",
		DisableDefaultText: true,
	}

	app := &cli.App{
		Name:        "due-cli",
		Version:     version.ToolVersion,
		Description: "due project building tool",
		Authors: []*cli.Author{{
			Name:  "fuxiao",
			Email: "yuebanfuxiao@gmail.com",
		}},
		Commands: []*cli.Command{
			create.Command,
			install.Command,
			upgrade.Command,
		},
	}

	if err := app.Run(os.Args); err != nil {
		fmt.Println(err)
	}
}
