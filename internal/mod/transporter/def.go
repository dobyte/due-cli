package transporter

import (
	_ "embed"
)

const (
	GRPC = "grpc"
	RPCX = "rpcx"
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
