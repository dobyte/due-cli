package registry

import (
	_ "embed"
)

const (
	Consul = "consul"
	Etcd   = "etcd"
	Nacos  = "nacos"
)

const (
	EtcdPackage   = "github.com/dobyte/due/registry/etcd"
	NacosPackage  = "github.com/dobyte/due/registry/nacos"
	ConsulPackage = "github.com/dobyte/due/registry/consul"
)

var (
	//go:embed consul.toml
	ConsulTemplate string
	//go:embed etcd.toml
	EtcdTemplate string
	//go:embed nacos.toml
	NacosTemplate string
)
