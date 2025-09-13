package log

import (
	"fmt"
	"log"
	"os/exec"
	"strings"
	"time"
)

func Fatal(title string, detail any) {
	println(log.Fatalf, title, detail)
}

func Error(title string, detail any) {
	println(log.Printf, title, detail)
}

func Info(title string) {
	fmt.Printf("%s %s\n", time.Now().Format("2006/01/02 15:04:05"), title)
}

func println(printf func(format string, v ...any), title string, detail any) {
	if err, ok := detail.(error); ok {
		if e, ok := err.(*exec.ExitError); ok {
			printf("%s:\n%s", title, strings.TrimSuffix(string(e.Stderr), "\n"))
		}
	}

	printf("%s:\n%v", title, detail)
}
