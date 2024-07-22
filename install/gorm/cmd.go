package gorm

import (
	"github.com/dobyte/due-cli/internal/exec"
	"github.com/urfave/cli/v2"
)

var Command = &cli.Command{
	Name:  "gorm",
	Usage: "Install the mongo toolchain",
	Action: func(ctx *cli.Context) error {
		exec.Install(exec.Package{
			Name:    "gorm-dao-generator",
			Module:  "github.com/dobyte/gorm-dao-generator",
			Version: "latest",
		})

		return nil
	},
}
