package utils

import (
	"fmt"
	"time"
	"strings"
	"strconv"
	"github.com/StackExchange/wmi"
)

type RAMDimm struct {
	Manufacturer string
	Model        string
	Size         int
	PartNumber   string
	Speed        int
	Voltage      int
}
type StorageDevice struct {
	Manufacturer string
	Model        string
	SerialNumber string
	Size         int
	Firmware     string
	Partitions   int
}
type CPU struct {
	Caption string
	Manufacturer string
	MaxClockSpeed int
	Name string
	Socket        string
}
type GPU struct {
	Manufacturer         string
	AdapterCompatibility string
	Caption              string
	Description	string
	Name string
	VideoProcessor string
	HorizontalResolution int
	VerticalResolution   int
}
type Mouse struct {
	Caption string
	Description string
	Manufacturer string
	Name string
}

type User struct {
	Name string
	Domain string
	Caption string
	SID	string
	AccountType int
}
type CDROM struct{
	Caption string
	Drive string
	Manufacturer string 
}
type OnBoardDevice struct{
	DeviceType int
	SerialNumber string
	Description string
}
type SystemInfo struct {
	Users          []User
	Uptime time.Time //TODO (incorrect value returned by WMI)
	TotalRamGB     int
	RAM            []RAMDimm
	OS             struct {
		// TODO (no relevant info observed at iteration 1)
		Name         string
		Version      string
		Architecture string
		InstallDate  time.Time
	}
	BIOS struct {
		Manufacturer string 		
		Name string
		SerialNumber string
		Version string 
	}
	MotherBoard struct {
		Manufacturer string
		Model        string
		Name         string
		SerialNumber string
		SKU          string
		Product      string
		OnBoardDevices []OnBoardDevice
	}
	TotalStorageGB int
	StorageDevices    []StorageDevice
	CPUs          []CPU
	GPUs          []GPU
	Monitors      []struct {
		//TODO (no relevant info observed at iteration 1)
	}
	Keyboards []struct {
		//TODO (no relevant info observed at iteration 1)
	}
	Mouses []Mouse
	AudioOutput []struct {
		//TODO (no relevant info observed at iteration 1)
	}
	SystemEnclosure []struct {
		//TODO (no relevant info observed at iteration 1)
	}
	PowerSupply []struct {
		//TODO (no relevant info observed at iteration 1)
	}
	USBControllers []struct {
		//TODO (no relevant info observed at iteration 1)
	}
	CDROMs []CDROM
}

func AppendIfSignificantlyPopulated(original string, printarray []string)string{
	str :=""
	for _, element := range printarray{
			if ElementInStringArray(element, []string{"Unknown","INVALID","(Standard disk drives)","(Standard system devices)","Default string","","\t","\t\t"," ","  ","   ","    ","     ","      ","       "}){
				return original
			}
			str = str + element			
		}
	return original + str
}

