package node

import (
	"github.com/dobyte/due-cli/create/node/template"
	"github.com/dobyte/due-cli/internal/etc"
	"github.com/dobyte/due-cli/internal/etc/cluster"
	"github.com/dobyte/due-cli/internal/flag"
	"github.com/dobyte/due-cli/internal/gen"
	"github.com/dobyte/due-cli/internal/version"
	"github.com/urfave/cli/v2"
	"path"
)

var Command = &cli.Command{
	Name:  "node",
	Usage: "create a new node project",
	Flags: []cli.Flag{
		flag.Name,
		flag.Module,
		flag.Dir,
		flag.Codec,
		flag.Locator,
		flag.Registry,
		flag.Transporter,
	},
	Action: func(ctx *cli.Context) error {
		var (
			output      string
			name        = ctx.String("name")
			module      = ctx.String("module")
			dir         = ctx.String("dir")
			codec       = ctx.String("codec")
			locator     = ctx.String("locator")
			registry    = ctx.String("registry")
			transporter = ctx.String("transporter")
			replaces    = map[string]string{
				"VarModule":      module,
				"VarCodec":       codec,
				"VarLocator":     locator,
				"VarRegistry":    registry,
				"VarTransporter": transporter,
				"VarGoVersion":   version.GoVersion,
			}
		)

		switch {
		case len(module) != 0:
			output = path.Join(dir, path.Base(module))
		case len(name) != 0:
			output = path.Join(dir, name)
		default:
			return cli.Exit("The module or name of the project is required.", 86)
		}

		etc := etc.NewEtc()
		etc.AddLog()
		etc.AddTask()
		etc.AddPacket()
		etc.AddCluster(cluster.Node)
		etc.AddLocator(locator)
		etc.AddRegistry(registry)
		etc.AddTransportServer(transporter)
		etc.AddTransportClient(transporter)

		makefiles := make([]*gen.Makefile, 0, 3)
		makefiles = append(makefiles, &gen.Makefile{
			Out:      template.MainOutput,
			Tpl:      template.MainTemplate,
			Replaces: replaces,
		}, &gen.Makefile{
			Out:      template.AppOutput,
			Tpl:      template.AppTemplate,
			Replaces: replaces,
		}, &gen.Makefile{
			Out:      template.RouteOutput,
			Tpl:      template.RouteTemplate,
			Replaces: replaces,
		}, &gen.Makefile{
			Out:      template.CoreOutput,
			Tpl:      template.CoreTemplate,
			Replaces: replaces,
		}, &gen.Makefile{
			Out:      template.ProtocolOutput,
			Tpl:      template.ProtocolTemplate,
			Replaces: replaces,
		}, &gen.Makefile{
			Out:      etc.Output(),
			Tpl:      etc.Template(),
			Replaces: replaces,
		})

		if len(module) != 0 {
			makefiles = append(makefiles, &gen.Makefile{
				Out:      template.GoModOutput,
				Tpl:      template.GoModTemplate,
				Replaces: replaces,
			})
		}

		return gen.NewGenerator(output).Make(makefiles...)
	},
}
