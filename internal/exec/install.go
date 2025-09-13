package exec

import (
	"fmt"
	"os/exec"
	"sync"

	"github.com/dobyte/due-cli/internal/log"
)

type Package struct {
	Name    string
	Module  string
	Version string
}

const (
	installFailure = "install failure"
	installSuccess = "install success"
)

// Install 安装包
func Install(packages ...Package) {
	var wg sync.WaitGroup

	wg.Add(len(packages))

	for _, item := range packages {
		go func(p Package) {
			if _, err := exec.Command("go", "install", fmt.Sprintf("%s@%s", p.Module, p.Version)).Output(); err != nil {
				log.Error(installFailure, err)
			} else {
				log.Info(installSuccess)
			}

			wg.Done()
		}(item)
	}

	wg.Wait()
}
