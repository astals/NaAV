package Others

import (
	"fmt"

	"../../utils"
)

var regKeysAndValueNames = map[string][]string{
	//"HKEY_CURRENT_USER\\SOFTWARE\\Wine":       []string{"a"},
	//"HKEY_CURRENT_USER\\SOFTWARE\\test\\test": []string{"a"},
}

// order is relevant
var significantParentTrees = []string{"HKEY_CURRENT_USER\\SOFTWARE\\test"}

func InstallOtherRegistryKeys() {
	fmt.Printf("Creating Other Registry keys on system:\n")
	utils.InstallRegkeys(regKeysAndValueNames, 3, "\t")
}

func UninstallOtherRegistryKeys() {
	fmt.Printf("Removing Other registry keys on system\n")
	utils.UninstallRegkeys(regKeysAndValueNames, 3, "\t")
	fmt.Printf("\t - Purging trees\n")
	utils.SafePurgeTrees(significantParentTrees, 3, "\t\t")
}

func CheckOtherRegistryKeys() {
	fmt.Printf("Checking Other registry keys on system\n")
	utils.CheckRegkeys(regKeysAndValueNames, 3, "\t")
	fmt.Printf("\t - Detecting parent trees\n")
	utils.CheckTrees(significantParentTrees, 3, "\t\t")

}
