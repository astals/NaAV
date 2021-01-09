package modes

import (
	"../modules/VMware"
	"../modules/VirtualBox"
)

func Install() {
	VMware.InstallVMwareDrivers()
	VirtualBox.InstallVirtualBoxGuestAdditionsFiles()
}
