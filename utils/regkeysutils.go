package utils

import (
	"errors"
	"fmt"
	"os/exec"
	"regexp"
	"strings"

	"golang.org/x/sys/windows/registry"
)

func InstallRegkeys(regkeys map[string][]string, verbosityLevel int, printPrepend string) {
	okOperations := 0
	NokOperations := 0
	for key, value := range regkeys {
		k, err := CreateRetrieveRegKey(key)
		if err != nil {
			printIfEnoughLevel(fmt.Sprintf("%s [!] ER-RU001 can't create/retrieve registry key %s ,%s \n", printPrepend, key, err), verbosityLevel, 1)
			NokOperations++
			continue
		}
		okOperations++
		printIfEnoughLevel(fmt.Sprintf("%s [i] Successfully created/retrieved registry key %s \n", printPrepend, key), verbosityLevel, 3)
		for _, v := range value {
			if !ExistsValuename(k, v) {
				err = WriteValue(k, v)
				if err != nil {
					printIfEnoughLevel(fmt.Sprintf("%s [!] ER-RU002 can't create registry namevalue %s  on registry key %s ,%s \n", printPrepend, v, key, err), verbosityLevel, 1)
					NokOperations++
					continue
				}
				printIfEnoughLevel(fmt.Sprintf("%s [i] Successfully created %s on registry key %s \n", printPrepend, v, key), verbosityLevel, 3)
			} else {
				printIfEnoughLevel(fmt.Sprintf("%s [i] Skipping namevalue %s on registry key %s, variable already exits \n", printPrepend, v, key), verbosityLevel, 2)
			}
			okOperations++
		}
	}
	printIfEnoughLevel(fmt.Sprintf("%s [i] Performed %d of %d operations\n", printPrepend, okOperations, okOperations+NokOperations), verbosityLevel, 0)
}

func UninstallRegkeys(regkeys map[string][]string, verbosityLevel int, printPrepend string) {
	okOperations := 0
	NokOperations := 0
	for key, value := range regkeys {
		exists := ExistsRegKey(key)
		if !exists {
			printIfEnoughLevel(fmt.Sprintf("%s Skipping registry key %s & it's namevalues, not found \n", printPrepend, key), verbosityLevel, 2)
			okOperations = okOperations + 1 + len(value)
			continue
		}
		k, err := CreateRetrieveRegKey(key)
		if err != nil {
			printIfEnoughLevel(fmt.Sprintf("%s [!] ER-RU002 can't retrieve registry key %s ,%s \n", printPrepend, key, err), verbosityLevel, 1)
			continue
		}
		valuenames, _ := GeyKeyValueNames(k)
		for _, v := range value {
			if ElementInStringArray(v, valuenames) {
				valuecontent := GetValue(k, v)
				if valuecontent != "NaAV" {
					printIfEnoughLevel(fmt.Sprintf("%s Skipping value %s on registry key %s, not NaAV value \n", printPrepend, v, key), verbosityLevel, 2)
					okOperations++
					continue
				}
				err = DeleteValue(k, v)
				if err != nil {
					printIfEnoughLevel(fmt.Sprintf("%s [!] ER-RU003 Unable to delete value %s on registry key %s \n", printPrepend, v, key), verbosityLevel, 1)
					NokOperations++
					continue
				} else {
					printIfEnoughLevel(fmt.Sprintf("%s [i] Successfully deleted valuename %s on key %s \n", printPrepend, v, key), verbosityLevel, 3)
					okOperations++
				}
			}
		}
		if len(valuenames) == 0 {
			err = DeleteKey(key)
		}
		if err != nil {
			printIfEnoughLevel(fmt.Sprintf("%s [!] ER-RU004 Unable to delete key %s \n", printPrepend, key), verbosityLevel, 1)
			NokOperations++
		} else {
			printIfEnoughLevel(fmt.Sprintf("%s [i] Successfully deleted key %s \n", printPrepend, key), verbosityLevel, 3)
			okOperations++
		}
	}
	printIfEnoughLevel(fmt.Sprintf("%s [i] Performed %d of %d operations\n", printPrepend, okOperations, okOperations+NokOperations), verbosityLevel, 0)
}

