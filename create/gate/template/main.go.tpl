package main

import (
    {{- if eq .Network "ws"}}
    "net/http"
    {{- end}}
	"github.com/dobyte/due/locate/{{.Locate}}/{{.DueMajorVersion}}"
	"github.com/dobyte/due/network/{{.Network}}/{{.DueMajorVersion}}"
	"github.com/dobyte/due/registry/{{.Registry}}/{{.DueMajorVersion}}"
	"github.com/dobyte/due/{{.DueMajorVersion}}"
	"github.com/dobyte/due/{{.DueMajorVersion}}/cluster/gate"
)

func main() {
	// 创建容器
	container := due.NewContainer()
    // 创建网络服务器
    server := {{.Network}}.NewServer()
    {{- if eq .Network "ws"}}
    // 监听HTTP连接升级WS协议
    server.OnUpgrade(func(w http.ResponseWriter, r *http.Request) bool {
        return true
    })
    {{- end}}
	// 创建用户定位器
	locator := {{.Locate}}.NewLocator()
	// 创建服务注册发现
	registry := {{.Registry}}.NewRegistry()
	// 创建网关组件
	component := gate.NewGate(
		gate.WithServer(server),
		gate.WithLocator(locator),
		gate.WithRegistry(registry),
	)
	// 添加网关组件
	container.Add(component)
	// 启动容器
	container.Serve()
}