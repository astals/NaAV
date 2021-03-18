package utils

import (
	"fmt"
	"time"
	"strings"

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
	MaxClockSpeed uint32
	Name string
	Socket        string
}
type GPU struct {
	Manufacturer         string
	GPUName              string
	RefreshRate          int
	HorizontalResolution int
	VerticalResolution   int
	Controller           string
	PartNumber           string
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
		//debug shit
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
	fmt.Print(q)
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
		tmp.MaxClockSpeed = element.MaxClockSpeed
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



