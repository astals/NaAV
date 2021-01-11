package modes

import (
	"../modules/VMware"
	"../modules/VirtualBox"
)

func Check() {
	VMware.CheckVMwareDrivers()
	VMware.CheckVMwareGuestFiles()
	VirtualBox.CheckVirtualBoxGuestAdditionsFiles()
}