func CheckTrees(trees []string, verbosityLevel int, printPrepend string) {
	detected := 0
	for _, tree := range trees {
		exists := ExistsRegKey(tree)
		if exists {
			detected++
			printIfEnoughLevel(fmt.Sprintf("%s [i] detected %s \n", printPrepend, tree), verbosityLevel, 1)

		} else {
			printIfEnoughLevel(fmt.Sprintf("%s [i] Not detected %s \n", printPrepend, tree), verbosityLevel, 2)
		}
	}
	printIfEnoughLevel(fmt.Sprintf("%s [i] detected %d of %d trees\n", printPrepend, detected, len(trees)), verbosityLevel, 0)
}

func SafePurgeTrees(trees []string, verbosityLevel int, printPrepend string) {
	okOperations := 0
	for _, tree := range trees {
		exists := ExistsRegKey(tree)
		if exists {
			key, err := CreateRetrieveRegKey(tree)
			if err != nil {
				printIfEnoughLevel(fmt.Sprintf("%s [!] ER-RU005 can't retrieve registry key %s ,%s \n", printPrepend, tree, err), verbosityLevel, 1)
				continue
			}
			values, _ := GeyKeyValueNames(key)
			subkeys, _ := key.ReadSubKeyNames(-1)
			if len(subkeys) == 0 && len(values) == 0 {
				err := DeleteKey(tree)
				if err != nil {
					printIfEnoughLevel(fmt.Sprintf("%s [!] ER-RU006 Unable to delete key %s, %s \n", printPrepend, tree, err), verbosityLevel, 1)
				} else {
					printIfEnoughLevel(fmt.Sprintf("%s [i] Deleted key %s \n", printPrepend, tree), verbosityLevel, 3)
					okOperations++
				}
			}
		} else {
			printIfEnoughLevel(fmt.Sprintf("%s [i] Skipped key %s, not found \n", printPrepend, tree), verbosityLevel, 2)
			okOperations++
		}
	}
	printIfEnoughLevel(fmt.Sprintf("%s [i] Performed %d of %d operations\n", printPrepend, okOperations, len(trees)), verbosityLevel, 0)
}

func CheckRegkeys(regkeys map[string][]string, verbosityLevel int, printPrepend string) {
	detected := 0
	Ndetected := 0
	for key, value := range regkeys {
		exists := ExistsRegKey(key)
		if !exists {
			printIfEnoughLevel(fmt.Sprintf("%s [i] Not detected %s & it's namevalues\n", printPrepend, key), verbosityLevel, 2)
			Ndetected = Ndetected + 1 + len(value)
			continue
		}
		detected++
		printIfEnoughLevel(fmt.Sprintf("%s [i] Detected %s \n", printPrepend, key), verbosityLevel, 1)
		k, _ := CreateRetrieveRegKey(key)
		valuenames, _ := GeyKeyValueNames(k)
		for _, v := range value {
			if ElementInStringArray(v, valuenames) {
				detected++
				printIfEnoughLevel(fmt.Sprintf("%s [i] Detected %s value on key %s \n", printPrepend, v, key), verbosityLevel, 1)
			} else {
				Ndetected++
				printIfEnoughLevel(fmt.Sprintf("%s [i] Not detected %s value on key %s \n", printPrepend, v, key), verbosityLevel, 2)
			}
		}
	}
	printIfEnoughLevel(fmt.Sprintf("%s [i] Detected %d of %d Registry Keys & Valuenames \n", printPrepend, detected, detected+Ndetected), verbosityLevel, 0)
}

func ExistsRegKey(completepath string) bool {
	// TODO until i find a better way...
	var regex = regexp.MustCompile("^[a-zA-Z0-9_\\ ]*")
	if !regex.MatchString(completepath) {
		printIfEnoughLevel(fmt.Sprintf("------------> [!] who is my dirty boy? <------------\n"), 0, 0)
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

func printIfEnoughLevel(printstring string, verbosityLevel int, requiredMinLevel int) {
	if verbosityLevel >= requiredMinLevel {
		fmt.Print(printstring)
	}
}