func PrintHardwareInfo(systemInfo *SystemInfo) {
	// RAM //
	if systemInfo.TotalRamGB !=0{
		PrintIfEnoughLevelAndPropulated(AppendIfSignificantlyPopulated("", []string{"TotalRAM: ", strconv.Itoa(systemInfo.TotalRamGB)," GB\n"}),HARDWARE_RECOGNITION_MESSAGE)
	}else{
		PrintIfEnoughLevelAndPropulated(AppendIfSignificantlyPopulated("", []string{"TotalRAM: ", strconv.Itoa(systemInfo.TotalRamGB)," GB -> OK, if Win32_PhysicalMemory returned nothing, this is a VM \n"}),HARDWARE_RECOGNITION_MESSAGE)
	}
	for i, tmp := range systemInfo.RAM{
		str := fmt.Sprintf("\tDimm %d -> ", i+1)
		str = AppendIfSignificantlyPopulated(str,[]string{"Manufacturer: ", tmp.Manufacturer, ", "})
		str = AppendIfSignificantlyPopulated(str,[]string{"Model: ", tmp.Model, ", "})
		str = AppendIfSignificantlyPopulated(str,[]string{"Size: ", strconv.Itoa(tmp.Size)," GB, "})
		str = AppendIfSignificantlyPopulated(str,[]string{"PartNumber: ", tmp.PartNumber, ", "})
		str = AppendIfSignificantlyPopulated(str,[]string{"Speed: ", strconv.Itoa(tmp.Speed), ", "})
		str = AppendIfSignificantlyPopulated(str,[]string{"Voltage: ", strconv.Itoa(tmp.Voltage), ", "})
		PrintIfEnoughLevel(fmt.Sprintf("%s \n",strings.Trim(str,", ")), HARDWARE_RECOGNITION_MESSAGE)
	}
	// STORAGE //
	PrintIfEnoughLevelAndPropulated(AppendIfSignificantlyPopulated("", []string{"Total Storage: ", strconv.Itoa(systemInfo.TotalStorageGB)," GB\n"}),HARDWARE_RECOGNITION_MESSAGE)
	for i, tmp := range systemInfo.StorageDevices{
		str := fmt.Sprintf("\tDevice %d -> ", i+1)
		str = AppendIfSignificantlyPopulated(str,[]string{"Manufacturer: ", tmp.Manufacturer, ", "})
		str = AppendIfSignificantlyPopulated(str,[]string{"Model: ", tmp.Model, ", "})
		str = AppendIfSignificantlyPopulated(str,[]string{"Size: ", strconv.Itoa(tmp.Size)," GB, "})
		str = AppendIfSignificantlyPopulated(str,[]string{"SerialNumber: ", tmp.SerialNumber, ", "})
		str = AppendIfSignificantlyPopulated(str,[]string{"Firmware: ", tmp.Firmware, ", "})
		//str = AppendIfSignificantlyPopulated(str,[]string{"Partitions: ", strconv.Itoa(tmp.Partitions), ", "})
		PrintIfEnoughLevel(fmt.Sprintf("%s \n",strings.Trim(str,", ")), HARDWARE_RECOGNITION_MESSAGE)
	}

	PrintIfEnoughLevelAndPropulated(AppendIfSignificantlyPopulated("", []string{"CPUs: ", strconv.Itoa(len(systemInfo.CPUs)),"\n"}),HARDWARE_RECOGNITION_MESSAGE)
	for i, tmp := range systemInfo.CPUs{
		str := fmt.Sprintf("\tCPU %d -> ", i+1)
		str = AppendIfSignificantlyPopulated(str,[]string{"Manufacturer: ", tmp.Manufacturer, ", "})
		str = AppendIfSignificantlyPopulated(str,[]string{"Caption: ", tmp.Caption, ", "})
		str = AppendIfSignificantlyPopulated(str,[]string{"MaxClockSpeed: ", strconv.Itoa(tmp.MaxClockSpeed), ", "})
		str = AppendIfSignificantlyPopulated(str,[]string{"Name: ", tmp.Name, ", "})
		str = AppendIfSignificantlyPopulated(str,[]string{"Socket: ", tmp.Socket, ", "})
		PrintIfEnoughLevel(fmt.Sprintf("%s \n",strings.Trim(str,", ")), HARDWARE_RECOGNITION_MESSAGE)
	}
	PrintIfEnoughLevelAndPropulated(AppendIfSignificantlyPopulated("", []string{"GPUs: ", strconv.Itoa(len(systemInfo.CPUs)),"\n"}),HARDWARE_RECOGNITION_MESSAGE)
	for i, tmp := range systemInfo.GPUs{
		resolution :=""
		if tmp.HorizontalResolution != 0 && tmp.VerticalResolution != 0 {
			resolution = strconv.Itoa(tmp.HorizontalResolution) + "x" + strconv.Itoa(tmp.VerticalResolution)
		}
		str := fmt.Sprintf("\tGPU %d -> ", i+1)
		str = AppendIfSignificantlyPopulated(str,[]string{"Manufacturer: ", tmp.Manufacturer, ", "})
		str = AppendIfSignificantlyPopulated(str,[]string{"Caption: ", tmp.Caption, ", "})
		str = AppendIfSignificantlyPopulated(str,[]string{"AdapterCompatibility: ", tmp.AdapterCompatibility, ", "})
		str = AppendIfSignificantlyPopulated(str,[]string{"Name: ", tmp.Name, ", "})
		str = AppendIfSignificantlyPopulated(str,[]string{"VideoProcessor: ", tmp.VideoProcessor, ", "})
		str = AppendIfSignificantlyPopulated(str,[]string{"Description: ", tmp.Description, ", "})
		str = AppendIfSignificantlyPopulated(str,[]string{"Resolution: ", resolution, ", "})
		PrintIfEnoughLevel(fmt.Sprintf("%s \n",strings.Trim(str,", ")), HARDWARE_RECOGNITION_MESSAGE)
	}
	PrintIfEnoughLevelAndPropulated(AppendIfSignificantlyPopulated("", []string{"Mouses: ", strconv.Itoa(len(systemInfo.Mouses)),"\n"}),HARDWARE_RECOGNITION_MESSAGE)
	for i, tmp := range systemInfo.Mouses{
		str := fmt.Sprintf("\tMouse %d -> ", i+1)
		str = AppendIfSignificantlyPopulated(str,[]string{"Manufacturer: ", tmp.Manufacturer,", "})
		str = AppendIfSignificantlyPopulated(str,[]string{"Caption: ", tmp.Caption,", "})
		str = AppendIfSignificantlyPopulated(str,[]string{"Description: ", tmp.Description,", "})
		str = AppendIfSignificantlyPopulated(str,[]string{"Name: ", tmp.Name,", "})
		PrintIfEnoughLevel(fmt.Sprintf("%s \n",strings.Trim(str,", ")), HARDWARE_RECOGNITION_MESSAGE)
	}
	PrintIfEnoughLevelAndPropulated(AppendIfSignificantlyPopulated("", []string{"Users: ", strconv.Itoa(len(systemInfo.Users)),"\n"}),HARDWARE_RECOGNITION_MESSAGE)
	for i, tmp := range systemInfo.Users{
		str := fmt.Sprintf("\tUser %d -> ", i+1)
		str = AppendIfSignificantlyPopulated(str,[]string{"Name: ", tmp.Name,", "})
		str = AppendIfSignificantlyPopulated(str,[]string{"Domain: ", tmp.Domain,", "})
		str = AppendIfSignificantlyPopulated(str,[]string{"Caption: ", tmp.Caption,", "})
		str = AppendIfSignificantlyPopulated(str,[]string{"AccountType: ", strconv.Itoa(tmp.AccountType), ", "})
		str = AppendIfSignificantlyPopulated(str,[]string{"SID: ", tmp.SID,", "})
		PrintIfEnoughLevel(fmt.Sprintf("%s \n",strings.Trim(str,", ")), HARDWARE_RECOGNITION_MESSAGE)
	}
	PrintIfEnoughLevelAndPropulated(AppendIfSignificantlyPopulated("", []string{"CDROMs: ", strconv.Itoa(len(systemInfo.CDROMs)),"\n"}),HARDWARE_RECOGNITION_MESSAGE)
	for i, tmp := range systemInfo.CDROMs{
		str := fmt.Sprintf("\tCDROM %d -> ", i+1)
		str = AppendIfSignificantlyPopulated(str,[]string{"Manufacturer: ", tmp.Manufacturer,", "})
		str = AppendIfSignificantlyPopulated(str,[]string{"Caption: ", tmp.Caption,", "})
		str = AppendIfSignificantlyPopulated(str,[]string{"Drive: ", tmp.Drive,", "})
		PrintIfEnoughLevel(fmt.Sprintf("%s \n",strings.Trim(str,", ")), HARDWARE_RECOGNITION_MESSAGE)
	}
	str := fmt.Sprintf("MotherBoard -> ")
	str = AppendIfSignificantlyPopulated(str,[]string{"Manufacturer: ", systemInfo.MotherBoard.Manufacturer,", "})
	str = AppendIfSignificantlyPopulated(str,[]string{"Product: ", systemInfo.MotherBoard.Product,", "})
	str = AppendIfSignificantlyPopulated(str,[]string{"Name: ", systemInfo.MotherBoard.Name,", "})
	str = AppendIfSignificantlyPopulated(str,[]string{"Model: ", systemInfo.MotherBoard.Model,", "})
	str = AppendIfSignificantlyPopulated(str,[]string{"SerialNumber: ", systemInfo.MotherBoard.SerialNumber,", "})
	str = AppendIfSignificantlyPopulated(str,[]string{"SKU: ", systemInfo.MotherBoard.SKU,", "})	
	PrintIfEnoughLevel(fmt.Sprintf("%s \n",strings.Trim(str,", ")), HARDWARE_RECOGNITION_MESSAGE)
	for i, tmp := range systemInfo.MotherBoard.OnBoardDevices{
		str := fmt.Sprintf("\tOnBoardDevice %d -> ", i+1)
		str = AppendIfSignificantlyPopulated(str,[]string{"DeviceType: ", strconv.Itoa(tmp.DeviceType),", "})
		str = AppendIfSignificantlyPopulated(str,[]string{"SerialNumber: ", tmp.SerialNumber,", "})
		str = AppendIfSignificantlyPopulated(str,[]string{"Description: ", tmp.Description,", "})
		PrintIfEnoughLevel(fmt.Sprintf("%s \n",strings.Trim(str,", ")), HARDWARE_RECOGNITION_MESSAGE)
	}
	str = fmt.Sprintf("BIOS -> ")
	str = AppendIfSignificantlyPopulated(str,[]string{"Manufacturer: ", systemInfo.BIOS.Manufacturer,", "})
	str = AppendIfSignificantlyPopulated(str,[]string{"Name: ", systemInfo.BIOS.Name,", "})
	str = AppendIfSignificantlyPopulated(str,[]string{"SerialNumber: ", systemInfo.BIOS.SerialNumber,", "})
	str = AppendIfSignificantlyPopulated(str,[]string{"Version: ", systemInfo.BIOS.Version,", "})
	PrintIfEnoughLevel(fmt.Sprintf("%s \n",strings.Trim(str,", ")), HARDWARE_RECOGNITION_MESSAGE)
	
}

