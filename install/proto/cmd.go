package proto

import (
	"github.com/dobyte/due/cmd/v2/internal/exec"
	"github.com/urfave/cli/v2"
)

var Command = &cli.Command{
	Name:  "proto",
	Usage: "Install the protobuf toolchain",
	Action: func(ctx *cli.Context) error {
		exec.Install(exec.Package{
			Name:    "protoc-gen-gofast",
			Module:  "github.com/gogo/protobuf/protoc-gen-gofast",
			Version: "latest",
		})

		return nil
	},
}
