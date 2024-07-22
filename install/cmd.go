package install

import (
	"github.com/dobyte/due/cmd/v2/install/gorm"
	"github.com/dobyte/due/cmd/v2/install/grpc"
	"github.com/dobyte/due/cmd/v2/install/mongo"
	"github.com/dobyte/due/cmd/v2/install/proto"
	"github.com/dobyte/due/cmd/v2/install/rpcx"
	"github.com/urfave/cli/v2"
)

var Command = &cli.Command{
	Name:        "install",
	Usage:       "Install and update the toolchain",
	Description: "Install and update the toolchain",
	Subcommands: []*cli.Command{
		grpc.Command,
		rpcx.Command,
		gorm.Command,
		mongo.Command,
		proto.Command,
	},
}