func PopulateRamInfo(systemInfo *SystemInfo) {
	type Win32_PhysicalMemory struct {
		Manufacturer string
		Model        string
		Capacity     uint64
		MaxVoltage   uint32
		PartNumber   string
		Speed        int
	}
	var dst []Win32_PhysicalMemory
	q := wmi.CreateQuery(&dst, "")
	err := wmi.Query(q, &dst)
	if err != nil {
		PrintIfEnoughLevel(fmt.Sprintf("%s [!] ER-HU001 Error querying WMI: %s\n", "", err), OPERATION_ERROR_MESSAGE)
	}
	systemInfo.TotalRamGB = 0
	for _, element := range dst {
		var tmp RAMDimm
		tmp.Size = int(element.Capacity) / 1073741824
		tmp.Manufacturer = element.Manufacturer
		tmp.Model = element.Model
		tmp.Voltage = int(element.MaxVoltage)
		tmp.PartNumber =  strings.Trim(element.PartNumber," ")
		tmp.Speed = element.Speed
		systemInfo.TotalRamGB += tmp.Size
		systemInfo.RAM = append(systemInfo.RAM, tmp)
	}
}
func PopulateMouses(systemInfo *SystemInfo) {
	type Win32_PointingDevice struct{
		Caption string
		Description string
		Manufacturer string
		Name string
	}
	// TODO, search more significant variables
	var dst []Win32_PointingDevice
	q := wmi.CreateQuery(&dst, "")
	err := wmi.Query(q, &dst)
	if err != nil {
		PrintIfEnoughLevel(fmt.Sprintf("%s [!] ER-HU002 Error querying WMI: %s\n", "", err), OPERATION_ERROR_MESSAGE)
	}
	for _, element := range dst {
		var tmp Mouse
		tmp.Caption = element.Caption
		tmp.Description = element.Description
		tmp.Manufacturer = element.Manufacturer
		tmp.Name = element.Name
		systemInfo.Mouses = append(systemInfo.Mouses, tmp)
	}
}

