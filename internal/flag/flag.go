package flag

import "github.com/urfave/cli/v2"

var (
	Dir = &cli.StringFlag{
		Name:        "dir",
		Usage:       "specify the project directory",
		Value:       "./",
		DefaultText: "./",
	}

	Alone = &cli.BoolFlag{
		Name:  "alone",
		Usage: "specify whether the project uses alone module",
		Value: false,
	}

	Module = &cli.StringFlag{
		Name:     "module",
		Usage:    "specify the project module",
		Required: true,
	}

	Codec = &cli.StringFlag{
		Name:        "codec",
		Usage:       "specify the project codec; optional: json, proto",
		Value:       "json",
		DefaultText: "json",
	}

	Network = &cli.StringFlag{
		Name:        "network",
		Usage:       "specify the project network; optional: ws, tcp, kcp",
		Value:       "ws",
		DefaultText: "ws",
	}

	Locator = &cli.StringFlag{
		Name:        "locator",
		Usage:       "specify the project locator; optional: redis",
		Value:       "redis",
		DefaultText: "redis",
	}

	Registry = &cli.StringFlag{
		Name:        "registry",
		Usage:       "specify the project registry; optional: consul, etcd, nacos",
		Value:       "consul",
		DefaultText: "consul",
	}

	Transporter = &cli.StringFlag{
		Name:        "transporter",
		Usage:       "specify the project transporter; optional: grpc, rpcx",
		Value:       "grpc",
		DefaultText: "grpc",
	}
)

type Flag struct {
	Name      string
	Optionals []string
}
