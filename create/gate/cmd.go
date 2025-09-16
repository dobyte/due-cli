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
			dir        string
			name       string
			alone      bool
			module     string
			goVersion  string
			dueVersion string
			network    string
			locator    string
			registry   string
		)

		if err := survey.AskOne(&survey.Input{
			Message: "specify the project directory:",
			Default: "./testdata",
		}, &dir); err != nil {
			log.Fatal(createFailure, err)
		}

		if err := survey.AskOne(&survey.Input{
			Message: "specify the project name:",
			Default: "gate",
		}, &name); err != nil {
			log.Fatal(createFailure, err)
		}

		if err := survey.AskOne(&survey.Confirm{
			Message: "whether it is an alone project:",
			Default: false,
		}, &alone); err != nil {
			log.Fatal(createFailure, err)
		}

		isNeedGenMod := alone || !os.IsFile(filepath.Join(dir, "go.mod"))

		if isNeedGenMod {
			if err := survey.AskOne(&survey.Input{
				Message: "specify the project module:",
				Default: fmt.Sprintf("github.com/due/%s", name),
			}, &module); err != nil {
				log.Fatal(createFailure, err)
			}

			if err := survey.AskOne(&survey.Input{
				Message: "specify the go version:",
			}, &goVersion); err != nil {
				log.Fatal(createFailure, err)
			}

			if err := survey.AskOne(&survey.Input{
				Message: "specify the due version:",
				Default: "latest",
			}, &dueVersion); err != nil {
				log.Fatal(createFailure, err)
			}
		}

		if err := survey.AskOne(&survey.Select{
			Message: "choose a network component:",
			Options: []string{"ws", "tcp"},
			Default: "ws",
		}, &network); err != nil {
			log.Fatal(createFailure, err)
		}

		if err := survey.AskOne(&survey.Select{
			Message: "choose a locator component:",
			Options: []string{"redis"},
			Default: "redis",
		}, &locator); err != nil {
			log.Fatal(createFailure, err)
		}

		if err := survey.AskOne(&survey.Select{
			Message: "choose a registry component:",
			Options: []string{"nacos", "etcd", "consul"},
			Default: "nacos",
		}, &registry); err != nil {
			log.Fatal(createFailure, err)
		}

		if ok, err := os.IsEmptyDir(filepath.Join(dir, name)); err != nil {
			log.Fatal(createFailure, err)
		} else if !ok {
			log.Fatal(createFailure, "the project dir is not empty.")
		}

		var (
			commands  *mod.Commands
			variables = &gen.Variables{
				Name:      name,
				Module:    module,
				Network:   network,
				Locator:   locator,
				Registry:  registry,
				GoVersion: goVersion,
			}
		)

		if isNeedGenMod {
			if module == "" {
				log.Fatal(createFailure, "the module of the project is required.")
			}

			if goVersion == "" {
				if v, err := vs.ParseGoVersion(); err != nil {
					log.Fatal(createFailure, err)
				} else {
					goVersion = v
				}
			}

			if dueVersion == "" {
				dueVersion = "latest"
			}

			full, major, sha, err := vs.ParseDueVersion(dueVersion)
			if err != nil {
				log.Fatal(createFailure, err)
			}

			variables.GoVersion = goVersion
			variables.DueFullVersion = full
			variables.DueMajorVersion = major

			if alone {
				commands = mod.NewCommands(filepath.Join(dir, name), full, sha)
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
		etc.AddLocator(locator)
		etc.AddRegistry(registry)
		etc.AddNetworkServer(network)

		makefiles := make([]*gen.Makefile, 0, 3)
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

		if err := gen.NewGenerator(filepath.Join(dir, name)).Make(makefiles...); err != nil {
			log.Fatal(createFailure, err)
		}

		commands.AddLock(locator)
		commands.AddRegistry(registry)
		commands.AddNetwork(network)

		if err := commands.Run(); err != nil {
			log.Fatal(createFailure, err)
		}

		log.Info(createSuccess)

		return nil
	},
}
