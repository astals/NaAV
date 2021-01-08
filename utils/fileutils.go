package utils

import (
	"io/ioutil"
	"os"
)

func CopyFile(origin string, destination string) error {
	bytesRead, err := ioutil.ReadFile(origin)
	err = ioutil.WriteFile(destination, bytesRead, 0755)
	return err
}
func FileExists(filename string) (bool, error) {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false, err
	}
	return !info.IsDir(), nil
}
