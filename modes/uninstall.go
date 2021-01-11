package modes

import (
	"../modules/VMware"
	"../modules/VirtualBox"
)

func Uninstall() {
	VMware.UninstallVMwareDrivers()
	VMware.UninstallVMwareGuestFiles()
	VirtualBox.UninstallVirtualBoxDrivers()
	VirtualBox.UninstallVirtualBoxGuestAdditionsFiles()
}
