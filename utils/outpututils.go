package utils

import "fmt"

var VerbosityLevel int //=1

// required vervosity levels
const (
	SUMMARY_MESSAGE           = 0
	BASIC_INFORMATION_MESSAGE = 0
	OPERATION_ERROR_MESSAGE   = 1
	FOUND_ITEM_MESSAGE        = 2
	OPERATION_SKIPPED_MESSAGE = 2
	ITEM_NOT_FOUND_MESSAGE    = 3
	OPERATION_SUCCESS_MESSAGE = 3
)

func PrintIfEnoughLevel(printstring string, requiredMinLevel int) {
	if VerbosityLevel >= requiredMinLevel {
		fmt.Print(printstring)
	}
}

func SetVerbosityLevel() {
	VerbosityLevel = 1
}
