package helpers

import (
	"fmt"
)

func DoAndLog(label string, command string) (string, error) {
	Log(label, command)
	return ExecCommand(command)
}

func Log(label string, text string) {
	fmt.Printf("gelp %s: %s\n", label, text)
}
