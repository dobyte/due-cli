package base

import (
	_ "embed"
)

const (
	Package = "github.com/dobyte/due"
)

var (
	//go:embed base.toml
	BaseTemplate string
	//go:embed packet.toml
	PacketTemplate string
	//go:embed task.toml
	TaskTemplate string
	//go:embed log.toml
	LogTemplate string
)
