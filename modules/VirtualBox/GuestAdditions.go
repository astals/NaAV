package VirtualBox

import (
	"fmt"

	"../../utils"
)

var files = map[string]string{
	"C:\\Program Files\\Oracle\\VirtualBox Guest Additions\\DIFxAPI.dll":                              "resources\\dummy",
	"C:\\Program Files\\Oracle\\VirtualBox Guest Additions\\iexplore.ico":                             "resources\\dummy",
	"C:\\Program Files\\Oracle\\VirtualBox Guest Additions\\install_drivers.log":                      "resources\\dummy",
	"C:\\Program Files\\Oracle\\VirtualBox Guest Additions\\Oracle VM VirtualBox Guest Additions.url": "resources\\dummy",
	"C:\\Program Files\\Oracle\\VirtualBox Guest Additions\\uninst.exe":                               "resources\\dummy",
	"C:\\Program Files\\Oracle\\VirtualBox Guest Additions\\VBoxControl.exe":                          "resources\\dummy",
	"C:\\Program Files\\Oracle\\VirtualBox Guest Additions\\VBoxDisp.dll":                             "resources\\dummy",
	"C:\\Program Files\\Oracle\\VirtualBox Guest Additions\\VBoxDispD3D-x86.dll":                      "resources\\dummy",
	"C:\\Program Files\\Oracle\\VirtualBox Guest Additions\\VBoxDispD3D.dll":                          "resources\\dummy",
	"C:\\Program Files\\Oracle\\VirtualBox Guest Additions\\VBoxDrvInst.exe":                          "resources\\dummy",
	"C:\\Program Files\\Oracle\\VirtualBox Guest Additions\\VBoxGL-x86.dll":                           "resources\\dummy",
	"C:\\Program Files\\Oracle\\VirtualBox Guest Additions\\VBoxGL.dll":                               "resources\\dummy",
	"C:\\Program Files\\Oracle\\VirtualBox Guest Additions\\VBoxGuest.cat":                            "resources\\dummy",
	"C:\\Program Files\\Oracle\\VirtualBox Guest Additions\\VBoxGuest.inf":                            "resources\\dummy",
	"C:\\Program Files\\Oracle\\VirtualBox Guest Additions\\VBoxGuest.sys":                            "resources\\dummy",
	"C:\\Program Files\\Oracle\\VirtualBox Guest Additions\\VBoxICD-x86.dll":                          "resources\\dummy",
	"C:\\Program Files\\Oracle\\VirtualBox Guest Additions\\VBoxICD.dll":                              "resources\\dummy",
	"C:\\Program Files\\Oracle\\VirtualBox Guest Additions\\VBoxMouse.cat":                            "resources\\dummy",
	"C:\\Program Files\\Oracle\\VirtualBox Guest Additions\\VBoxMouse.inf":                            "resources\\dummy",
	"C:\\Program Files\\Oracle\\VirtualBox Guest Additions\\VBoxMouse.sys":                            "resources\\dummy",
	"C:\\Program Files\\Oracle\\VirtualBox Guest Additions\\VBoxNine-x86.dll":                         "resources\\dummy",
	"C:\\Program Files\\Oracle\\VirtualBox Guest Additions\\VBoxNine.dll":                             "resources\\dummy",
	"C:\\Program Files\\Oracle\\VirtualBox Guest Additions\\VBoxSVGA-x86.dll":                         "resources\\dummy",
	"C:\\Program Files\\Oracle\\VirtualBox Guest Additions\\VBoxSVGA.dll":                             "resources\\dummy",
	"C:\\Program Files\\Oracle\\VirtualBox Guest Additions\\VBoxTray.exe":                             "resources\\dummy",
	"C:\\Program Files\\Oracle\\VirtualBox Guest Additions\\VBoxVideo.cat":                            "resources\\dummy",
	"C:\\Program Files\\Oracle\\VirtualBox Guest Additions\\VBoxVideo.inf":                            "resources\\dummy",
	"C:\\Program Files\\Oracle\\VirtualBox Guest Additions\\VBoxVideo.sys":                            "resources\\dummy",
	"C:\\Program Files\\Oracle\\VirtualBox Guest Additions\\VBoxWddm.cat":                             "resources\\dummy",
	"C:\\Program Files\\Oracle\\VirtualBox Guest Additions\\VBoxWddm.inf":                             "resources\\dummy",
	"C:\\Program Files\\Oracle\\VirtualBox Guest Additions\\VBoxWddm.sys":                             "resources\\dummy",
	"C:\\Program Files\\Oracle\\VirtualBox Guest Additions\\VBoxWHQLFake.exe":                         "resources\\dummy",
}

func InstallVirtualBoxGuestAdditionsFiles() {
	fmt.Printf("Copying VirtualBox Guest Additions files on system:\n")
	okOperations := 0
	for destination, origin := range files {
		err := utils.SafeCopy(origin, destination, "\t")
		if err == nil {
			okOperations++
		}
	}
	fmt.Printf("\t [i] Performed %d of %d operations\n", okOperations, len(files))
}

func UninstallVirtualBoxGuestAdditionsFiles() {
	fmt.Printf("Removing VirtualBox Guest Additions files on system:\n")
	okOperations := 0
	for destination, _ := range files {
		success, err := utils.DeleteIfIsNaAVFile(destination, "\t")
		if success && err == nil {
			okOperations++
		}
	}
	fmt.Printf("\t [i] Performed %d of %d operations\n", okOperations, len(files))
}

func CheckVirtualBoxGuestAdditionsFiles() {
	fmt.Printf("Checking VirtualBox Guest Additions files on system:\n")
	okOperations := 0
	for destination, _ := range files {
		res, _ := utils.FileExists(destination)
		if res {
			okOperations++
		}
	}
	fmt.Printf("\t [i] Found %d of %d files\n", okOperations, len(files))
}
