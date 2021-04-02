package modules

import (
	"fmt"

	"../utils"
)

func InstallRegkeys(regkeys map[string][]string, VerbosePlatformName string) {
	if len(regkeys) == 0 {
		utils.PrintIfEnoughLevel(fmt.Sprintf("Skipped install %s Registry Keys & ValueNames\n", VerbosePlatformName), utils.OPERATION_SKIPPED_MESSAGE)
		return
	}
	utils.PrintIfEnoughLevel(fmt.Sprintf("Creating %s Registry Keys & ValueNames \n", VerbosePlatformName), utils.SUMMARY_MESSAGE)
	okOperations := 0
	NoOKperations := 0
	SkippedOperations :=0
	for key, value := range regkeys {
		k, err := utils.CreateRetrieveRegKey(key)
		if err != nil {
			utils.PrintIfEnoughLevel(fmt.Sprintf("%s [!] FR-RU001 can't create/retrieve registry key %s ,%s \n", "\t", key, err), utils.OPERATION_ERROR_MESSAGE)
			NoOKperations++
			continue
		}
		okOperations++
		utils.PrintIfEnoughLevel(fmt.Sprintf("%s [i] Successfully created/retrieved registry key %s \n", "\t", key), utils.OPERATION_SUCCESS_MESSAGE)
		for _, v := range value {
			if !utils.ExistsValuename(k, v) {
				err = utils.WriteValue(k, v)
				if err != nil {
					utils.PrintIfEnoughLevel(fmt.Sprintf("%s [!] ER-FR002 can't create registry namevalue %s  on registry key %s ,%s \n", "\t", v, key, err), utils.OPERATION_ERROR_MESSAGE)
					NoOKperations++
					continue
				}
				utils.PrintIfEnoughLevel(fmt.Sprintf("%s [i] Successfully created %s on registry key %s \n", "\t", v, key), utils.OPERATION_SUCCESS_MESSAGE)
			} else {
				utils.PrintIfEnoughLevel(fmt.Sprintf("%s [i] Skipping namevalue %s on registry key %s, variable already exits \n", "\t", v, key), utils.OPERATION_SKIPPED_MESSAGE)
				SkippedOperations++
			}
			okOperations++
		}
	}
	utils.PrintIfEnoughLevel(fmt.Sprintf("%s [i] Successfully performed %d of %d operations (%d skipped)\n", "\t", okOperations, okOperations+NoOKperations+SkippedOperations,SkippedOperations), utils.SUMMARY_MESSAGE)
}

func UninstallRegkeys(regkeys map[string][]string, VerbosePlatformName string) {
	if len(regkeys) == 0 {
		utils.PrintIfEnoughLevel(fmt.Sprintf("Skipped uninstall %s Registry Keys & ValueNames\n", VerbosePlatformName), utils.OPERATION_SKIPPED_MESSAGE)
		return
	}
	utils.PrintIfEnoughLevel(fmt.Sprintf("Removing %s Registry Keys & ValueNames \n", VerbosePlatformName), utils.BASIC_INFORMATION_MESSAGE)
	okOperations := 0
	NoOKperations := 0
	SkippedOperations := 0
	for key, value := range regkeys {
		exists := utils.ExistsRegKey(key)
		if !exists {
			utils.PrintIfEnoughLevel(fmt.Sprintf("%s Skipping registry key %s & it's namevalues, not found \n", "\t", key), utils.OPERATION_SKIPPED_MESSAGE)
			SkippedOperations = SkippedOperations + 1 + len(value)
			continue
		}
		k, err := utils.CreateRetrieveRegKey(key)
		if err != nil {
			utils.PrintIfEnoughLevel(fmt.Sprintf("%s [!] ER-FR003 can't retrieve registry key %s ,%s \n", "\t", key, err), utils.OPERATION_ERROR_MESSAGE)
			continue
		}
		valuenames, _ := utils.GeyKeyValueNames(k)
		for _, v := range value {
			if utils.ElementInStringArray(v, valuenames) {
				valuecontent := utils.GetValue(k, v)
				if valuecontent != "NaAV" {
					utils.PrintIfEnoughLevel(fmt.Sprintf("%s Skipping value %s on registry key %s, not NaAV value \n", "\t", v, key), utils.OPERATION_SKIPPED_MESSAGE)
					SkippedOperations++
					continue
				}
				err = utils.DeleteValue(k, v)
				if err != nil {
					utils.PrintIfEnoughLevel(fmt.Sprintf("%s [!] ER-FR004 Unable to delete value %s on registry key %s \n", "\t", v, key), utils.OPERATION_ERROR_MESSAGE)
					NoOKperations++
					continue
				} else {
					utils.PrintIfEnoughLevel(fmt.Sprintf("%s [i] Successfully deleted valuename %s on key %s \n", "\t", v, key), utils.OPERATION_SUCCESS_MESSAGE)
					okOperations++
				}
			}
		}
		valuenames, _ = utils.GeyKeyValueNames(k)
		if len(valuenames) == 0 {
			err = utils.DeleteKey(key)
		}
		if err != nil {
			utils.PrintIfEnoughLevel(fmt.Sprintf("%s [!] ER-FR005 Unable to delete key %s: %s  \n", "\t", key, err), utils.OPERATION_ERROR_MESSAGE)
			NoOKperations++
		} else {
			utils.PrintIfEnoughLevel(fmt.Sprintf("%s [i] Successfully deleted key %s \n", "\t", key), utils.OPERATION_SUCCESS_MESSAGE)
			okOperations++
		}
	}
	utils.PrintIfEnoughLevel(fmt.Sprintf("%s [i] Successfully performed %d of %d operations (%d skipped)\n", "\t", okOperations, okOperations+NoOKperations+SkippedOperations,SkippedOperations), utils.SUMMARY_MESSAGE)
}

