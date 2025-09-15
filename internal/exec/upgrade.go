package exec

import (
	"fmt"
	"os/exec"
	"path/filepath"
	"regexp"
	"time"

	"github.com/dobyte/due-cli/internal/log"
	"github.com/dobyte/due-cli/internal/os"
	"github.com/dobyte/due-cli/internal/version"
)

const (
	dueFrameworkPrefix  = "github.com/dobyte/due"
	dueFrameworkTagsUrl = "https://api.github.com/repos/dobyte/due/git/refs/tags"
)

const (
	upgradeFailure = "upgrade failure"
	upgradeSuccess = "upgrade success"
)

type tag struct {
	Status string `json:"status,omitempty"`
	Ref    string `json:"ref,omitempty"`
	Object struct {
		Sha string `json:"sha,omitempty"`
	} `json:"object,omitempty"`
}

// Upgrade 升级框架
func Upgrade(dir string, v string) {
	if !os.IsDir(dir) {
		log.Fatal(upgradeFailure, "the dir is not a directory")
	}

	mod := filepath.Join(dir, "go.mod")

	if !os.IsFile(mod) {
		log.Fatal(upgradeFailure, "the go.mod file does not exist")
	}

	content, err := os.ReadFile(mod)
	if err != nil {
		log.Fatal(upgradeFailure, err)
	}

	full, major, sha, err := version.ParseDueVersion(v)
	if err != nil {
		log.Fatal(upgradeFailure, err)
	}

	pkg := fmt.Sprintf("%s/%s@%s", dueFrameworkPrefix, major, full)
	cmd := exec.Command("go", "get", pkg)
	cmd.Dir = dir
	cmd.WaitDelay = 30 * time.Second

	if _, err = cmd.Output(); err != nil {
		log.Fatal(upgradeFailure, err)
	}

	reg := regexp.MustCompile(fmt.Sprintf(`(%s/\w+/\w+)/v\d+`, dueFrameworkPrefix))
	rst := reg.FindAllStringSubmatch(string(content), -1)

	for _, group := range rst {
		if len(group) != 2 {
			continue
		}

		pkg = fmt.Sprintf("%s/%s@%s", group[1], major, sha)
		cmd = exec.Command("go", "get", pkg)
		cmd.Dir = dir
		cmd.WaitDelay = 30 * time.Second

		if _, err = cmd.Output(); err != nil {
			log.Fatal(upgradeFailure, err)
		}
	}

	log.Info(upgradeSuccess)
}
