package rpcx

import (
	"github.com/dobyte/due-cli/internal/exec"
	"github.com/urfave/cli/v2"
)

var Command = &cli.Command{
	Name:  "rpcx",
	Usage: "Install the rpcx toolchain",
	Action: func(ctx *cli.Context) error {
		exec.Install(exec.Package{
			Name:    "protoc-gen-rpcx",
			Module:  "github.com/rpcxio/protoc-gen-rpcx",
			Version: "latest",
		})

		return nil
	},
}
