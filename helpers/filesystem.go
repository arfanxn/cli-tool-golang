package helpers

import (
	"io/ioutil"
	"os"
	"path/filepath"
)

func CreateFile(bytes []byte, destination string) error {
	// Create the directory if it doesn't exist
	err := os.MkdirAll(filepath.Dir(destination), os.ModePerm)
	if err != nil {
		return err
	}
	// Save the file to destination
	return ioutil.WriteFile(destination, bytes, 0644)
}
