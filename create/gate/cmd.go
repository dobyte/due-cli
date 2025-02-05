package gate

import (
	"github.com/dobyte/due-cli/create/gate/template"
	"github.com/dobyte/due-cli/internal/etc"
	"github.com/dobyte/due-cli/internal/etc/cluster"
	"github.com/dobyte/due-cli/internal/flag"
	"github.com/dobyte/due-cli/internal/gen"
	"github.com/dobyte/due-cli/internal/version"
	"github.com/urfave/cli/v2"
	"path"
)

var Command = &cli.Command{
	Name:  "gate",
	Usage: "create a new gate project",
	Flags: []cli.Flag{
		flag.Name,
		flag.Module,
		flag.Dir,
		flag.Network,
		flag.Locator,
		flag.Registry,
	},
	Action: func(ctx *cli.Context) error {
		var (
			output   string
			name     = ctx.String("name")
			module   = ctx.String("module")
			dir      = ctx.String("dir")
			network  = ctx.String("network")
			locator  = ctx.String("locator")
			registry = ctx.String("registry")
			replaces = map[string]string{
				"VarModule":    module,
				"VarNetwork":   network,
				"VarLocator":   locator,
				"VarRegistry":  registry,
				"VarGoVersion": version.GoVersion,
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
		etc.AddPacket()
		etc.AddCluster(cluster.Gate)
		etc.AddLocator(locator)
		etc.AddRegistry(registry)
		etc.AddNetworkServer(network)

		makefiles := make([]*gen.Makefile, 0, 3)
		makefiles = append(makefiles, &gen.Makefile{
			Out:      template.MainOutput,
			Tpl:      template.MainTemplate,
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
