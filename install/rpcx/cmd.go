package rpcx

import (
	"github.com/dobyte/due-cli/internal/exec"
	"github.com/dobyte/due-cli/internal/flag"
	"github.com/urfave/cli/v2"
)

var Command = &cli.Command{
	Name:  "rpcx",
	Usage: "install the rpcx toolchain",
	Flags: []cli.Flag{
		flag.Version,
	},
	Action: func(ctx *cli.Context) error {
		exec.Install(exec.Package{
			Name:    "protoc-gen-rpcx",
			Module:  "github.com/rpcxio/protoc-gen-rpcx",
			Version: ctx.String(flag.Version.Name),
		})

		return nil
	},
}
