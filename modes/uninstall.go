package modes

import (
	"../modules/Others"
)

func Uninstall() {
	/*
		VMware.UninstallVMwareDrivers()
		VMware.UninstallVMwareGuestFiles()
		VirtualBox.UninstallVirtualBoxDrivers()
		VirtualBox.UninstallVirtualBoxGuestAdditionsFiles()
	*/
	Others.UninstallOtherRegistryKeys()
}
