package template

import (
	_ "embed"
)

// MainOutput 主文件输出位置
const MainOutput = "main.go"

// MainTemplate 主文件模板
//
//go:embed main.go.tpl
var MainTemplate string
