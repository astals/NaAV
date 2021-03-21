package modules

import (
	"fmt"

	"../utils"
)

func InstallFiles(Files []string, VerboseFileTypeName string, VerbosePlatformName string) {
	if len(Files) == 0 {
		utils.PrintIfEnoughLevel(fmt.Sprintf("Skipped install %s %s\n", VerbosePlatformName, VerboseFileTypeName), utils.OPERATION_SKIPPED_MESSAGE)
		return
	}
	utils.PrintIfEnoughLevel(fmt.Sprintf("Creating %s %s:\n", VerbosePlatformName, VerboseFileTypeName), utils.BASIC_INFORMATION_MESSAGE)
	okOperations := 0
	SkippedOperations :=0
	for _, file := range Files {
		exists, _ := utils.FileExists(file)
		if exists{
			utils.PrintIfEnoughLevel(fmt.Sprintf("%s Skipping file %s\n", "/t", file), utils.OPERATION_SKIPPED_MESSAGE)
			SkippedOperations++
		}
		err := utils.SafeCopy("resources\\dummyfile", file, "\t")
		if err == nil {
			okOperations++
		}
	}
	utils.PrintIfEnoughLevel(fmt.Sprintf("\t [i] Successfully performed %d of %d operations (%d skipped)\n", okOperations, len(Files),SkippedOperations), utils.SUMMARY_MESSAGE)
}

func UninstallFiles(Files []string, VerboseFileTypeName string, VerbosePlatformName string) {
	if len(Files) == 0 {
		utils.PrintIfEnoughLevel(fmt.Sprintf("Skipped uninstall %s %s\n", VerbosePlatformName, VerboseFileTypeName), utils.OPERATION_SKIPPED_MESSAGE)
		return
	}
	utils.PrintIfEnoughLevel(fmt.Sprintf("Removing %s %s:\n", VerbosePlatformName, VerboseFileTypeName), utils.BASIC_INFORMATION_MESSAGE)
	okOperations := 0
	SkippedOperations :=0
	for _, file := range Files {
		res, _ := utils.FileExists(file)
		if res {
			success, err := utils.DeleteIfIsNaAVFile(file, "\t")
			if success && err == nil {
				okOperations++
			}
			if !success && fmt.Sprint(err) == "not NaAV file"{
				utils.PrintIfEnoughLevel(fmt.Sprintf("%s Skipping file %s, not NaAV file\n", "/t", file), utils.OPERATION_SKIPPED_MESSAGE)
				SkippedOperations++
			}
		} else {
			SkippedOperations++
		}
	}
	utils.PrintIfEnoughLevel(fmt.Sprintf("\t [i] Successfully performed %d of %d operations (%d skipped)\n", okOperations, len(Files),SkippedOperations), utils.SUMMARY_MESSAGE)

}

func CheckFiles(Files []string, VerboseFileTypeName string, VerbosePlatformName string) {
	if len(Files) == 0 {
		utils.PrintIfEnoughLevel(fmt.Sprintf("Skipped check %s %s\n", VerbosePlatformName, VerboseFileTypeName), utils.OPERATION_SKIPPED_MESSAGE)
		return
	}
	utils.PrintIfEnoughLevel(fmt.Sprintf("Checking %s %s:\n", VerbosePlatformName, VerboseFileTypeName), utils.BASIC_INFORMATION_MESSAGE)
	okOperations := 0
	for _, file := range Files {
		res, _ := utils.FileExists(file)
		if res {
			okOperations++
		}
	}
	utils.PrintIfEnoughLevel(fmt.Sprintf("\t [i] Found %d of %d files\n", okOperations, len(Files)), utils.SUMMARY_MESSAGE)
}
