package utils

import (
	"encoding/json"
	"fmt"
	"os"
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
}

func GetConfigfile(ConfigFile string) string {
	if ConfigFile == "" {
		ConfigFile = "resources\\defaultconfig.json"
	}
	ConfigFile, err := FindAbsolutePath(ConfigFile)
	if err != nil {
		PrintIfEnoughLevel(fmt.Sprintf("[!] ER-CO001 Unable to find config file, %s \n", err), OPERATION_ERROR_MESSAGE)
		os.Exit(2)
	}
	return ConfigFile
}
func ReadConfigFile(ConfigFile string) Configuration {
	bytesRead, err := ReadFile(ConfigFile)
	if err != nil {
		PrintIfEnoughLevel(fmt.Sprintf("[!] ER-CO002 Unable to read config file, %s \n", err), OPERATION_ERROR_MESSAGE)
		os.Exit(5)
	}
	var Config Configuration
	err = json.Unmarshal(bytesRead, &Config)
	if err != nil {
		PrintIfEnoughLevel(fmt.Sprintf("[!] ER-CO003 Unable to read config file, %s \n", err), OPERATION_ERROR_MESSAGE)
		os.Exit(11)
	}
	return Config
}
