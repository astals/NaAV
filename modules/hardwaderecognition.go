package modules

/*
github.com/StackExchange/wmi
https://github.com/jaypipes/ghw
https://pkg.go.dev/golang.org/x/sys
https://github.com/osquery
https://github.com/digitalocean/go-smbios
COM
github.com/go-ole/go-ole
---------------------------------------------------------
https://underc0de.org/foro/visual-studio-net/(vb-net)-anti-virtual-machine/
https://docs.microsoft.com/en-us/windows/win32/wmisdk/wmi-providers
*/

/*
PENDING ADD?
	CPU Temp
*/

import (
	"../utils"
)

func CheckHardware() {
	var systemInfo utils.SystemInfo
	utils.PopulateRamInfo(&systemInfo)
	utils.PopulateStorageInfo(&systemInfo)
	utils.PopulateHardBIOSInfo(&systemInfo)
	utils.PopulateMotherBoardInfo(&systemInfo)
	utils.PopulateCPUsInfo(&systemInfo)
	utils.PopulateMouses(&systemInfo)
	utils.PopulateGPUs(&systemInfo)
	utils.PopulateUsers(&systemInfo)
	utils.PopulateCDROMs(&systemInfo)	
	utils.PrintHardwareInfo(&systemInfo)
}
