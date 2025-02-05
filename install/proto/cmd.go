package proto

import (
	"github.com/dobyte/due-cli/internal/exec"
	"github.com/urfave/cli/v2"
)

var Command = &cli.Command{
	Name:  "proto",
	Usage: "install the protobuf toolchain",
	Action: func(ctx *cli.Context) error {
		exec.Install(exec.Package{
			Name:    "protoc-gen-go",
			Module:  "google.golang.org/protobuf/cmd/protoc-gen-go",
			Version: "latest",
		})

		return nil
	},
}
