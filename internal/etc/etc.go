package etc

import (
	"github.com/dobyte/due-cli/internal/etc/base"
	"github.com/dobyte/due-cli/internal/etc/cluster"
	"github.com/dobyte/due-cli/internal/etc/locator"
	"github.com/dobyte/due-cli/internal/etc/network"
	"github.com/dobyte/due-cli/internal/etc/registry"
	"github.com/dobyte/due-cli/internal/etc/transporter"
	"strings"
)

type Etc struct {
	templates []string
}

func NewEtc() *Etc {
	e := &Etc{}
	e.templates = append(e.templates, base.Template)
	return e
}

// Output 输出目录
func (e *Etc) Output() string {
	return "etc/etc.toml"
}

// Template 模板
func (e *Etc) Template() string {
	return strings.Join(e.templates, "\n")
}

// AddLog 添加日志配置
func (e *Etc) AddLog() *Etc {
	e.templates = append(e.templates, base.LogTemplate)

	return e
}

// AddTask 添加任务配置
func (e *Etc) AddTask() *Etc {
	e.templates = append(e.templates, base.TaskTemplate)

	return e
}

// AddPacket 添加打包配置
func (e *Etc) AddPacket() *Etc {
	e.templates = append(e.templates, base.PacketTemplate)

	return e
}

// AddLocator 添加定位器配置
func (e *Etc) AddLocator(name string) *Etc {
	switch name {
	case locator.Redis:
		e.templates = append(e.templates, locator.RedisTemplate)
	default:
		// ignore
	}

	return e
}

// AddRegistry 添加服务注册发现配置
func (e *Etc) AddRegistry(name string) *Etc {
	switch name {
	case registry.Consul:
		e.templates = append(e.templates, registry.ConsulTemplate)
	case registry.Etcd:
		e.templates = append(e.templates, registry.EtcdTemplate)
	case registry.Nacos:
		e.templates = append(e.templates, registry.NacosTemplate)
	default:
		// ignore
	}

	return e
}

// AddCluster 添加集群配置
func (e *Etc) AddCluster(name string) *Etc {
	switch name {
	case cluster.Gate:
		e.templates = append(e.templates, cluster.GateTemplate)
	case cluster.Node:
		e.templates = append(e.templates, cluster.NodeTemplate)
	case cluster.Mesh:
		e.templates = append(e.templates, cluster.MeshTemplate)
	case cluster.Client:
		e.templates = append(e.templates, cluster.ClientTemplate)
	default:
		// ignore
	}

	return e
}

// AddTransportServer 添加传输服务器配置
func (e *Etc) AddTransportServer(name string) *Etc {
	switch name {
	case transporter.RPCX:
		e.templates = append(e.templates, transporter.RPCXServerTemplate)
	case transporter.GRPC:
		e.templates = append(e.templates, transporter.GRPCServerTemplate)
	default:
		// ignore
	}

	return e
}

// AddTransportClient 添加传输客户端配置
func (e *Etc) AddTransportClient(name string) *Etc {
	switch name {
	case transporter.RPCX:
		e.templates = append(e.templates, transporter.RPCXClientTemplate)
	case transporter.GRPC:
		e.templates = append(e.templates, transporter.GRPCClientTemplate)
	default:
		// ignore
	}

	return e
}

// AddNetworkServer 添加网络服务器配置
func (e *Etc) AddNetworkServer(name string) *Etc {
	switch name {
	case network.WS:
		e.templates = append(e.templates, network.WSServerTemplate)
	case network.TCP:
		e.templates = append(e.templates, network.TCPServerTemplate)
	default:
		// ignore
	}

	return e
}