func PopulateUsers(systemInfo *SystemInfo) {
	type Win32_UserAccount struct{
		Name string
		Domain string
		Caption string
		SID	string
		AccountType uint32
	}
	// TODO, search more significant variables
	var dst []Win32_UserAccount
	q := wmi.CreateQuery(&dst, "")
	err := wmi.Query(q, &dst)
	if err != nil {
		PrintIfEnoughLevel(fmt.Sprintf("%s [!] ER-HU003 Error querying WMI: %s\n", "", err), OPERATION_ERROR_MESSAGE)
	}
	for _, element := range dst {
		var tmp User
		tmp.Name = element.Name
		tmp.Domain = element.Domain
		tmp.Caption = element.Caption
		tmp.SID = element.SID
		tmp.AccountType = int(element.AccountType)
		systemInfo.Users = append(systemInfo.Users, tmp)
	}
}

func PopulateCDROMs(systemInfo *SystemInfo) {
	type Win32_CDROMDrive struct{
		Caption string
		Drive string
		Manufacturer string 
	}
	// TODO, search more significant variables
	var dst []Win32_CDROMDrive
	q := wmi.CreateQuery(&dst, "")
	err := wmi.Query(q, &dst)
	if err != nil {
		PrintIfEnoughLevel(fmt.Sprintf("%s [!] ER-HU004 Error querying WMI: %s\n", "", err), OPERATION_ERROR_MESSAGE)
	}
	for _, element := range dst {
		var tmp CDROM
		tmp.Drive = element.Drive
		tmp.Manufacturer = element.Manufacturer
		tmp.Caption = element.Caption
		systemInfo.CDROMs = append(systemInfo.CDROMs, tmp)
	}
}

