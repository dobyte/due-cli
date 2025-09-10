package gen

import (
	"github.com/dobyte/due-cli/internal/os"
	tpl "github.com/dobyte/due-cli/internal/template"
)

// MakeGlobalFile 生成全局文件
func MakeGlobalFile(replaces map[string]string) error {
	dir := "./"

	makefiles := make([]*Makefile, 0, 2)

	if !os.IsFile(dir + tpl.GoModOutput) {
		makefiles = append(makefiles, &Makefile{
			Out:      tpl.GoModOutput,
			Tpl:      tpl.GoModTemplate,
			Replaces: replaces,
		})
	}

	if !os.IsFile(dir + tpl.GitignoreOutput) {
		makefiles = append(makefiles, &Makefile{
			Out:      tpl.GitignoreOutput,
			Tpl:      tpl.GitignoreTemplate,
			Replaces: replaces,
		})
	}

	return NewGenerator(dir).Make(makefiles...)
}
