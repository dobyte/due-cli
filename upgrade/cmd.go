package upgrade

import (
	"fmt"

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
		if err := exec.Upgrade(ctx.String(flag.Dir.Name), ctx.String(flag.Version.Name)); err != nil {
			fmt.Printf("due upgrade failed: %v\n", err)
		} else {
			fmt.Printf("due upgrade success\n")
		}

		return nil
	},
}
