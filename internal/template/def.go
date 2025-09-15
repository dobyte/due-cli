package template

import (
	_ "embed"
)

const (
	GoModOutput     = "go.mod"
	GitignoreOutput = ".gitignore"
)

var (
	//go:embed go.mod.tpl
	GoModTemplate string
	//go:embed gitignore.tpl
	GitignoreTemplate string
)
