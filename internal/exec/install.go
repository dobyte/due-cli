package exec

import (
	"fmt"
	"os/exec"
	"strings"
	"sync"
)

type Package struct {
	Name    string
	Module  string
	Version string
}

// Install 安装包
func Install(packages ...Package) {
	var wg sync.WaitGroup

	wg.Add(len(packages))

	for _, item := range packages {
		go func(p Package) {
			if _, err := exec.Command("go", "install", fmt.Sprintf("%s@%s", p.Module, p.Version)).Output(); err != nil {
				fmt.Printf("%s installed failed: %v\n", p.Name, strings.TrimSuffix(string(err.(*exec.ExitError).Stderr), "\n"))
			} else {
				fmt.Printf("%s installed successfully.\n", p.Name)
			}

			wg.Done()
		}(item)
	}

	wg.Wait()
}
