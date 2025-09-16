package component

import (
	_ "embed"
)

const (
	Http  = "http"
	PProf = "pprof"
)

const (
	HttpPackage = "github.com/dobyte/due/component/http"
)

var (
	//go:embed http.toml
	HttpTemplate string
	//go:embed pprof.toml
	PProfTemplate string
)
