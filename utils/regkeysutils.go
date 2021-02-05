package utils

import (
	"errors"
	"fmt"
	"os/exec"
	"regexp"
	"strings"

	"golang.org/x/sys/windows/registry"
)

func ExistsRegKey(completepath string) bool {
	// TODO until i find a better way...
	var regex = regexp.MustCompile("^[a-zA-Z0-9_\\ ]*")
	if !regex.MatchString(completepath) {
		PrintIfEnoughLevel(fmt.Sprintf("------------> [!] who is my dirty boy? <------------\n"), -99999)
		return false
	}
	_, err := exec.Command("reg", "query", completepath).Output()
	if err != nil {
		return false
	}
	return true
}

func ExistsValuename(key registry.Key, valuename string) bool {
	valuenames, _ := GeyKeyValueNames(key)
	return ElementInStringArray(valuename, valuenames)
}

func CreateRetrieveRegKey(completepath string) (registry.Key, error) {
	var key registry.Key
	if ValidCompletePath(completepath) {
		key, path := splitKeyString(completepath)
		key, _, err := registry.CreateKey(key, path, registry.QUERY_VALUE|registry.SET_VALUE)
		if err != nil {
			return key, err
		}
		return key, nil
	}
	return key, errors.New("Invalid key path")
}

func WriteValue(key registry.Key, valuename string) error {
	//check if value exists
	return key.SetStringValue(valuename, "NaAV")
}

func GetValue(key registry.Key, valuename string) string {
	val, _, _ := key.GetStringValue(valuename)
	return val
}

func DeleteValue(key registry.Key, valuename string) error {
	return key.DeleteValue(valuename)
}

func GeyKeyValueNames(k registry.Key) ([]string, error) {
	return k.ReadValueNames(-1)
}

func DeleteKey(completepath string) error {
	key, path := splitKeyString(completepath)
	return registry.DeleteKey(key, path)
}

func splitKeyString(completepath string) (registry.Key, string) {
	if strings.HasPrefix(completepath, "HKEY_CLASSES_ROOT\\") {
		return registry.CLASSES_ROOT, strings.Replace(completepath, "HKEY_CLASSES_ROOT\\", "", 1)
	}
	if strings.HasPrefix(completepath, "HKEY_CURRENT_USER\\") {
		return registry.CURRENT_USER, strings.Replace(completepath, "HKEY_CURRENT_USER\\", "", 1)
	}
	if strings.HasPrefix(completepath, "HKEY_LOCAL_MACHINE\\") {
		return registry.LOCAL_MACHINE, strings.Replace(completepath, "HKEY_LOCAL_MACHINE\\", "", 1)
	}
	if strings.HasPrefix(completepath, "HKEY_USERS\\") {
		return registry.USERS, strings.Replace(completepath, "HKEY_USERS\\", "", 1)
	}
	if strings.HasPrefix(completepath, "HKEY_CURRENT_CONFIG\\") {
		return registry.CURRENT_CONFIG, strings.Replace(completepath, "HKEY_CURRENT_CONFIG\\", "", 1)
	}
	return registry.NONE, ""
}

func ValidCompletePath(completepath string) bool {
	if strings.HasPrefix(completepath, "HKEY_CLASSES_ROOT\\") || strings.HasPrefix(completepath, "HKEY_CURRENT_USER\\") || strings.HasPrefix(completepath, "HKEY_LOCAL_MACHINE\\") || strings.HasPrefix(completepath, "HKEY_USERS\\") || strings.HasPrefix(completepath, "HKEY_CURRENT_CONFIG\\") {
		return true
	}
	return false
}
