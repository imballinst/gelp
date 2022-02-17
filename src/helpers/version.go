package gelp

import (
	"os"
)

func EmitVersion() {
	lastRevision, err := ExecCommand("git", "rev-list --tags --max-count=1")
	if err != nil {
		panic(err)
	}

	lastTag, err := ExecCommand("git", "describe --tags "+lastRevision)
	if err != nil {
		panic(err)
	}

	path, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	err = WriteToFile(path+"/version.txt", lastTag)
	if err != nil {
		panic(err)
	}
}
