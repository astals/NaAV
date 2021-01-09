package VMware

import (
	"fmt"
	"os"

	"../../utils"
)

var files = map[string]string{
	"C:\\WINDOWS\\system32\\drivers\\vmhgfs.sys":  "resources\\dummy",
	"C:\\WINDOWS\\system32\\drivers\\vmmouse.sys": "resources\\dummy",
}

func InstallVMwareDrivers() {
	fmt.Printf("Copying VMware host drivers on system:\n")
	okOperations := 0
	for destination, origin := range files {
		err := utils.SafeCopy(origin, destination, "\t")
		if err == nil {
			okOperations++
		}
	}
	fmt.Printf("\t [i] Performed %d of %d operations\n", okOperations, len(files))
}

func UninstallVMwareDrivers() {
	fmt.Printf("Removing VMware host drivers on system:\n")
	okOperations := 0
	for destination, _ := range files {
		data, err := utils.ReadFile(destination)
		if err == nil && string(data) == "NaAV" {
			err := os.Remove(destination)
			if err == nil {
				okOperations++
			} else {
				fmt.Printf("\t [!] ER-VMW001 Error removing file %s, %s\n", destination, err)
			}
		} else if err != nil {
			fmt.Printf("\t [!] ER-VMW002 Error removing file %s, %s\n", destination, err)
		} else {
			fmt.Printf("\t Skipping file %s, not NaAV file\n", destination)
			okOperations++
		}
	}
	fmt.Printf("\t [i] Performed %d of %d operations\n", okOperations, len(files))
}

func CheckVMwareDrivers() {
	okOperations := 0
	for destination, _ := range files {
		res, _ := utils.FileExists(destination)
		if res {
			okOperations++
		}
	}
	fmt.Printf("\t [i] Found %d of %d files\n", okOperations, len(files))
}
