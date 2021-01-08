package VMware

import (
	"fmt"

	"../../utils"
)

func InstallVMwareDrivers() {
	files := map[string]string{
		"C:\\WINDOWS\\system32\\drivers\\vmhgfs.sys":  "resources\\VMware\\vmhgfs.sys",
		"C:\\WINDOWS\\system32\\drivers\\vmmouse.sys": "resources\\VMware\\vmmouse.sys",
	}
	fmt.Print("Copying VMware host drivers on system: ")
	return
	okOperations := 0
	for key, value := range files {
		//driverVerboseName= strings.Split(key, "\\")[len(strings.Split(key, "\\"))-1]
		exists, err := utils.FileExists(key)
		if err != nil {
			fmt.Print("\t [-] Skyping driver %s, %s", key, err)
		}
		if exists {
			fmt.Print("\t Skyping driver %s, file already exists", key)
		} else {
			err = utils.CopyFile(value, key)
			if err == nil {
				okOperations++
			} else {
				fmt.Print("\t [-] Error copying driver %s, %s", key, err)
			}
		}
	}
}

func UninstallVMwareDrivers() {
	//TODO
}

func CheckVMwareDrivers() {
	//TODO
}
