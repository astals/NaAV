package modes

import (
	"../modules/VMware"
)

func Check() {
	VMware.CheckVMwareDrivers()
}
