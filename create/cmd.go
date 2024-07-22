package create

import (
	"github.com/dobyte/due/cmd/v2/create/gate"
	"github.com/dobyte/due/cmd/v2/create/mesh"
	"github.com/dobyte/due/cmd/v2/create/node"
	"github.com/urfave/cli/v2"
)

var Command = &cli.Command{
	Name:        "create",
	Usage:       "Create a new project",
	Description: "Create a new project",
	Subcommands: []*cli.Command{
		gate.Command,
		node.Command,
		mesh.Command,
	},
}
