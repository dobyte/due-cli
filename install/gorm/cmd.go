package gorm

import (
	"github.com/dobyte/due-cli/internal/exec"
	"github.com/dobyte/due-cli/internal/flag"
	"github.com/urfave/cli/v2"
)

var Command = &cli.Command{
	Name:  "gorm",
	Usage: "install the gorm toolchain",
	Flags: []cli.Flag{
		flag.Version,
	},
	Action: func(ctx *cli.Context) error {
		exec.Install(exec.Package{
			Name:    "gorm-dao-generator",
			Module:  "github.com/dobyte/gorm-dao-generator",
			Version: ctx.String(flag.Version.Name),
		})

		return nil
	},
}
