package node

import (
	"github.com/dobyte/due/cmd/v2/create/node/template"
	"github.com/dobyte/due/cmd/v2/internal/etc"
	"github.com/dobyte/due/cmd/v2/internal/etc/cluster"
	"github.com/dobyte/due/cmd/v2/internal/flag"
	"github.com/dobyte/due/cmd/v2/internal/gen"
	"github.com/dobyte/due/cmd/v2/internal/version"
	"github.com/urfave/cli/v2"
)

var Command = &cli.Command{
	Name:  "node",
	Usage: "Create a new node project",
	Flags: []cli.Flag{
		flag.Name,
		flag.Dir,
		flag.Codec,
		flag.Locator,
		flag.Registry,
		flag.Transporter,
	},
	Action: func(ctx *cli.Context) error {
		var (
			name        = ctx.String("name")
			dir         = ctx.String("dir")
			codec       = ctx.String("codec")
			locator     = ctx.String("locator")
			registry    = ctx.String("registry")
			transporter = ctx.String("transporter")
			replaces    = map[string]string{
				"VarName":             name,
				"VarCodec":            codec,
				"VarLocator":          locator,
				"VarRegistry":         registry,
				"VarTransporter":      transporter,
				"VarGoVersion":        version.GoVersion,
				"VarFrameworkVersion": version.FrameworkVersion,
			}
		)

		if len(name) == 0 {
			return cli.Exit("Ginger croutons are not in the soup", 86)
		}

		etc := etc.NewEtc()
		etc.AddLog()
		etc.AddTask()
		etc.AddPacket()
		etc.AddCluster(cluster.Node)
		etc.AddLocator(locator)
		etc.AddRegistry(registry)
		etc.AddTransportClient(transporter)

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
				Out:      template.AppOutput,
				Tpl:      template.AppTemplate,
				Replaces: replaces,
			},
			&gen.Makefile{
				Out:      template.RouteOutput,
				Tpl:      template.RouteTemplate,
				Replaces: replaces,
			},
			&gen.Makefile{
				Out:      template.CoreOutput,
				Tpl:      template.CoreTemplate,
				Replaces: replaces,
			},
			&gen.Makefile{
				Out:      template.ProtocolOutput,
				Tpl:      template.ProtocolTemplate,
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
