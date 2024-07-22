package template

// MainOutput 主文件输出位置
const MainOutput = "main.go"

// MainTemplate 主文件模板
const MainTemplate = `
package main

import (
	"github.com/dobyte/due/locate/${VarLocator}/v2"
	"github.com/dobyte/due/network/${VarNetwork}/v2"
	"github.com/dobyte/due/registry/${VarRegistry}/v2"
	"github.com/dobyte/due/v2"
	"github.com/dobyte/due/v2/cluster/gate"
)

func main() {
	// 创建容器
	container := due.NewContainer()
	// 创建网络服务器
	server := ${VarNetwork}.NewServer()
	// 创建用户定位器
	locator := ${VarLocator}.NewLocator()
	// 创建服务注册发现
	registry := ${VarRegistry}.NewRegistry()
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
`
