package upgrade

import (
	"github.com/dobyte/due-cli/internal/exec"
	"github.com/dobyte/due-cli/internal/flag"
	"github.com/urfave/cli/v2"
)

var Command = &cli.Command{
	Name:        "upgrade",
	Usage:       "upgrade or rollback due to the specified version",
	Description: "upgrade or rollback due to the specified version",
	Flags: []cli.Flag{
		flag.Dir,
		flag.Version,
	},
	Action: func(ctx *cli.Context) error {
		exec.Upgrade(ctx.String(flag.Dir.Name), ctx.String(flag.Version.Name))

		return nil
	},
}
