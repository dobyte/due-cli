package template

import (
	_ "embed"
)

const MainOutput = "main.go"

//go:embed main.go.tpl
var MainTemplate string
