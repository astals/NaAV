package utils

import "fmt"

var VerbosityLevel int = 1

// required vervosity levels
const (
	SUMMARY_MESSAGE           = 0
	BASIC_INFORMATION_MESSAGE = 0
	HARDWARE_RECOGNITION_MESSAGE = 0
	OPERATION_ERROR_MESSAGE   = 1
	OPERATION_SKIPPED_MESSAGE = 2
	OPERATION_SUCCESS_MESSAGE = 3
	FOUND_ITEM_MESSAGE        = 3
	ITEM_NOT_FOUND_MESSAGE    = 4	
)

func PrintIfEnoughLevel(printstring string, requiredMinLevel int) {
	// TODO append \n at here
	if VerbosityLevel >= requiredMinLevel {
		fmt.Print(printstring)
	}

}
func PrintIfEnoughLevelAndPropulated(printstring string, requiredMinLevel int) {
	// TODO append \n at here
	if VerbosityLevel >= requiredMinLevel {
		if ElementInStringArray(printstring, []string{"", " "}){
			return
		}
		fmt.Print(printstring)
	}

}
/*
func SetVerbosityLevel(level int) {
	VerbosityLevel = level
}
*/
