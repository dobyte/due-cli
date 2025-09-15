package mod

import (
	"fmt"
	"os/exec"
	"strings"
	"time"

	"github.com/dobyte/due-cli/internal/mod/locator"
	"github.com/dobyte/due-cli/internal/mod/registry"
)

const defaultWaitDelay = 30 * time.Second

type Commands struct {
	dir          string
	sha          string
	waitDelay    time.Duration
	fullVersion  string
	majorVersion string
	commands     []*exec.Cmd
}

func NewCmd(dir string, version string, majorVersion string, sha string) *Commands {
	return &Commands{
		dir:          dir,
		sha:          sha,
		waitDelay:    defaultWaitDelay,
		fullVersion:  version,
		majorVersion: strings.Split(version, ".")[0],
	}
}

// AddLocator 添加定位器
func (c *Commands) AddLocator(name string) *Commands {
	var pkg string

	switch name {
	case locator.Redis:
		pkg = locator.RedisPackage
	default:
		// ignore
	}

	if pkg != "" {
		cmd := exec.Command("go", "get", fmt.Sprintf("%s/%s@%s", pkg, c.fullVersion, c.sha))
		cmd.Dir = c.dir
		cmd.WaitDelay = c.waitDelay

		c.commands = append(c.commands, cmd)
	}

	return c
}

// AddRegistry 添加服务注册发现配置
func (c *Commands) AddRegistry(name string) *Commands {
	var pkg string

	switch name {
	case registry.Consul:
		pkg = registry.ConsulPackage
	case registry.Etcd:
		pkg = registry.EtcdPackage
	case registry.Nacos:
		pkg = registry.NacosPackage
	default:
		// ignore
	}

	if pkg != "" {
		cmd := exec.Command("go", "get", fmt.Sprintf("%s/%s@%s", pkg, c.fullVersion, c.sha))
		cmd.Dir = c.dir
		cmd.WaitDelay = c.waitDelay

		c.commands = append(c.commands, cmd)
	}

	return c
}
