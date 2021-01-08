package VMware

import (
	"fmt"

	"../../utils"
)

func InstallVMwareDrivers() {
	files := map[string]string{
		"C:\\WINDOWS\\system32\\drivers\\vmhgfs.sys":  "resources\\dummy",
		"C:\\WINDOWS\\system32\\drivers\\vmmouse.sys": "resources\\dummy",
	}
	fmt.Printf("Copying VMware host drivers on system:\n")
	okOperations := 0
	for key, value := range files {
		err := utils.SafeCopy(value, key, "\t")
		if err == nil {
			okOperations++
		}
	}
	fmt.Printf("\t Successfuly performed %d of %d operations\n", okOperations, len(files))
}

func UninstallVMwareDrivers() {
	//TODO
}

func CheckVMwareDrivers() {
	//TODO
}
