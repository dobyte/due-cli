package gen

import (
	"os"
	"path/filepath"
	"text/template"
)

type Makefile struct {
	Out       string            // 输出位置
	Tpl       string            // 模板文件
	Replaces  map[string]string // 替换内容
	Variables *Variables        // 替换变量
}

type Variables struct {
	Name            string // 名称
	Lock            string // 分布式锁
	Codec           string // 编码
	Cache           string // 缓存
	Module          string // 模块
	Locate          string // 定位器
	Config          string // 配置中心
	Network         string // 网络
	Registry        string // 注册中心
	Transport       string // 传输层
	Eventbus        string // 事件总线
	GoVersion       string // go版本
	DueFullVersion  string // due版本
	DueMajorVersion string // due主版本
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
		if err := g.doWrite(filepath.Join(g.dir, makefile.Out), makefile.Tpl, makefile.Variables); err != nil {
			return err
		}
	}

	return nil
}

// 写入文件内容
func (g *Generator) doWrite(file string, tpl string, variables *Variables) error {
	if err := os.MkdirAll(filepath.Dir(file), os.ModePerm); err != nil {
		return err
	}

	tmpl, err := template.New(file).Parse(tpl)
	if err != nil {
		return err
	}

	f, err := os.OpenFile(file, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, os.ModePerm)
	if err != nil {
		return err
	}
	defer f.Close()

	return tmpl.Execute(f, variables)
}
