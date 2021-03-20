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
type HardDrive struct {
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

type SystemInfo struct {
	users          []string
	LastBootUpTime time.Time
	TotalRamGB     int
	RAM            []RAMDimm
	SwapSize       int
	OS             struct {
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
	Motherboard struct {
		Manufacturer string
		Model        string
		Name         string
		SerialNumber string
		SKU          string
		Product      string
	}
	TotalDrivesGB int
	HardDrives    []HardDrive
	CPUs          []CPU
	GPUs          []GPU
	Monitors      []struct {
	}
	Keyboards []struct {
	}
	Mouses []Mouse
	AudioOutput []struct {
	}
}

func AppendIfSignificantlyPopulated(original string, printarray []string)string{
	str :=""
	for _, element := range printarray{
			if ElementInStringArray(element, []string{"", " ", "Unknown","(Standard disk drives)","(Standard system devices)","Default string"}){
				return original
			}
			str = str + element			
		}
	return original + str
}

func PrintHardwareInfo(systemInfo *SystemInfo) {
	// RAM //
	PrintIfEnoughLevelAndPropulated(AppendIfSignificantlyPopulated("", []string{"TotalRAM: ", strconv.Itoa(systemInfo.TotalRamGB)," GB\n"}),HARDWARE_RECOGNITION_MESSAGE)
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
	PrintIfEnoughLevelAndPropulated(AppendIfSignificantlyPopulated("", []string{"Total HardDrive Storage: ", strconv.Itoa(systemInfo.TotalDrivesGB)," GB\n"}),HARDWARE_RECOGNITION_MESSAGE)
	for i, tmp := range systemInfo.HardDrives{
		str := fmt.Sprintf("\tHDD %d -> ", i+1)
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
	str := fmt.Sprintf("BIOS -> ")
	str = AppendIfSignificantlyPopulated(str,[]string{"Manufacturer: ", systemInfo.BIOS.Manufacturer,", "})
	str = AppendIfSignificantlyPopulated(str,[]string{"Name: ", systemInfo.BIOS.Name,", "})
	str = AppendIfSignificantlyPopulated(str,[]string{"SerialNumber: ", systemInfo.BIOS.SerialNumber,", "})
	str = AppendIfSignificantlyPopulated(str,[]string{"Version: ", systemInfo.BIOS.Version,", "})
	PrintIfEnoughLevel(fmt.Sprintf("%s \n",strings.Trim(str,", ")), HARDWARE_RECOGNITION_MESSAGE)
	str = fmt.Sprintf("MotherBoard -> ")
	str = AppendIfSignificantlyPopulated(str,[]string{"Manufacturer: ", systemInfo.Motherboard.Manufacturer,", "})
	str = AppendIfSignificantlyPopulated(str,[]string{"Product: ", systemInfo.Motherboard.Product,", "})
	str = AppendIfSignificantlyPopulated(str,[]string{"Name: ", systemInfo.Motherboard.Name,", "})
	str = AppendIfSignificantlyPopulated(str,[]string{"Model: ", systemInfo.Motherboard.Model,", "})
	str = AppendIfSignificantlyPopulated(str,[]string{"SerialNumber: ", systemInfo.Motherboard.SerialNumber,", "})
	str = AppendIfSignificantlyPopulated(str,[]string{"SKU: ", systemInfo.Motherboard.SKU,", "})	
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
	// TODO: handle err
	fmt.Print(err)
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
	// TODO: handle err
	fmt.Print(err)
	for _, element := range dst {
		var tmp Mouse
		tmp.Caption = element.Caption
		tmp.Description = element.Description
		tmp.Manufacturer = element.Manufacturer
		tmp.Name = element.Name
		systemInfo.Mouses = append(systemInfo.Mouses, tmp)
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
	// TODO: handle err
	fmt.Print(err)
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


func PopulateHardDrivesInfo(systemInfo *SystemInfo) {
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
	// TODO: handle err
	fmt.Print(err)
	systemInfo.TotalDrivesGB = 0
	for _, element := range dst {
		var tmp HardDrive
		tmp.Size = int(element.Size) / 1073741824
		tmp.Manufacturer = element.Manufacturer
		tmp.Model = element.Model
		tmp.SerialNumber = strings.Trim(element.SerialNumber," ")
		tmp.Firmware = element.FirmwareRevision
		tmp.Partitions = int(element.Partitions)
		systemInfo.TotalDrivesGB += tmp.Size
		systemInfo.HardDrives = append(systemInfo.HardDrives, tmp)
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
	// TODO: handle err
	fmt.Print(err)
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
	fmt.Print(err)
	// TODO handle err
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
	var dst []Win32_BaseBoard
	q := wmi.CreateQuery(&dst, "")
	err := wmi.Query(q, &dst)
	fmt.Print(err)
	systemInfo.Motherboard.Manufacturer=dst[0].Manufacturer
	systemInfo.Motherboard.Model=dst[0].Model
	systemInfo.Motherboard.Name=dst[0].Name
	systemInfo.Motherboard.SerialNumber=dst[0].SerialNumber
	systemInfo.Motherboard.SKU=dst[0].SKU
	systemInfo.Motherboard.Product=dst[0].Product
}



