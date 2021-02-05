package modules

import (
	"encoding/json"
	"fmt"
	"os"

	"../utils"
)

type PlatformDataStruct struct {
	Drivers           []string
	Files             []string
	NetworkInterfaces []string
	Processes         []string
}

type Configuration struct {
	VMware struct {
		PlatformData PlatformDataStruct
	}
	VirtualBox struct {
		PlatformData PlatformDataStruct
	}
	QEMU struct {
		PlatformData PlatformDataStruct
	}
	HyperV struct {
		platformData PlatformDataStruct
	}
	Parallels struct {
		PlatformData PlatformDataStruct
	}
	Other struct {
		PlatformData PlatformDataStruct
	}
	AnalysisTools struct {
		Processes []string
	}
}

func GetConfigfile(ConfigFile string) string {
	if ConfigFile == "" {
		ConfigFile = "resources//defaultconfig.json"
	}
	ConfigFile, err := utils.FindAbsolutePath(ConfigFile)
	if err != nil {
		utils.PrintIfEnoughLevel(fmt.Sprintf("[!] ER-CO001 Unable to find config file, %s \n", err), utils.OPERATION_ERROR_MESSAGE)
		os.Exit(2)
	}
	return ConfigFile
}
func ReadConfigFile(ConfigFile string) Configuration {
	bytesRead, err := utils.ReadFile(ConfigFile)
	if err != nil {
		utils.PrintIfEnoughLevel(fmt.Sprintf("[!] ER-CO002 Unable to read config file, %s \n", err), utils.OPERATION_ERROR_MESSAGE)
		os.Exit(5)
	}
	var Config Configuration
	err = json.Unmarshal(bytesRead, &Config)
	if err != nil {
		utils.PrintIfEnoughLevel(fmt.Sprintf("[!] ER-CO003 Unable to read config file, %s \n", err), utils.OPERATION_ERROR_MESSAGE)
		os.Exit(11)
	}
	return Config
}
