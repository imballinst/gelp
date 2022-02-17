package gelp

import (
	"bytes"
	"fmt"
	"os/exec"
)

// This is a helper function to emit a more helpful error code.
func ExecCommand(command string, arg ...string) (string, error) {
	cmd := exec.Command(command, arg...)
	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr
	err := cmd.Run()

	outString := out.String()
	if err != nil {
		// Error logging.
		fmt.Println(fmt.Sprint(err) + ": " + stderr.String())
	}
	return outString, err
}
