package modules

import (
	"fmt"

	"../utils"
)

func InstallNetworkInterfaces(adapters map[string]string, VerbosePlatformName string) {
	if len(adapters) == 0 {
		utils.PrintIfEnoughLevel(fmt.Sprintf("Skipped install %s network interfaces\n", VerbosePlatformName), utils.OPERATION_SKIPPED_MESSAGE)
		return
	}
	utils.PrintIfEnoughLevel(fmt.Sprintf("Creating %s network interfaces:\n", VerbosePlatformName), utils.BASIC_INFORMATION_MESSAGE)
	okOperations := 0
	for name, mac := range adapters {
		err := utils.CreateNetworkAdapter(name, mac, "\t")
		if err == nil {
			okOperations++
		}
	}
	utils.PrintIfEnoughLevel(fmt.Sprintf("%s [i] Successfully performed %d of %d operations\n", "\t", okOperations, len(adapters)), utils.SUMMARY_MESSAGE)

}
func UninstallNetworkInterfaces(adapters map[string]string, VerbosePlatformName string) {
	if len(adapters) == 0 {
		utils.PrintIfEnoughLevel(fmt.Sprintf("Skipped uninstall %s network interfaces\n", VerbosePlatformName), utils.OPERATION_SKIPPED_MESSAGE)
		return
	}
	utils.PrintIfEnoughLevel(fmt.Sprintf("Removing %s network interfaces:\n", VerbosePlatformName), utils.BASIC_INFORMATION_MESSAGE)
	okOperations := 0
	SkippedOperations := 0
	for name, mac := range adapters {
		if !utils.ExistsNetworkAdapterStatingWith(mac, "\t") {
			SkippedOperations++
			continue
		}
		err := utils.DeleteNetworkAdapter(name, "\t")
		if err == nil {
			okOperations++
		}
	}
	utils.PrintIfEnoughLevel(fmt.Sprintf("%s [i] Successfully performed %d of %d operations (%d skipped)\n", "\t", okOperations, len(adapters),SkippedOperations), utils.SUMMARY_MESSAGE)
}

func CheckNetworkInterfaces(adapters map[string]string, VerbosePlatformName string) {
	if len(adapters) == 0 {
		utils.PrintIfEnoughLevel(fmt.Sprintf("Skipped check %s network interfaces\n", VerbosePlatformName), utils.OPERATION_SKIPPED_MESSAGE)
		return
	}
	utils.PrintIfEnoughLevel(fmt.Sprintf("Checking %s network interfaces:\n", VerbosePlatformName), utils.BASIC_INFORMATION_MESSAGE)
	detected := 0
	for _, mac := range adapters {
		if utils.ExistsNetworkAdapterStatingWith(mac[0:8], "\t") {
			detected++
		}
	}
	utils.PrintIfEnoughLevel(fmt.Sprintf("%s [i] Detected %d of %d mac addresses \n", "\t", detected, len(adapters)), 0)
}
