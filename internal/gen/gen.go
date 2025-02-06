package gen

import (
	"github.com/dobyte/due-cli/internal/os"
	tpl "github.com/dobyte/due-cli/internal/template"
)

// MakeGlobalGoModFile 生成全局的Go Mod 文件
func MakeGlobalGoModFile(replaces map[string]string) error {
	dir := "./"

	if os.IsFile(dir + tpl.GoModOutput) {
		return nil
	}

	return NewGenerator(dir).Make(&Makefile{
		Out:      tpl.GoModOutput,
		Tpl:      tpl.GoModTemplate,
		Replaces: replaces,
	})
}
