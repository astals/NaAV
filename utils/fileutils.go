package utils

import (
	"fmt"
	"io/ioutil"
	"os"
)

//driverVerboseName= strings.Split(key, "\\")[len(strings.Split(key, "\\"))-1]
func SafeCopy(origin string, destination string, printPreapend string) error {
	exists, err := FileExists(origin)
	if err != nil {
		fmt.Printf("%s [!] ER-FU001 Skyping file %s, %s \n", printPreapend, origin, err)
		return err
	}
	if !exists {
		fmt.Printf("%s ER-FU002 Skyping file %s, file not found \n", printPreapend, origin)
		return err
	}
	exists, err = FileExists(destination)
	if err != nil && !os.IsNotExist(err) {
		fmt.Printf("%s [!] ER-FU003 Skyping file %s, %s \n", printPreapend, destination, err)
		return err
	}
	if exists {
		fmt.Printf("%s ER-FU004 Skyping file %s, file already exists \n", printPreapend, destination)
		return nil
	}
	err = CopyFile(origin, destination)
	if err != nil {
		fmt.Printf("%s [!] ER-FU005 Error copying file %s, %s \n", printPreapend, destination, err)
		return err
	}
	return nil
}

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
