package helpers

import (
	"bytes"
	"fmt"
	"os/exec"
	"strings"
)

// This is a helper function to emit a more helpful error code.
func ExecCommand(commandWithArgs string) (string, error) {
	array := []string{}
	segment := ""
	isWhitespaceIgnored := false

	for i := 0; i < len(commandWithArgs); i++ {
		char := string(commandWithArgs[i])

		if char == "'" {
			isWhitespaceIgnored = !isWhitespaceIgnored
		}

		if char == " " && !isWhitespaceIgnored {
			array = append(array, segment)
			segment = ""
			continue
		}

		// Ignore the "'" character.
		if char != "'" {
			segment = segment + char
		}

		if i+1 == len(commandWithArgs) {
			array = append(array, segment)
		}
	}

	cmd := exec.Command(array[0], array[1:]...)
	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr
	err := cmd.Run()

	outString := strings.Trim(out.String(), "\n")
	fmt.Println(array)
	fmt.Println(outString)
	if err != nil {
		// Error logging.
		fmt.Println(fmt.Sprint(err) + ": " + stderr.String())
	}
	return outString, err
}
