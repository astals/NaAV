package VirtualBox

import (
	"../../utils"
)

var regKeysAndValueNames = map[string][]string{
	"\\HKEY_LOCAL_MACHINE\\HARDWARE\\ACPI\\DSDT\\VBOX__\\VBOXBIOS\\00000002": []string{"00000000"},
}

func InstallVirtualBoxRegistry() {
	utils.InstallRegkeys(regKeysAndValueNames, "\t")
}
