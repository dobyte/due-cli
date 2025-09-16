package node

import (
	"github.com/dobyte/due-cli/internal/flag"
	"github.com/urfave/cli/v2"
)

var Command = &cli.Command{
	Name:  "node",
	Usage: "create a new node project",
	Flags: []cli.Flag{
		flag.Dir,
		flag.Name,
		flag.Module,
		flag.Codec,
		flag.Locator,
		flag.Registry,
		flag.Transporter,
	},
	Action: func(ctx *cli.Context) error {
		// var (
		// 	dir         = ctx.String("dir")
		// 	name        = ctx.String("name")
		// 	module      = ctx.String("module")
		// 	codec       = ctx.String("codec")
		// 	locator     = ctx.String("locator")
		// 	registry    = ctx.String("registry")
		// 	transporter = ctx.String("transporter")
		// 	replaces    = map[string]string{
		// 		"VarName":        name,
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
		// etc.AddCluster(cluster.Node)
		// etc.AddLocator(locator)
		// etc.AddRegistry(registry)
		// etc.AddTransportClient(transporter)
		// etc.AddTransportServer(transporter)

		// if err := gen.NewGenerator(dir).Make(&gen.Makefile{
		// 	Out:      template.MainOutput,
		// 	Tpl:      template.MainTemplate,
		// 	Replaces: replaces,
		// }, &gen.Makefile{
		// 	Out:      template.AppOutput,
		// 	Tpl:      template.AppTemplate,
		// 	Replaces: replaces,
		// }, &gen.Makefile{
		// 	Out:      template.RouteOutput,
		// 	Tpl:      template.RouteTemplate,
		// 	Replaces: replaces,
		// }, &gen.Makefile{
		// 	Out:      template.CoreOutput,
		// 	Tpl:      template.CoreTemplate,
		// 	Replaces: replaces,
		// }, &gen.Makefile{
		// 	Out:      template.ProtocolOutput,
		// 	Tpl:      template.ProtocolTemplate,
		// 	Replaces: replaces,
		// }, &gen.Makefile{
		// 	Out:      etc.Output(),
		// 	Tpl:      etc.Template(),
		// 	Replaces: replaces,
		// }); err != nil {
		// 	return cli.Exit(err, 1)
		// }

		// if err := gen.MakeGlobalFile(replaces); err != nil {
		// 	return cli.Exit(err, 1)
		// }

		return nil
	},
}
