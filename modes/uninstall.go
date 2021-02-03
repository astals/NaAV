package modes

import (
	"../modules/VirtualBox"
)

func Uninstall() {
	/*
		VMware.UninstallVMwareDrivers()
		VMware.UninstallVMwareGuestFiles()
		VirtualBox.UninstallVirtualBoxDrivers()
		VirtualBox.UninstallVirtualBoxGuestAdditionsFiles()
		Others.UninstallOtherRegistryKeys()
	*/
	VirtualBox.UninstallVirtualBoxFakeGuestNetworkInterface()
}
