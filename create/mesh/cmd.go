package mesh

import (
	"github.com/dobyte/due-cli/create/mesh/template"
	"github.com/dobyte/due-cli/internal/etc"
	"github.com/dobyte/due-cli/internal/etc/cluster"
	"github.com/dobyte/due-cli/internal/flag"
	"github.com/dobyte/due-cli/internal/gen"
	"github.com/dobyte/due-cli/internal/version"
	"github.com/urfave/cli/v2"
)

var Command = &cli.Command{
	Name:  "mesh",
	Usage: "Create a new mesh project",
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

		etc := etc.NewEtc()
		etc.AddLog()
		etc.AddTask()
		etc.AddPacket()
		etc.AddCluster(cluster.Mesh)
		etc.AddLocator(locator)
		etc.AddRegistry(registry)
		etc.AddTransportServer(transporter)
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
				Out:      etc.Output(),
				Tpl:      etc.Template(),
				Replaces: replaces,
			},
		)
	},
}
