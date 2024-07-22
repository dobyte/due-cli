package gate

import (
	"github.com/dobyte/due/cmd/v2/create/gate/template"
	"github.com/dobyte/due/cmd/v2/internal/etc"
	"github.com/dobyte/due/cmd/v2/internal/etc/cluster"
	"github.com/dobyte/due/cmd/v2/internal/flag"
	"github.com/dobyte/due/cmd/v2/internal/gen"
	"github.com/dobyte/due/cmd/v2/internal/version"
	"github.com/urfave/cli/v2"
)

var Command = &cli.Command{
	Name:  "gate",
	Usage: "Create a new gate project",
	Flags: []cli.Flag{
		flag.Name,
		flag.Dir,
		flag.Network,
		flag.Locator,
		flag.Registry,
	},
	Action: func(ctx *cli.Context) error {
		var (
			name     = ctx.String("name")
			dir      = ctx.String("dir")
			network  = ctx.String("network")
			locator  = ctx.String("locator")
			registry = ctx.String("registry")
			replaces = map[string]string{
				"VarName":             name,
				"VarNetwork":          network,
				"VarLocator":          locator,
				"VarRegistry":         registry,
				"VarGoVersion":        version.GoVersion,
				"VarFrameworkVersion": version.FrameworkVersion,
			}
		)

		etc := etc.NewEtc()
		etc.AddLog()
		etc.AddPacket()
		etc.AddCluster(cluster.Gate)
		etc.AddLocator(locator)
		etc.AddRegistry(registry)
		etc.AddNetworkServer(network)

		return gen.NewGenerator(name, dir).Make(
			&gen.Makefile{
				Out:      template.GoModOutput,
				Tpl:      template.GoModTemplate,
				Replaces: replaces,
			},
			&gen.Makefile{
				Out:      template.MainOutput,
				Tpl:      template.MainTemplate,
				Replaces: replaces,
			},
			&gen.Makefile{
				Out:      etc.Output(),
				Tpl:      etc.Template(),
				Replaces: replaces,
			},
		)
	},
}
