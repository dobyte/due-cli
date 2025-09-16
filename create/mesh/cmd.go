package mesh

import (
	"github.com/dobyte/due-cli/internal/flag"
	"github.com/urfave/cli/v2"
)

var Command = &cli.Command{
	Name:  "mesh",
	Usage: "create a new mesh project",
	Flags: []cli.Flag{
		flag.Dir,
		// flag.Alone,
		flag.Module,
		flag.Codec,
		flag.Locator,
		flag.Registry,
		flag.Transporter,
	},
	Action: func(ctx *cli.Context) error {
		// var (
		// 	dir         = ctx.String("dir")
		// 	alone       = ctx.Bool("alone")
		// 	module      = ctx.String("module")
		// 	codec       = ctx.String("codec")
		// 	locator     = ctx.String("locator")
		// 	registry    = ctx.String("registry")
		// 	transporter = ctx.String("transporter")
		// 	replaces    = map[string]string{
		// 		"VarModule":      module,
		// 		"VarCodec":       codec,
		// 		"VarLocator":     locator,
		// 		"VarRegistry":    registry,
		// 		"VarTransporter": transporter,
		// 		// "VarGoVersion":   version.GoVersion,
		// 	}
		// )

		// if module == "" {
		// 	return cli.Exit("the module of the project is required.", 1)
		// }

		// if ok, err := os.IsEmptyDir(dir); err != nil {
		// 	return cli.Exit(err, 1)
		// } else if !ok {
		// 	return cli.Exit("the project dir is not empty.", 1)
		// }

		// etc := etc.NewEtc()
		// etc.AddLog()
		// etc.AddTask()
		// etc.AddPacket()
		// etc.AddCluster(cluster.Mesh)
		// etc.AddLocator(locator)
		// etc.AddRegistry(registry)
		// etc.AddTransportServer(transporter)
		// etc.AddTransportClient(transporter)

		// makefiles := make([]*gen.Makefile, 0, 5)
		// makefiles = append(makefiles, &gen.Makefile{
		// 	Out:      template.MainOutput,
		// 	Tpl:      template.MainTemplate,
		// 	Replaces: replaces,
		// }, &gen.Makefile{
		// 	Out:      template.AppOutput,
		// 	Tpl:      template.AppTemplate,
		// 	Replaces: replaces,
		// }, &gen.Makefile{
		// 	Out:      etc.Output(),
		// 	Tpl:      etc.Template(),
		// 	Replaces: replaces,
		// }, &gen.Makefile{
		// 	Out:      tpl.GitignoreOutput,
		// 	Tpl:      tpl.GitignoreTemplate,
		// 	Replaces: replaces,
		// })

		// if alone {
		// 	makefiles = append(makefiles, &gen.Makefile{
		// 		Out:      tpl.GoModOutput,
		// 		Tpl:      tpl.GoModTemplate,
		// 		Replaces: replaces,
		// 	})
		// } else {
		// 	if err := gen.MakeGlobalFile(replaces); err != nil {
		// 		return cli.Exit(err, 1)
		// 	}
		// }

		// if err := gen.NewGenerator(dir).Make(makefiles...); err != nil {
		// 	return cli.Exit(err, 1)
		// }

		return nil
	},
}
