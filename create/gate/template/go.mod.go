package template

const GoModOutput = `go.mod`

const GoModTemplate = `
module ${VarName}

go ${VarGoVersion}

require (
	github.com/dobyte/due/v2 ${VarFrameworkVersion}
	github.com/dobyte/due/locate/${VarLocator}/v2 ${VarFrameworkVersion}
	github.com/dobyte/due/network/${VarNetwork}/v2 ${VarFrameworkVersion}
	github.com/dobyte/due/registry/${VarRegistry}/v2 ${VarFrameworkVersion}
)
`
