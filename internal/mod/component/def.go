package component

import (
	_ "embed"
)

const (
	Http  = "http"
	PProf = "pprof"
)

var (
	//go:embed http.toml
	HttpTemplate string
	//go:embed pprof.toml
	PProfTemplate string
)
