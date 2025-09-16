package template

import (
	_ "embed"
)

const (
	MainOutput = "main.go"
)

var (
	//go:embed main.go.tpl
	MainTemplate string
)
