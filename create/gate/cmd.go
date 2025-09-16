package gate

import (
	"fmt"
	"path/filepath"

	"github.com/AlecAivazis/survey/v2"
	"github.com/dobyte/due-cli/create/gate/template"
	"github.com/dobyte/due-cli/internal/gen"
	"github.com/dobyte/due-cli/internal/log"
	"github.com/dobyte/due-cli/internal/mod"
	"github.com/dobyte/due-cli/internal/mod/cluster"
	"github.com/dobyte/due-cli/internal/os"
	tpl "github.com/dobyte/due-cli/internal/template"
	vs "github.com/dobyte/due-cli/internal/version"
	"github.com/urfave/cli/v2"
)

const (
	createFailure = "create project failure"
	createSuccess = "create project success"
)

var Command = &cli.Command{
	Name:  "gate",
	Usage: "create a new gate project",
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
			Default: "gate",
		}, &variables.Name); err != nil {
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
			Message: "choose a network component:",
			Options: []string{"ws", "tcp"},
			Default: "ws",
		}, &variables.Network); err != nil {
			log.Fatal(createFailure, err)
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
		etc.AddCluster(cluster.Gate)
		etc.AddLocate(variables.Locate)
		etc.AddRegistry(variables.Registry)
		etc.AddNetworkServer(variables.Network)

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

		if isNeedGenMod {
			makefiles = append(makefiles, &gen.Makefile{
				Out:       tpl.GoModOutput,
				Tpl:       tpl.GoModTemplate,
				Variables: variables,
			}, &gen.Makefile{
				Out:       tpl.GitignoreOutput,
				Tpl:       tpl.GitignoreTemplate,
				Variables: variables,
			})
		}

		if err := gen.NewGenerator(filepath.Join(dir, variables.Name)).Make(makefiles...); err != nil {
			log.Fatal(createFailure, err)
		}

		commands.AddLocate(variables.Locate)
		commands.AddNetwork(variables.Network)
		commands.AddRegistry(variables.Registry)

		if err := commands.Run(); err != nil {
			log.Fatal(createFailure, err)
		}

		log.Info(createSuccess)

		return nil
	},
}
