package mongo

import (
	"github.com/dobyte/due-cli/internal/exec"
	"github.com/dobyte/due-cli/internal/flag"
	"github.com/urfave/cli/v2"
)

var Command = &cli.Command{
	Name:  "mongo",
	Usage: "install the mongo toolchain",
	Flags: []cli.Flag{
		flag.Version,
	},
	Action: func(ctx *cli.Context) error {
		exec.Install(exec.Package{
			Name:    "mongo-dao-generator",
			Module:  "github.com/dobyte/mongo-dao-generator",
			Version: ctx.String(flag.Version.Name),
		})

		return nil
	},
}
