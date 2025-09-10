package gate

import (
	"github.com/dobyte/due-cli/create/gate/template"
	"github.com/dobyte/due-cli/internal/etc"
	"github.com/dobyte/due-cli/internal/etc/cluster"
	"github.com/dobyte/due-cli/internal/gen"
	"path"
)

// Make 生成网关示例项目
func Make(dir, locator, registry, network string) error {
	e := etc.NewEtc()
	e.AddLog()
	e.AddPacket()
	e.AddCluster(cluster.Gate)
	e.AddLocator(locator)
	e.AddRegistry(registry)
	e.AddNetworkServer(network)

	replaces := map[string]string{
		"VarNetwork":  network,
		"VarLocator":  locator,
		"VarRegistry": registry,
	}

	return gen.NewGenerator(path.Join(dir, cluster.Gate)).Make(&gen.Makefile{
		Out:      template.MainOutput,
		Tpl:      template.MainTemplate,
		Replaces: replaces,
	}, &gen.Makefile{
		Out:      e.Output(),
		Tpl:      e.Template(),
		Replaces: replaces,
	})
}
