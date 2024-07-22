package template

const ProtocolOutput = "app/logic/protocol.go"

const ProtocolTemplate = `
package logic

type GreetReq struct {
	Message string ${VarSymbol}json:"message"${VarSymbol}
}

type GreetRes struct {
	Message string ${VarSymbol}json:"message"${VarSymbol}
}
`
