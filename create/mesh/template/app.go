package template

const AppOutput = "app/app.go"

const AppTemplate = `
package app

import "github.com/dobyte/due/v2/cluster/mesh"

func Init(proxy *mesh.Proxy) {
	// TODO: init service
}
`
