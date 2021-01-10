package modes

import (
	"encoding/json"
	"fmt"
	"os"

	"../modules/VMware"
	"../modules/VirtualBox"
	"../utils"
)

type Configuration struct {
	WMware struct {
		FakeDrivers                bool
		FakeGuestNetworkInterfaces bool
		FakeProcesses              []string
	}
	VirtualBox struct {
		FakeVirtualBoxGuestAdditionsFiles bool
		FakeGuestNetworkInterfaces        bool
		FakeProcesses                     []string
	}
	AnalysisTools struct {
		FakeProcesses []string
	}
}

func Install(ConfigFile string) {
	Config := ReadConfigFile(ConfigFile)
	if Config.WMware.FakeDrivers {
		VMware.InstallVMwareDrivers()
	}
	if Config.VirtualBox.FakeVirtualBoxGuestAdditionsFiles {
		VirtualBox.InstallVirtualBoxGuestAdditionsFiles()
	}
}

func ReadConfigFile(ConfigFile string) Configuration {
	var bytesRead []byte
	var err error
	/* read file */
	res, _ := utils.FileExists(ConfigFile)
	if res == true {
		bytesRead, err = utils.ReadFile(ConfigFile)
		if err != nil {
			fmt.Printf("Unable to read config file, %s", err)
			os.Exit(5)
		}
	} else {
		cwd, _ := os.Getwd()
		ConfigFile = cwd + "\\" + ConfigFile
		res, _ = utils.FileExists(ConfigFile)
		if res == true {
			bytesRead, err = utils.ReadFile(ConfigFile)
			if err != nil {
				fmt.Printf("Unable to read config file, %s", err)
				os.Exit(5)
			}
		} else {
			fmt.Printf("Unable to read config file, %s", err)
			os.Exit(2)
		}
	}
	/* unmarshall json */
	var Config Configuration
	err = json.Unmarshal(bytesRead, &Config)
	if err != nil {
		fmt.Printf("Unable to read config file, %s", err)
		os.Exit(11)
	}
	return Config
}
