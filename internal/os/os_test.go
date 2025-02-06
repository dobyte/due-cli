package os_test

import (
	"github.com/dobyte/due-cli/internal/os"
	"testing"
)

func TestIsEmptyDir(t *testing.T) {
	ok, err := os.IsEmptyDir("./gate")

	t.Log(ok, err)
}
