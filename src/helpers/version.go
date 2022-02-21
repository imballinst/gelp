package helpers

import (
	"fmt"
)

func GetVersion() string {
	version := Version
	if version == "" {
		version = "dev"
	}
	return fmt.Sprintf("Current version is: %s\n", version)
}
