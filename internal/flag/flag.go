package flag

import "github.com/urfave/cli/v2"

var (
	Name = &cli.StringFlag{
		Name:     "name",
		Usage:    "specify the project name",
		Required: true,
	}

	Dir = &cli.StringFlag{
		Name:        "dir",
		Usage:       "specify the project directory",
		Value:       "./",
		DefaultText: "./",
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
		Usage:       "specify the project transporter; optional: rpcx, grpc",
		Value:       "rpcx",
		DefaultText: "rpcx",
	}
)

type Flag struct {
	Name      string
	Optionals []string
}
