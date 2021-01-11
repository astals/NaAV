package VMware

import (
	"fmt"

	"../../utils"
)

var driverfiles = map[string]string{
	"C:\\WINDOWS\\system32\\drivers\\vmhgfs.sys":  "resources\\dummy",
	"C:\\WINDOWS\\system32\\drivers\\vmmouse.sys": "resources\\dummy",
}

func InstallVMwareDrivers() {
	fmt.Printf("Copying VMware host drivers on system:\n")
	okOperations := 0
	for destination, origin := range driverfiles {
		err := utils.SafeCopy(origin, destination, "\t")
		if err == nil {
			okOperations++
		}
	}
	fmt.Printf("\t [i] Performed %d of %d operations\n", okOperations, len(driverfiles))
}

func UninstallVMwareDrivers() {
	fmt.Printf("Removing VMware host drivers on system:\n")
	okOperations := 0
	for destination, _ := range driverfiles {
		res, _ := utils.FileExists(destination)
		if res {
			success, err := utils.DeleteIfIsNaAVFile(destination, "\t")
			if success && err == nil {
				okOperations++
			}
		} else {
			okOperations++
		}
	}
	fmt.Printf("\t [i] Performed %d of %d operations\n", okOperations, len(driverfiles))
}

func CheckVMwareDrivers() {
	fmt.Printf("Checking  VMware host drivers on system:\n")
	okOperations := 0
	for destination, _ := range driverfiles {
		res, _ := utils.FileExists(destination)
		if res {
			okOperations++
		}
	}
	fmt.Printf("\t [i] Found %d of %d files\n", okOperations, len(driverfiles))
}
