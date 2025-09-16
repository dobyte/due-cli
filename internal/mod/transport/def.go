package transport

import (
	_ "embed"
)

const (
	GRPC = "grpc"
	RPCX = "rpcx"
)

const (
	GRPCPackage = "github.com/dobyte/due/transport/grpc"
	RPCXPackage = "github.com/dobyte/due/transport/rpcx"
)

var (
	//go:embed grpc_client.toml
	GRPCClientTemplate string
	//go:embed grpc_server.toml
	GRPCServerTemplate string
	//go:embed rpcx_client.toml
	RPCXClientTemplate string
	//go:embed rpcx_server.toml
	RPCXServerTemplate string
)
