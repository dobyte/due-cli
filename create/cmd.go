package create

import (
	"github.com/dobyte/due-cli/create/gate"
	"github.com/dobyte/due-cli/create/node"
	"github.com/urfave/cli/v2"
)

var Command = &cli.Command{
	Name:        "create",
	Usage:       "create a new cluster project",
	Description: "create a new cluster project",
	Subcommands: []*cli.Command{
		gate.Command,
		node.Command,
	},
	//Flags: []cli.Flag{
	//	flag.Dir,
	//	flag.Module,
	//	flag.Codec,
	//	flag.Network,
	//	flag.Locator,
	//	flag.Registry,
	//	flag.Transporter,
	//},
	//Action: func(ctx *cli.Context) error {
	//	var (
	//		dir         = ctx.String("dir")
	//		module      = ctx.String("module")
	//		codec       = ctx.String("codec")
	//		network     = ctx.String("network")
	//		locator     = ctx.String("locator")
	//		registry    = ctx.String("registry")
	//		transporter = ctx.String("transporter")
	//		replaces    = map[string]string{
	//			"VarModule":      module,
	//			"VarCodec":       codec,
	//			"VarLocator":     locator,
	//			"VarRegistry":    registry,
	//			"VarTransporter": transporter,
	//			"VarGoVersion":   version.GoVersion,
	//		}
	//	)
	//
	//	if module == "" {
	//		return cli.Exit("the module of the project is required.", 1)
	//	}
	//
	//	if ok, err := os.IsEmptyDir(dir); err != nil {
	//		return cli.Exit(err, 1)
	//	} else if !ok {
	//		return cli.Exit("the project dir is not empty.", 1)
	//	}
	//
	//	if err := gate.Make(dir, locator, registry, network); err != nil {
	//		return cli.Exit(err, 1)
	//	}
	//
	//	if err := gen.NewGenerator(dir).Make(&gen.Makefile{
	//		Out:      tpl.GoModOutput,
	//		Tpl:      tpl.GoModTemplate,
	//		Replaces: replaces,
	//	}, &gen.Makefile{
	//		Out:      tpl.GitignoreOutput,
	//		Tpl:      tpl.GitignoreTemplate,
	//		Replaces: replaces,
	//	}); err != nil {
	//		return cli.Exit(err, 1)
	//	}
	//
	//	return nil
	//},
}
