package cluster

import (
	_ "embed"
)

const (
	Gate   = "gate"
	Node   = "node"
	Mesh   = "mesh"
	Client = "client"
)

var (
	//go:embed gate.toml
	GateTemplate string
	//go:embed node.toml
	NodeTemplate string
	//go:embed mesh.toml
	MeshTemplate string
	//go:embed client.toml
	ClientTemplate string
)
