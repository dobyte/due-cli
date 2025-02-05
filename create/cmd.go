package create

import (
	"github.com/dobyte/due-cli/create/gate"
	"github.com/dobyte/due-cli/create/mesh"
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
		mesh.Command,
	},
}
