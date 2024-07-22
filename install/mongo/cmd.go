package mongo

import (
	"github.com/dobyte/due/cmd/v2/internal/exec"
	"github.com/urfave/cli/v2"
)

var Command = &cli.Command{
	Name:  "mongo",
	Usage: "Install the mongo toolchain",
	Action: func(ctx *cli.Context) error {
		exec.Install(exec.Package{
			Name:    "mongo-dao-generator",
			Module:  "github.com/dobyte/mongo-dao-generator",
			Version: "latest",
		})

		return nil
	},
}
