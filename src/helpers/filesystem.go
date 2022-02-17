package gelp

import (
	"io/ioutil"
	"os"
)

func WriteToFile(filePath string, content string) error {
	_, err := os.Stat(filePath)
	var f *os.File

	if err != nil {
		// Create file if doesn't exist.
		f, err = os.Create(filePath)
	} else {
		// Open file.
		f, err = os.OpenFile(filePath, os.O_RDWR, 0644)
	}

	// Don't forget to close it later.
	defer f.Close()
	// This is error from creating/opening the file.
	if err != nil {
		return err
	}

	str := []byte(content)
	f.Write(str)

	return nil
}

func ReadFileContent(filePath string) (string, error) {
	var content string

	byteContent, err := ioutil.ReadFile(filePath)
	if err != nil {
		return content, err
	}

	return string(byteContent), nil
}
