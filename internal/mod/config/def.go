package config

import (
	_ "embed"
)

const (
	File   = "file"
	Consul = "consul"
	Etcd   = "etcd"
	Nacos  = "nacos"
)

const (
	EtcdPackage   = "github.com/dobyte/due/config/etcd"
	NacosPackage  = "github.com/dobyte/due/config/nacos"
	ConsulPackage = "github.com/dobyte/due/config/consul"
)

var (
	//go:embed file.toml
	FileTemplate string
	//go:embed consul.toml
	ConsulTemplate string
	//go:embed etcd.toml
	EtcdTemplate string
	//go:embed nacos.toml
	NacosTemplate string
)
