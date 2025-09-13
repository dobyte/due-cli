package gate

import (
	"fmt"

	"github.com/AlecAivazis/survey/v2"
	"github.com/urfave/cli/v2"
)

var Command = &cli.Command{
	Name:  "gate",
	Usage: "create a new gate project",
	// Flags: []cli.Flag{
	// 	flag.Dir,
	// 	// flag.Alone,
	// 	flag.Module,
	// 	flag.Network,
	// 	flag.Locator,
	// 	flag.Registry,
	// },
	Action: func(ctx *cli.Context) error {
		var qs = []*survey.Question{
			{
				Name:      "name",
				Prompt:    &survey.Input{Message: "What is your name?"},
				Validate:  survey.Required,
				Transform: survey.Title,
			},
			{
				Name: "color",
				Prompt: &survey.Select{
					Message: "Choose a color:",
					Options: []string{"red", "blue", "green"},
					Default: "red",
				},
			},
			{
				Name:   "age",
				Prompt: &survey.Input{Message: "How old are you?"},
			},
		}

		// the answers will be written to this struct
		answers := struct {
			Name          string // survey will match the question and field names
			FavoriteColor string `survey:"color"` // or you can tag fields to match a specific name
			Age           int    // if the types don't match, survey will convert it
		}{}

		// perform the questions
		if err := survey.Ask(qs, &answers); err != nil {
			return err
		}

		fmt.Printf("%s chose %s.", answers.Name, answers.FavoriteColor)

		// var (
		// 	dir = ctx.String(flag.Dir.Name)
		// 	// alone    = ctx.Bool(flag.Alone.Name)
		// 	module   = ctx.String(flag.Module.Name)
		// 	network  = ctx.String(flag.Network.Name)
		// 	locator  = ctx.String(flag.Locator.Name)
		// 	registry = ctx.String(flag.Registry.Name)
		// 	replaces = map[string]string{
		// 		"VarModule":    module,
		// 		"VarNetwork":   network,
		// 		"VarLocator":   locator,
		// 		"VarRegistry":  registry,
		// 		"VarGoVersion": version.GoVersion,
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
		// etc.AddPacket()
		// etc.AddCluster(cluster.Gate)
		// etc.AddLocator(locator)
		// etc.AddRegistry(registry)
		// etc.AddNetworkServer(network)

		// makefiles := make([]*gen.Makefile, 0, 3)
		// makefiles = append(makefiles, &gen.Makefile{
		// 	Out:      template.MainOutput,
		// 	Tpl:      template.MainTemplate,
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
