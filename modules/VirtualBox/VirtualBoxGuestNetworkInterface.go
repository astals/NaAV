package VirtualBox

//Set-NetAdapter -Name "NaAV VBox" -MacAddress "" -Confirm:$false
import (
	"fmt"

	"../../utils"
)

func InstallVirtualBoxFakeGuestNetworkInterface() {
	_ = utils.CreateNetworkAdapter("NaAV VBox", "08:00:27:e1:ee:e7", "\t")
}

func UninstallVirtualBoxFakeGuestNetworkInterface() {
	_ = utils.DeleteNetworkAdapter("NaAV VBox", "\t")
}

func CheckVirtualBoxFakeGuestNetworkInterface() {
	fmt.Printf("Checking VirtualBox host drivers on system:\n")
	utils.ExistsNetworkAdapter("08:00:27:e1:ee:e7", "\t")
}
