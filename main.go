package main

import (
	"fmt"
	"github.com/dobyte/due/cmd/v2/create"
	"github.com/dobyte/due/cmd/v2/install"
	"github.com/dobyte/due/cmd/v2/internal/version"
	"github.com/urfave/cli/v2"
	"os"
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
		},
	}

	if err := app.Run(os.Args); err != nil {
		fmt.Println(err)
	}
}
