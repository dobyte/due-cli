package mod

import (
	"fmt"
	"os/exec"
	"strings"
	"time"

	"github.com/dobyte/due-cli/internal/mod/base"
	"github.com/dobyte/due-cli/internal/mod/component"
	"github.com/dobyte/due-cli/internal/mod/config"
	"github.com/dobyte/due-cli/internal/mod/locator"
	"github.com/dobyte/due-cli/internal/mod/lock"
	"github.com/dobyte/due-cli/internal/mod/network"
	"github.com/dobyte/due-cli/internal/mod/registry"
	"github.com/dobyte/due-cli/internal/mod/transporter"
)

const defaultWaitDelay = 30 * time.Second

type Commands struct {
	dir       string
	sha       string
	full      string
	major     string
	commands  []*exec.Cmd
	waitDelay time.Duration
}

func NewCommands(dir string, version string, sha string) *Commands {
	c := &Commands{}
	c.dir = dir
	c.sha = sha
	c.full = version
	c.major = strings.Split(version, ".")[0]
	c.commands = make([]*exec.Cmd, 0, 1)
	c.waitDelay = defaultWaitDelay

	return c
}

// AddBase 添加主包
func (c *Commands) AddBase() {
	c.addPackage(base.Package, c.full)
}

// AddLocator 添加定位器
func (c *Commands) AddLocator(name string) *Commands {
	switch name {
	case locator.Redis:
		c.addSubPackage(locator.RedisPackage)
	default:
		// ignore
	}

	return c
}

// AddRegistry 添加服务注册发现配置
func (c *Commands) AddRegistry(name string) *Commands {
	switch name {
	case registry.Etcd:
		c.addSubPackage(registry.EtcdPackage)
	case registry.Nacos:
		c.addSubPackage(registry.NacosPackage)
	case registry.Consul:
		c.addSubPackage(registry.ConsulPackage)
	default:
		// ignore
	}

	return c
}

// AddTransporter 添加 transporter
func (c *Commands) AddTransporter(name string) *Commands {
	switch name {
	case transporter.GRPC:
		c.addSubPackage(transporter.GRPCPackage)
	case transporter.RPCX:
		c.addSubPackage(transporter.RPCXPackage)
	default:
		// ignore
	}

	return c
}

// AddNetwork 添加网络
func (c *Commands) AddNetwork(name string) *Commands {
	switch name {
	case network.WS:
		c.addSubPackage(network.WSPackage)
	case network.TCP:
		c.addSubPackage(network.TCPPackage)
	default:
		// ignore
	}

	return c
}

// AddLock 添加锁
func (c *Commands) AddLock(name string) *Commands {
	switch name {
	case lock.Redis:
		c.addSubPackage(lock.RedisPackage)
	case lock.Memcache:
		c.addSubPackage(lock.MemcachePackage)
	default:
		// ignore
	}

	return c
}

// AddConfig 添加配置
func (c *Commands) AddConfig(name string) *Commands {
	switch name {
	case config.Etcd:
		c.addSubPackage(config.EtcdPackage)
	case config.Nacos:
		c.addSubPackage(config.NacosPackage)
	case config.Consul:
		c.addSubPackage(config.ConsulPackage)
	default:
		// ignore
	}

	return c
}

// AddComponent 添加组件
func (c *Commands) AddComponent(name string) *Commands {
	switch name {
	case component.Http:
		c.addSubPackage(component.HttpPackage)
	default:
		// ignore
	}

	return c
}

// Run 运行命令
func (c *Commands) Run() error {
	c.addGoModTidy()

	for _, cmd := range c.commands {
		if _, err := cmd.Output(); err != nil {
			return err
		}
	}

	return nil
}

// 添加子包
func (c *Commands) addSubPackage(pkg string) {
	c.addPackage(pkg, c.sha)
}

// 添加包命令
func (c *Commands) addPackage(pkg, sha string) {
	cmd := exec.Command("go", "get", fmt.Sprintf("%s/%s@%s", pkg, c.major, sha))
	cmd.Dir = c.dir
	cmd.WaitDelay = c.waitDelay

	c.commands = append(c.commands, cmd)
}

// 添加go mod tidy
func (c *Commands) addGoModTidy() *Commands {
	cmd := exec.Command("go", "mod", "tidy")
	cmd.Dir = c.dir
	cmd.WaitDelay = c.waitDelay

	c.commands = append(c.commands, cmd)
	return c
}
