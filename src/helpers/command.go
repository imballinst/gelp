package helpers

import (
	"bytes"
	"fmt"
	"os/exec"
	"strings"
)

// This is a helper function to emit a more helpful error code.
func ExecCommand(commandWithArgs string) (string, error) {
	array := strings.Split(commandWithArgs, " ")

	cmd := exec.Command(array[0], array[1:]...)
	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr
	err := cmd.Run()

	outString := strings.Trim(out.String(), "\n")
	if err != nil {
		// Error logging.
		fmt.Println(fmt.Sprint(err) + ": " + stderr.String())
	}
	return outString, err
}
