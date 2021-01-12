package Others

import (
	"../../utils"
)

var regKeysAndValueNames = map[string][]string{
	"HKEY_CURRENT_USER\\SOFTWARE\\Wine":       []string{"a"},
	"HKEY_CURRENT_USER\\SOFTWARE\\test\\test": []string{"a"},
}

func InstallOtherRegistryKeys() {
	utils.InstallRegkeys(regKeysAndValueNames, "\t")
}
