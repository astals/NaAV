package modes

import (
	"../modules/VMware"
)

func Install() {
	VMware.InstallVMwareDrivers()
}
