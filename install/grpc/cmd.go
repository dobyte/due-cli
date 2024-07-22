package grpc

import (
	"github.com/dobyte/due-cli/internal/exec"
	"github.com/urfave/cli/v2"
)

var Command = &cli.Command{
	Name:  "grpc",
	Usage: "Install the grpc toolchain",
	Action: func(ctx *cli.Context) error {
		exec.Install(exec.Package{
			Name:    "protoc-gen-go",
			Module:  "google.golang.org/protobuf/cmd/protoc-gen-go",
			Version: "latest",
		}, exec.Package{
			Name:    "protoc-gen-go-grpc",
			Module:  "google.golang.org/grpc/cmd/protoc-gen-go-grpc",
			Version: "latest",
		})

		return nil
	},
}
