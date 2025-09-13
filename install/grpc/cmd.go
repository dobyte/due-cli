package grpc

import (
	"github.com/dobyte/due-cli/internal/exec"
	"github.com/dobyte/due-cli/internal/flag"
	"github.com/urfave/cli/v2"
)

var Command = &cli.Command{
	Name:  "grpc",
	Usage: "install the grpc toolchain",
	Flags: []cli.Flag{
		flag.Version,
	},
	Action: func(ctx *cli.Context) error {
		exec.Install(exec.Package{
			Name:    "protoc-gen-go-grpc",
			Module:  "google.golang.org/grpc/cmd/protoc-gen-go-grpc",
			Version: ctx.String(flag.Version.Name),
		})

		return nil
	},
}
