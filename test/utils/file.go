package utils

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

func FileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

func ReadFile(path string) ([]byte, error) {
	file, openErr := os.Open(path)
	if openErr != nil {
		return []byte{}, openErr
	}
	defer file.Close()
	bytes, _ := ioutil.ReadAll(file)
	return bytes
}

func WriteFile(path string, bytes []byte) error {
	return ioutil.WriteFile(path, bytes, 0644)
}

func RemoveFile(path string) error {
	return os.Remove(path)
}