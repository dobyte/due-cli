package template

const AppOutput = "app/app.go"

const AppTemplate = `
package app

import (
	"${VarName}/app/logic"
	"github.com/dobyte/due/v2/cluster/node"
)

func Init(proxy *node.Proxy) {
	logic.NewCore(proxy).Init()
}
`