func PopulateGPUs(systemInfo *SystemInfo) {
	type Win32_VideoController struct{
		AdapterCompatibility string
		Caption              string
		Description	string
		Name string
		VideoProcessor string
		CurrentHorizontalResolution uint32
		CurrentVerticalResolution   uint32
	}
	// TODO, search more significant variables
	var dst []Win32_VideoController
	q := wmi.CreateQuery(&dst, "")
	err := wmi.Query(q, &dst)
	if err != nil {
		PrintIfEnoughLevel(fmt.Sprintf("%s [!] ER-HU005 Error querying WMI: %s\n", "", err), OPERATION_ERROR_MESSAGE)
	}
	for _, element := range dst {
		var tmp GPU
		//tmp.Manufacturer = element.Manufacturer
		tmp.AdapterCompatibility = element.AdapterCompatibility
		tmp.Caption = element.Caption
		tmp.Description = element.Description
		tmp.Name = element.Name
		tmp.VideoProcessor = element.VideoProcessor
		tmp.HorizontalResolution = int(element.CurrentHorizontalResolution)
		tmp.VerticalResolution = int(element.CurrentVerticalResolution)
		systemInfo.GPUs = append(systemInfo.GPUs, tmp)
	}
}


func PopulateStorageInfo(systemInfo *SystemInfo) {
	type Win32_DiskDrive struct {
		Manufacturer string
		Model        string
		SerialNumber string
		Size     uint64
		Partitions   uint32
		FirmwareRevision     string
	}
	var dst []Win32_DiskDrive
	q := wmi.CreateQuery(&dst, "")
	err := wmi.Query(q, &dst)
	if err != nil {
		PrintIfEnoughLevel(fmt.Sprintf("%s [!] ER-HU006 Error querying WMI: %s\n", "", err), OPERATION_ERROR_MESSAGE)
	}
	systemInfo.TotalStorageGB = 0
	for _, element := range dst {
		var tmp StorageDevice
		tmp.Size = int(element.Size) / 1073741824
		tmp.Manufacturer = element.Manufacturer
		tmp.Model = element.Model
		tmp.SerialNumber = strings.Trim(element.SerialNumber," ")
		tmp.Firmware = element.FirmwareRevision
		tmp.Partitions = int(element.Partitions)
		systemInfo.TotalStorageGB += tmp.Size
		systemInfo.StorageDevices = append(systemInfo.StorageDevices, tmp)
	}

}

