package mod

import (
	"strings"

	"github.com/dobyte/due-cli/internal/mod/base"
	"github.com/dobyte/due-cli/internal/mod/cache"
	"github.com/dobyte/due-cli/internal/mod/cluster"
	"github.com/dobyte/due-cli/internal/mod/component"
	"github.com/dobyte/due-cli/internal/mod/config"
	"github.com/dobyte/due-cli/internal/mod/eventbus"
	"github.com/dobyte/due-cli/internal/mod/locate"
	"github.com/dobyte/due-cli/internal/mod/lock"
	"github.com/dobyte/due-cli/internal/mod/network"
	"github.com/dobyte/due-cli/internal/mod/registry"
	"github.com/dobyte/due-cli/internal/mod/transport"
)

type Etc struct {
	templates []string
}

func NewEtc() *Etc {
	e := &Etc{}
	e.templates = append(e.templates, base.BaseTemplate)
	return e
}

// Output 输出目录
func (e *Etc) Output() string {
	return "etc/etc.toml"
}

// Template 模板
func (e *Etc) Template() string {
	return strings.Join(e.templates, "\n\n")
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

// AddLocate 添加定位组件配置
func (e *Etc) AddLocate(name string) *Etc {
	switch name {
	case locate.Redis:
		e.templates = append(e.templates, locate.RedisTemplate)
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

// AddConfig 添加配置中心配置
func (e *Etc) AddConfig(name string) *Etc {
	switch name {
	case config.File:
		e.templates = append(e.templates, config.FileTemplate)
	case config.Consul:
		e.templates = append(e.templates, config.ConsulTemplate)
	case config.Etcd:
		e.templates = append(e.templates, config.EtcdTemplate)
	case config.Nacos:
		e.templates = append(e.templates, config.NacosTemplate)
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

// AddCache 添加缓存配置
func (e *Etc) AddCache(name string) *Etc {
	switch name {
	case cache.Redis:
		e.templates = append(e.templates, cache.RedisTemplate)
	case cache.Memcache:
		e.templates = append(e.templates, cache.MemcacheTemplate)
	default:
		// ignore
	}

	return e
}

// AddLock 添加分布式锁配置
func (e *Etc) AddLock(name string) *Etc {
	switch name {
	case lock.Redis:
		e.templates = append(e.templates, lock.RedisTemplate)
	case lock.Memcache:
		e.templates = append(e.templates, lock.MemcacheTemplate)
	default:
		// ignore
	}

	return e
}

// AddTransportServer 添加传输服务器配置
func (e *Etc) AddTransportServer(name string) *Etc {
	switch name {
	case transport.RPCX:
		e.templates = append(e.templates, transport.RPCXServerTemplate)
	case transport.GRPC:
		e.templates = append(e.templates, transport.GRPCServerTemplate)
	default:
		// ignore
	}

	return e
}

// AddTransportClient 添加传输客户端配置
func (e *Etc) AddTransportClient(name string) *Etc {
	switch name {
	case transport.RPCX:
		e.templates = append(e.templates, transport.RPCXClientTemplate)
	case transport.GRPC:
		e.templates = append(e.templates, transport.GRPCClientTemplate)
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

// AddNetworkClient 添加网络客户端配置
func (e *Etc) AddNetworkClient(name string) *Etc {
	switch name {
	case network.WS:
		e.templates = append(e.templates, network.WSClientTemplate)
	case network.TCP:
		e.templates = append(e.templates, network.TCPClientTemplate)
	default:
		// ignore
	}

	return e
}

// AddEventbus 添加事件总线配置
func (e *Etc) AddEventbus(name string) *Etc {
	switch name {
	case eventbus.Nats:
		e.templates = append(e.templates, eventbus.NatsTemplate)
	case eventbus.Redis:
		e.templates = append(e.templates, eventbus.RedisTemplate)
	case eventbus.Kafka:
		e.templates = append(e.templates, eventbus.KafkaTemplate)
	default:
		// ignore
	}

	return e
}

// AddComponent 添加组件配置
func (e *Etc) AddComponent(name string) *Etc {
	switch name {
	case component.Http:
		e.templates = append(e.templates, component.HttpTemplate)
	case component.PProf:
		e.templates = append(e.templates, component.PProfTemplate)
	default:
		// ignore
	}

	return e
}
