package utils

import (
	"errors"
	"fmt"
	"strings"

	"golang.org/x/sys/windows/registry"
)

func InstallRegkeys(regkeys map[string][]string, printPrepend string) {
	okOperations := 0
	for key, value := range regkeys {
		k, err := CreateRetrieveRegKey(key)
		if err != nil {
			fmt.Printf("%s [!] ER-FU001 can't create registry key %s ,%s \n", printPrepend, key, err)
			continue
		}
		for _, v := range value {
			err = WriteValue(k, v)
			if err != nil {
				fmt.Printf("%s [!] ER-FU002 can't create registry namevalue %s  on registry key %s ,%s \n", printPrepend, v, key, err)
				continue
			}
			okOperations++
		}
	}
	fmt.Printf("%s [i] Performed %d of %d operations\n", printPrepend, okOperations, len(regkeys))
}

/*
func UninstallRegkeys(regkeys map[string][]string, printPrepend string) error {
	okOperations := 0
	for key, value  := range regkeys {
		k, err := CreateRetrieveRegKey(key)
		if err != nil {
			continue
		}
		value, err = GetValue(k, key)
		if err != nil {
			continue
		}
		if value == "NaAV" {
			DeleteValueName(k)
		}
		if err != nil {
			continue
		}
		valuenames, err := GeyKeyValueNames(k)
		if err != nil {

		} else if len(valuenames) == 0 {
			err = DeleteKey(key.KeyComplatePath)
		}
		if err != nil {

		}
		okOperations++
	}
	fmt.Printf("%s [i] Performed %d of %d operations\n", printPrepend, okOperations, len(regkeys))
}

func CheckRegkeys(regkeys map[string][]string, printPrepend string) error {
	okOperations := 0
	paths := []string
	// checking keys
	for index, key := range regkeys {
		paths = append(paths, key.KeyComplatePath)
	}
	uniquepaths := utils.uniquestrings(paths)
	for index, path := range uniquepaths {
		if ExistsRegKey(path) {
		}
	}
	// checking valuenames
	for index, path := range regkeys {
		if ExistsRegKey(path) {
			CreateRetrieveRegKey(path)
			if ExistsValuename(key,valuename) {
			}

		}
	}
	fmt.Printf("%s [i] Performed %d of %d operations\n", printPrepend, okOperations, len(regkeys))
}*/

func ExistsRegKey(completepath string) (bool, error) {
	key, path := splitKeyString(completepath)
	key, exists, err := registry.CreateKey(key, path, registry.QUERY_VALUE|registry.SET_VALUE)
	if err != nil {
		return false, err
	}
	if exists {
		return true, nil
	} else {
		DeleteKey(completepath)
		return false, nil
	}
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
	return key.SetStringValue(valuename, "NaAV")
}

func GetValue(key registry.Key, valuename string) string {
	val, _, _ := key.GetStringValue(valuename)
	return val
}

func DeleteValueName(key registry.Key, valuename string) {
	key.DeleteValue(valuename)
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
