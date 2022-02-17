package main

import (
	"os"

	cmd "github.com/imballinst/gelp/src/cmd"
)

var version string

func main() {
	os.Setenv("version", version)
	cmd.Execute()
}
