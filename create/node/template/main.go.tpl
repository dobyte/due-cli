package main

import (
	"game-server/hall/app"

	{{- if ne .Config "none"}}
	cc "github.com/dobyte/due/config/{{.Config}}/{{.DueMajorVersion}}"
	{{- end}}
	"github.com/dobyte/due/locate/{{.Locate}}/{{.DueMajorVersion}}"
	"github.com/dobyte/due/registry/{{.Registry}}/{{.DueMajorVersion}}"
	{{- if ne .Transport "none"}}
	"github.com/dobyte/due/transport/{{.Transport}}/{{.DueMajorVersion}}"
	{{- end}}
	"github.com/dobyte/due/{{.DueMajorVersion}}"
	"github.com/dobyte/due/{{.DueMajorVersion}}/cluster/node"
	{{- if ne .Config "none"}}
	"github.com/dobyte/due/v2/config"
	{{- end}}
)

func main() {
	{{- if ne .Config "none"}}
	// 设置配置中心
	config.SetConfigurator(config.NewConfigurator(config.WithSources(cc.NewSource())))
	{{- end}}
	// 创建容器
	container := due.NewContainer()
	// 创建用户定位器
	locator := {{.Locate}}.NewLocator()
	// 创建服务注册发现
	registry := {{.Registry}}.NewRegistry()
	{{- if ne .Transport "none"}}
	// 创建RPC传输器
	transporter := {{.Transport}}.NewTransporter()
	{{- end}}
	// 创建节点组件
	component := node.NewNode(
		node.WithLocator(locator),
		node.WithRegistry(registry),
		{{- if ne .Transport "none"}}
		node.WithTransporter(transporter),
		{{- end}}
	)
	// 初始化应用
	app.Init(component.Proxy())
	// 添加节点组件
	container.Add(component)
	// 启动容器
	container.Serve()
}
