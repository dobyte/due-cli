package node

import (
	"fmt"
	"log"
	"path/filepath"

	"github.com/AlecAivazis/survey/v2"
	"github.com/dobyte/due-cli/create/node/template"
	"github.com/dobyte/due-cli/internal/gen"
	"github.com/dobyte/due-cli/internal/mod"
	"github.com/dobyte/due-cli/internal/mod/cluster"
	"github.com/dobyte/due-cli/internal/os"
	vs "github.com/dobyte/due-cli/internal/version"
	"github.com/urfave/cli/v2"
)

const (
	createFailure = "create project failure"
	createSuccess = "create project success"
)

var Command = &cli.Command{
	Name:  "node",
	Usage: "create a new node project",
	Action: func(ctx *cli.Context) error {
		var (
			dir       string
			alone     bool
			variables = &gen.Variables{}
		)

		if err := survey.AskOne(&survey.Confirm{
			Message: "whether it is an alone project:",
			Default: false,
		}, &alone); err != nil {
			log.Fatal(createFailure, err)
		}

		if err := survey.AskOne(&survey.Input{
			Message: "specify the project directory:",
			Default: "./testdata",
		}, &dir); err != nil {
			log.Fatal(createFailure, err)
		}

		if err := survey.AskOne(&survey.Input{
			Message: "specify the project name:",
			Default: "node",
		}, &variables.Name); err != nil {
			log.Fatal(createFailure, err)
		}

		if err := survey.AskOne(&survey.Select{
			Message: "choose a codec protocol:",
			Options: []string{"json", "proto"},
			Default: "json",
		}, &variables.Codec); err != nil {
			log.Fatal(createFailure, err)
		}

		isNeedGenMod := alone || !os.IsFile(filepath.Join(dir, "go.mod"))

		if isNeedGenMod {
			var defaultModule string

			if alone {
				defaultModule = fmt.Sprintf("github.com/due/%s", variables.Name)
			} else {
				defaultModule = "github.com/due"
			}

			if err := survey.AskOne(&survey.Input{
				Message: "specify the project module:",
				Default: defaultModule,
			}, &variables.Module); err != nil {
				log.Fatal(createFailure, err)
			}

			if err := survey.AskOne(&survey.Input{
				Message: "specify the go version:",
			}, &variables.GoVersion); err != nil {
				log.Fatal(createFailure, err)
			}

			if err := survey.AskOne(&survey.Input{
				Message: "specify the due version:",
				Default: "latest",
			}, &variables.DueFullVersion); err != nil {
				log.Fatal(createFailure, err)
			}
		}

		if err := survey.AskOne(&survey.Select{
			Message: "choose a locate component:",
			Options: []string{"redis"},
			Default: "redis",
		}, &variables.Locate); err != nil {
			log.Fatal(createFailure, err)
		}

		if err := survey.AskOne(&survey.Select{
			Message: "choose a registry component:",
			Options: []string{"nacos", "etcd", "consul"},
			Default: "nacos",
		}, &variables.Registry); err != nil {
			log.Fatal(createFailure, err)
		}

		if err := survey.AskOne(&survey.Select{
			Message: "choose a config component:",
			Options: []string{"none", "file", "nacos", "etcd", "consul"},
			Default: "none",
		}, &variables.Config); err != nil {
			log.Fatal(createFailure, err)
		}

		if err := survey.AskOne(&survey.Select{
			Message: "choose a lock component:",
			Options: []string{"none", "redis", "memcache"},
			Default: "none",
		}, &variables.Lock); err != nil {
			log.Fatal(createFailure, err)
		}

		if err := survey.AskOne(&survey.Select{
			Message: "choose a cache component:",
			Options: []string{"none", "redis", "memcache"},
			Default: "none",
		}, &variables.Cache); err != nil {
			log.Fatal(createFailure, err)
		}

		if err := survey.AskOne(&survey.Select{
			Message: "choose a eventbus component:",
			Options: []string{"none", "redis", "memcache"},
			Default: "none",
		}, &variables.Eventbus); err != nil {
			log.Fatal(createFailure, err)
		}

		if err := survey.AskOne(&survey.Select{
			Message: "choose a rpc transport component:",
			Options: []string{"none", "grpc", "rpcx"},
			Default: "none",
		}, &variables.Transport); err != nil {
			log.Fatal(createFailure, err)
		}

		if ok, err := os.IsEmptyDir(filepath.Join(dir, variables.Name)); err != nil {
			log.Fatal(createFailure, err)
		} else if !ok {
			log.Fatal(createFailure, "the project dir is not empty.")
		}

		var commands *mod.Commands

		if isNeedGenMod {
			if variables.Module == "" {
				log.Fatal(createFailure, "the module of the project is required.")
			}

			if variables.GoVersion == "" {
				if v, err := vs.ParseGoVersion(); err != nil {
					log.Fatal(createFailure, err)
				} else {
					variables.GoVersion = v
				}
			}

			if variables.DueFullVersion == "" {
				variables.DueFullVersion = "latest"
			}

			full, major, sha, err := vs.ParseDueVersion(variables.DueFullVersion)
			if err != nil {
				log.Fatal(createFailure, err)
			}

			variables.DueFullVersion = full
			variables.DueMajorVersion = major

			if alone {
				commands = mod.NewCommands(filepath.Join(dir, variables.Name), full, sha)
			} else {
				commands = mod.NewCommands(dir, full, sha)
			}

			commands.AddBase()
		} else {
			full, major, sha, err := vs.ParseDueVersionFromGoMod(dir)
			if err != nil {
				log.Fatal(createFailure, err)
			}

			variables.DueFullVersion = full
			variables.DueMajorVersion = major

			commands = mod.NewCommands(dir, full, sha)
		}

		etc := mod.NewEtc()
		etc.AddLog()
		etc.AddPacket()
		etc.AddCluster(cluster.Node)
		etc.AddLock(variables.Lock)
		etc.AddCache(variables.Cache)
		etc.AddConfig(variables.Config)
		etc.AddLocate(variables.Locate)
		etc.AddRegistry(variables.Registry)
		etc.AddEventbus(variables.Eventbus)
		etc.AddTransportClient(variables.Transport)
		etc.AddTransportServer(variables.Transport)

		makefiles := make([]*gen.Makefile, 0, 4)
		makefiles = append(makefiles, &gen.Makefile{
			Out:       template.MainOutput,
			Tpl:       template.MainTemplate,
			Variables: variables,
		}, &gen.Makefile{
			Out:       etc.Output(),
			Tpl:       etc.Template(),
			Variables: variables,
		})

		if err := gen.NewGenerator(filepath.Join(dir, variables.Name)).Make(makefiles...); err != nil {
			log.Fatal(createFailure, err)
		}

		commands.AddLock(variables.Lock)
		commands.AddCache(variables.Cache)
		commands.AddConfig(variables.Config)
		commands.AddLocate(variables.Locate)
		commands.AddRegistry(variables.Registry)
		commands.AddEventbus(variables.Eventbus)
		commands.AddTransport(variables.Transport)

		if err := commands.Run(); err != nil {
			log.Fatal(createFailure, err)
		}

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