func CheckRegkeys(regkeys map[string][]string, VerbosePlatformName string) {
	if len(regkeys) == 0 {
		utils.PrintIfEnoughLevel(fmt.Sprintf("Skipped check %s Registry Keys & ValueNames\n", VerbosePlatformName), utils.OPERATION_SKIPPED_MESSAGE)
		return
	}
	utils.PrintIfEnoughLevel(fmt.Sprintf("Checking %s Registry Keys & ValueNames \n", VerbosePlatformName), utils.SUMMARY_MESSAGE)
	detected := 0
	Ndetected := 0
	for key, value := range regkeys {
		exists := utils.ExistsRegKey(key)
		if !exists {
			utils.PrintIfEnoughLevel(fmt.Sprintf("%s [i] Not detected %s & it's namevalues\n", "\t", key), utils.ITEM_NOT_FOUND_MESSAGE)
			Ndetected = Ndetected + 1 + len(value)
			continue
		}
		detected++
		utils.PrintIfEnoughLevel(fmt.Sprintf("%s [i] Detected %s \n", "\t", key), utils.FOUND_ITEM_MESSAGE)
		k, _ := utils.CreateRetrieveRegKey(key)
		valuenames, _ := utils.GeyKeyValueNames(k)
		for _, v := range value {
			if utils.ElementInStringArray(v, valuenames) {
				detected++
				utils.PrintIfEnoughLevel(fmt.Sprintf("%s [i] Detected %s value on key %s \n", "\t", v, key), utils.FOUND_ITEM_MESSAGE)
			} else {
				Ndetected++
				utils.PrintIfEnoughLevel(fmt.Sprintf("%s [i] Not detected %s value on key %s \n", "\t", v, key), utils.ITEM_NOT_FOUND_MESSAGE)
			}
		}
	}
	utils.PrintIfEnoughLevel(fmt.Sprintf("%s [i] Detected %d of %d Registry Keys & Valuenames \n", "\t", detected, detected+Ndetected), utils.SUMMARY_MESSAGE)
}

func CheckTrees(trees []string) {
	if len(trees) == 0 {
		utils.PrintIfEnoughLevel(fmt.Sprintf("%s [i] Skipped check parent trees \n", "\t"), utils.OPERATION_SKIPPED_MESSAGE)
		return
	}
	utils.PrintIfEnoughLevel(fmt.Sprintf("%s --- Checking parent trees\n", "\t"), utils.BASIC_INFORMATION_MESSAGE)
	detected := 0
	for _, tree := range trees {
		exists := utils.ExistsRegKey(tree)
		if exists {
			detected++
			utils.PrintIfEnoughLevel(fmt.Sprintf("%s [i] detected %s \n", "\t\t", tree), utils.FOUND_ITEM_MESSAGE)

		} else {
			utils.PrintIfEnoughLevel(fmt.Sprintf("%s [i] Not detected %s \n", "\t\t", tree), utils.ITEM_NOT_FOUND_MESSAGE)
		}
	}
	utils.PrintIfEnoughLevel(fmt.Sprintf("%s [i] detected %d of %d trees\n", "\t\t", detected, len(trees)), utils.SUMMARY_MESSAGE)
}

func UninstallParentTrees(trees []string) {
	if len(trees) == 0 {
		utils.PrintIfEnoughLevel(fmt.Sprintf("%s [i] Skipped uninstall parent trees\n", "\t"), utils.OPERATION_SKIPPED_MESSAGE)
		return
	}
	utils.PrintIfEnoughLevel(fmt.Sprintf("%s --- Removing parent trees\n", "\t"), utils.BASIC_INFORMATION_MESSAGE)
	okOperations := 0
	SkippedOperations :=0
	for _, tree := range trees {
		exists := utils.ExistsRegKey(tree)
		if exists {
			key, err := utils.CreateRetrieveRegKey(tree)
			if err != nil {
				utils.PrintIfEnoughLevel(fmt.Sprintf("%s [!] ER-RU005 can't retrieve registry key %s ,%s \n", "\t\t", tree, err), utils.OPERATION_ERROR_MESSAGE)
				continue
			}
			values, _ := utils.GeyKeyValueNames(key)
			subkeys, _ := key.ReadSubKeyNames(-1)
			if len(subkeys) == 0 && len(values) == 0 {
				err := utils.DeleteKey(tree)
				if err != nil {
					utils.PrintIfEnoughLevel(fmt.Sprintf("%s [!] ER-RU006 Unable to delete key %s, %s \n", "\t\t", tree, err), utils.OPERATION_ERROR_MESSAGE)
				} else {
					utils.PrintIfEnoughLevel(fmt.Sprintf("%s [i] Deleted key %s \n", "\t\t", tree), utils.OPERATION_SUCCESS_MESSAGE)
					okOperations++
				}
			}
		} else {
			utils.PrintIfEnoughLevel(fmt.Sprintf("%s [i] Skipped key %s, not found \n", "\t\t", tree), utils.OPERATION_SKIPPED_MESSAGE)
			SkippedOperations++
		}
	}
	utils.PrintIfEnoughLevel(fmt.Sprintf("%s [i] Successfully performed %d of %d operations (%d skipped)\n", "\t\t", okOperations, len(trees),SkippedOperations), utils.SUMMARY_MESSAGE)
}
