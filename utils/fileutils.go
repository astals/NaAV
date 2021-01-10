package utils

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

//driverVerboseName= strings.Split(key, "\\")[len(strings.Split(key, "\\"))-1]
func SafeCopy(origin string, destination string, printPrepend string) error {
	split := strings.Split(destination, "\\")
	path := strings.Join(split[:len(split)-1], "\\")
	err := CreateFoldersPath(path)
	if err != nil {
		fmt.Printf("%s [!] ER-FU000 %s ,%s \n", printPrepend, path, err)
	}
	exists, err := FileExists(origin)
	/*if err != nil {
		fmt.Printf("%s [!] ER-FU001 Skyping file %s, %s \n", printPrepend, origin, err)
		return err
	}
	if !exists {
		fmt.Printf("%s ER-FU002 Skyping file %s, file not found \n", printPrepend, origin)
		return err
	}*/
	exists, err = FileExists(destination)
	if err != nil && !os.IsNotExist(err) {
		fmt.Printf("%s [!] ER-FU003 Skyping file %s, %s \n", printPrepend, destination, err)
		return err
	}
	if exists {
		fmt.Printf("%s ER-FU004 Skyping file %s, file already exists \n", printPrepend, destination)
		return nil
	}
	err = CopyFile(origin, destination)
	if err != nil {
		fmt.Printf("%s [!] ER-FU005 Error copying file %s, %s \n", printPrepend, destination, err)
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

func ReadFile(file string) ([]byte, error) {
	bytesRead, err := ioutil.ReadFile(file)
	return bytesRead, err
}

func DeleteIfIsNaAVFile(file string, printPrepend string) (bool, error) {
	data, err := ReadFile(file)
	if err == nil && string(data) == "NaAV" {
		err := os.Remove(file)
		if err == nil {
			return true, nil
		} else {
			fmt.Printf("%s [!] ER-FU006 Error removing file %s, %s\n", printPrepend, file, err)
			return false, err
		}
	} else if err != nil {
		fmt.Printf("%s [!] ER-FU007 Error removing file %s, %s\n", printPrepend, file, err)
		return false, err
	} else {
		fmt.Printf("%s Skipping file %s, not NaAV file\n", printPrepend, file)
		return true, nil
	}
}

func CreateFoldersPath(path string) error {
	return os.MkdirAll(path, os.ModePerm)
}

func FindAbsolutePath(file string) (string, error) {
	res, _ := FileExists(file)
	if res == true {
		return file, nil
	} else {
		cwd, _ := os.Getwd()
		file = cwd + "\\" + file
		res, _ = FileExists(file)
		if res == true {
			return file, nil
		}
	}
	return "", errors.New("file not found")
}
