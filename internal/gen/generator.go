package gen

import (
	"os"
	"path/filepath"
	"strings"
)

type Makefile struct {
	Out      string            // 输出位置
	Tpl      string            // 模板文件
	Replaces map[string]string // 替换内容
}

type Generator struct {
	dir string // 项目目录
}

func NewGenerator(dir string) *Generator {
	return &Generator{dir: dir}
}

// Make 生成文件
func (g *Generator) Make(makefiles ...*Makefile) error {
	for _, makefile := range makefiles {
		if err := g.doWrite(filepath.Join(g.dir, makefile.Out), makefile.Tpl, makefile.Replaces); err != nil {
			return err
		}
	}

	return nil
}

// 写入文件内容
func (g *Generator) doWrite(file string, tpl string, replaces map[string]string) error {
	s := os.Expand(tpl, func(s string) string {
		switch {
		case len(s) >= 3 && s[:3] == "Var":
			return replaces[s]
		case len(s) >= 6 && s[:6] == "Symbol":
			return replaces[s]
		default:
			return "$" + s
		}
	})

	if err := os.MkdirAll(filepath.Dir(file), os.ModePerm); err != nil {
		return err
	}

	return os.WriteFile(file, []byte(strings.TrimPrefix(s, "\n")), os.ModePerm)
}
