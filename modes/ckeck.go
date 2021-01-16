package modes

import (
	"../modules/Others"
	"../modules/VMware"
	"../modules/VirtualBox"
)

func Check() {
	VMware.CheckVMwareDrivers()
	VMware.CheckVMwareGuestFiles()
	VirtualBox.CheckVirtualBoxDrivers()
	VirtualBox.CheckVirtualBoxGuestAdditionsFiles()
	Others.CheckOtherRegistryKeys()
}
