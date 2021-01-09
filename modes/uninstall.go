package modes

import (
	"../modules/VMware"
	"../modules/VirtualBox"
)

func Uninstall() {
	VMware.UninstallVMwareDrivers()
	VirtualBox.UninstallVirtualBoxGuestAdditionsFiles()
}
