package modes

import (
	"../modules/VMware"
)

func Uninstall() {
	VMware.UninstallVMwareDrivers()
}
