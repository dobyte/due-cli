package template

// MainOutput 主文件输出位置
const MainOutput = "main.go"

// MainTemplate 主文件模板
const MainTemplate = `
package main

import (
	"${VarName}/app"
	"github.com/dobyte/due/locate/${VarLocator}/v2"
	"github.com/dobyte/due/registry/${VarRegistry}/v2"
	"github.com/dobyte/due/transport/${VarTransporter}/v2"
	"github.com/dobyte/due/v2"
	"github.com/dobyte/due/v2/cluster/mesh"
)

func main() {
	// 创建容器
	container := due.NewContainer()
	// 创建用户定位器
	locator := ${VarLocator}.NewLocator()
	// 创建服务注册发现
	registry := ${VarRegistry}.NewRegistry()
	// 创建RPC传输器
	transporter := ${VarTransporter}.NewTransporter()
	// 创建网格组件
	component := mesh.NewMesh(
		mesh.WithLocator(locator),
		mesh.WithRegistry(registry),
		mesh.WithTransporter(transporter),
	)
	// 初始化应用
	app.Init(component.Proxy())
	// 添加网格组件
	container.Add(component)
	// 启动容器
	container.Serve()
}
`
