package gelp

import (
	"fmt"
	"os"
)

func GetVersion() string {
	version := os.Getenv("version")
	if version == "" {
		version = "dev"
	}
	return fmt.Sprintf("Current version is: %s\n", version)
}
