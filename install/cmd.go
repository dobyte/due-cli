package install

import (
	"github.com/dobyte/due-cli/install/gorm"
	"github.com/dobyte/due-cli/install/grpc"
	"github.com/dobyte/due-cli/install/mongo"
	"github.com/dobyte/due-cli/install/proto"
	"github.com/dobyte/due-cli/install/rpcx"
	"github.com/urfave/cli/v2"
)

var Command = &cli.Command{
	Name:        "install",
	Usage:       "install and update the toolchain",
	Description: "install and update the toolchain",
	Subcommands: []*cli.Command{
		grpc.Command,
		rpcx.Command,
		gorm.Command,
		mongo.Command,
		proto.Command,
	},
}
