package utils

import (
	"encoding/json"
	"fmt"
	"os"
	"embed"
)

type PlatformDataStruct struct {
	Drivers  []string
	Files    []string
	Registry struct {
		KeysAndValueNames      map[string][]string
		OrderedKeysParentTrees []string // used for check and uninstall
	}
	NetworkInterfaces map[string]string
	Processes         []string
}

type Configuration struct {
	AnalysisTools []string
	HyperV        PlatformDataStruct
	Other         PlatformDataStruct
	Parallels     PlatformDataStruct
	QEMU          PlatformDataStruct
	VMware        PlatformDataStruct
	VirtualBox    PlatformDataStruct
	Sandboxie 	PlatformDataStruct
}


func LoadConfigFromFile(ConfigFile string) Configuration {
	ConfigFile, err := FindAbsolutePath(ConfigFile)
	if err != nil {
		PrintIfEnoughLevel(fmt.Sprintf("[!] ER-CO001 Unable to find config file, %s \n", err), OPERATION_ERROR_MESSAGE)
		os.Exit(2)
	}
	bytesRead, err := ReadFile(ConfigFile)
	if err != nil {
		PrintIfEnoughLevel(fmt.Sprintf("[!] ER-CO002 Unable to read config file, %s \n", err), OPERATION_ERROR_MESSAGE)
		os.Exit(5)
	}
	return LoadConfig(bytesRead)
}

func LoadConfigFromEmbededFS(ConfigFile embed.FS)Configuration{
	bytesRead, err := ConfigFile.ReadFile("resources/defaultconfig.json")
	if err != nil {
		PrintIfEnoughLevel(fmt.Sprintf("[!] ER-CO003 Unable to read config file, %s \n", err), OPERATION_ERROR_MESSAGE)
		os.Exit(5)
	}
	return LoadConfig(bytesRead)
}

func LoadConfig(bytesRead []byte) Configuration  {
	var Config Configuration
	err := json.Unmarshal(bytesRead, &Config)
	if err != nil {
		PrintIfEnoughLevel(fmt.Sprintf("[!] ER-CO004 Unable to read config file, %s \n", err), OPERATION_ERROR_MESSAGE)
		os.Exit(11)
	}
	return Config
}

func JoinAllProgramNames(config Configuration) []string {
	all_processes := append(config.AnalysisTools, config.HyperV.Processes...)
	all_processes = append(all_processes, config.Other.Processes...)
	all_processes = append(all_processes, config.Parallels.Processes...)
	all_processes = append(all_processes, config.QEMU.Processes...)
	all_processes = append(all_processes, config.VMware.Processes...)
	all_processes = append(all_processes, config.VirtualBox.Processes...)
	return all_processes
}
