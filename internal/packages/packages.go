package packages

import (
	"fmt"
	"os/exec"
	"time"
)

var Packages = map[string]string{
	"due":             "github.com/dobyte/due",
	"cache/memcache":  "github.com/dobyte/due/cache/memcache",
	"cache/redis":     "github.com/dobyte/due/cache/redis",
	"component/http":  "github.com/dobyte/due/component/http",
	"config/consul":   "github.com/dobyte/due/config/consul",
	"config/etcd":     "github.com/dobyte/due/config/etcd",
	"config/nacos":    "github.com/dobyte/due/config/nacos",
	"crypto/ecc":      "github.com/dobyte/due/crypto/ecc",
	"crypto/rsa":      "github.com/dobyte/due/crypto/rsa",
	"eventbus/kafka":  "github.com/dobyte/due/eventbus/kafka",
	"eventbus/nats":   "github.com/dobyte/due/eventbus/nats",
	"eventbus/redis":  "github.com/dobyte/due/eventbus/redis",
	"locate/redis":    "github.com/dobyte/due/locate/redis",
	"lock/memcache":   "github.com/dobyte/due/lock/memcache",
	"lock/redis":      "github.com/dobyte/due/lock/redis",
	"log/aliyun":      "github.com/dobyte/due/log/aliyun",
	"log/tencent":     "github.com/dobyte/due/log/tencent",
	"network/kcp":     "github.com/dobyte/due/network/kcp",
	"network/tcp":     "github.com/dobyte/due/network/tcp",
	"network/ws":      "github.com/dobyte/due/network/ws",
	"registry/consul": "github.com/dobyte/due/registry/consul",
	"registry/etcd":   "github.com/dobyte/due/registry/etcd",
	"registry/nacos":  "github.com/dobyte/due/registry/nacos",
	"transport/grpc":  "github.com/dobyte/due/transport/grpc",
	"transport/rpcx":  "github.com/dobyte/due/transport/rpcx",
}

func CMD(pkg, major, version, dir string) *exec.Cmd {
	cmd := exec.Command("go", "get", fmt.Sprintf("%s/%s@%s", pkg, major, version))
	cmd.Dir = dir
	cmd.WaitDelay = 30 * time.Second

	return cmd
}
