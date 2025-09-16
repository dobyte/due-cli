package template

import (
	_ "embed"
)

const (
	MainOutput = "main.go"
	AppOutput  = "app/app.go"
)

var (
	//go:embed main.go.tpl
	MainTemplate string
	//go:embed app.go.tpl
	AppTemplate string
)
