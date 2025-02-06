package version

import (
	"os/exec"
	"regexp"
)

var (
	GoVersion   = "1.22"  // Go版本
	ToolVersion = "1.0.0" // 工具版本
)

func init() {
	b, err := exec.Command("go", "version").Output()
	if err != nil {
		return
	}

	matches := regexp.MustCompile(`go version go(\d+\.\d+(?:\.\d+)?)`).FindStringSubmatch(string(b))

	if len(matches) > 0 {
		GoVersion = matches[1]
	}
}