func PopulateCPUsInfo(systemInfo *SystemInfo){
	type Win32_Processor struct{
		Caption string
		Manufacturer string
		MaxClockSpeed uint32
		Name string
	}
	var dst []Win32_Processor
	q := wmi.CreateQuery(&dst, "")
	err := wmi.Query(q, &dst)
	if err != nil {
		PrintIfEnoughLevel(fmt.Sprintf("%s [!] ER-HU007 Error querying WMI: %s\n", "", err), OPERATION_ERROR_MESSAGE)
	}
	for _, element := range dst {
		var tmp CPU
		tmp.Caption = element.Caption
		tmp.Manufacturer = element.Manufacturer
		tmp.MaxClockSpeed = int(element.MaxClockSpeed)
		tmp.Name = strings.Trim(element.Name," ")
		systemInfo.CPUs = append(systemInfo.CPUs, tmp)
	}
}

func PopulateHardBIOSInfo(systemInfo *SystemInfo){
	type CIM_BIOSElement struct{
		Manufacturer string 		
		Name string
		SerialNumber string
		Version string 
	}
	var dst []CIM_BIOSElement
	q := wmi.CreateQuery(&dst, "")
	err := wmi.Query(q, &dst)
	if err != nil {
		PrintIfEnoughLevel(fmt.Sprintf("%s [!] ER-HU008 Error querying WMI: %s\n", "", err), OPERATION_ERROR_MESSAGE)
	}
	systemInfo.BIOS.Manufacturer=dst[0].Manufacturer
	systemInfo.BIOS.Name=dst[0].Name
	systemInfo.BIOS.SerialNumber=dst[0].SerialNumber
	systemInfo.BIOS.Version=dst[0].Version
}

func PopulateMotherBoardInfo(systemInfo *SystemInfo){
	type Win32_BaseBoard struct {
		Manufacturer string
		Model        string
		Name         string
		SerialNumber string
		SKU          string
		Product      string
	}
	type Win32_OnboardDevice struct{
		DeviceType int
		SerialNumber string
		Description string
	}
	var dst []Win32_BaseBoard
	q := wmi.CreateQuery(&dst, "")
	err := wmi.Query(q, &dst)
	if err != nil {
		PrintIfEnoughLevel(fmt.Sprintf("%s [!] ER-HU009 Error querying WMI: %s\n", "", err), OPERATION_ERROR_MESSAGE)
	}
	systemInfo.MotherBoard.Manufacturer=dst[0].Manufacturer
	systemInfo.MotherBoard.Model=dst[0].Model
	systemInfo.MotherBoard.Name=dst[0].Name
	systemInfo.MotherBoard.SerialNumber=dst[0].SerialNumber
	systemInfo.MotherBoard.SKU=dst[0].SKU
	systemInfo.MotherBoard.Product=dst[0].Product
	// ON board Devices
	var dst2 []Win32_OnboardDevice
	q = wmi.CreateQuery(&dst2, "")
	err = wmi.Query(q, &dst2)
	if err != nil {
		PrintIfEnoughLevel(fmt.Sprintf("%s [!] ER-HU010 Error querying WMI: %s\n", "", err), OPERATION_ERROR_MESSAGE)
	}
	for _, element := range dst2 {
		var tmp OnBoardDevice
		tmp.SerialNumber = element.SerialNumber
		tmp.Description = element.Description
		tmp.DeviceType = element.DeviceType
		systemInfo.MotherBoard.OnBoardDevices = append(systemInfo.MotherBoard.OnBoardDevices, tmp)
	}
}