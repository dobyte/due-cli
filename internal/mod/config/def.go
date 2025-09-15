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
