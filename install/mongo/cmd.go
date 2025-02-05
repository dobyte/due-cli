package mongo

import (
	"github.com/dobyte/due-cli/internal/exec"
	"github.com/urfave/cli/v2"
)

var Command = &cli.Command{
	Name:  "mongo",
	Usage: "install the mongo toolchain",
	Action: func(ctx *cli.Context) error {
		exec.Install(exec.Package{
			Name:    "mongo-dao-generator",
			Module:  "github.com/dobyte/mongo-dao-generator",
			Version: "latest",
		})

		return nil
	},
}
