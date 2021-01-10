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
	ConfigFile, err := utils.FindAbsolutePath(ConfigFile)
	if err != nil {
		fmt.Printf("[!] ER-IN001 Unable to find config file")
		os.Exit(2)
	}
	Config := ReadConfigFile(ConfigFile)
	cloneInstallFiles(ConfigFile)
	if Config.WMware.FakeDrivers {
		VMware.InstallVMwareDrivers()
	}
	if Config.VirtualBox.FakeVirtualBoxGuestAdditionsFiles {
		VirtualBox.InstallVirtualBoxGuestAdditionsFiles()
	}
}

func ReadConfigFile(ConfigFile string) Configuration {
	bytesRead, err := utils.ReadFile(ConfigFile)
	if err != nil {
		fmt.Printf("[!] ER-IN002 Unable to read config file, %s", err)
		os.Exit(5)
	}
	var Config Configuration
	err = json.Unmarshal(bytesRead, &Config)
	if err != nil {
		fmt.Printf("[!] ER-IN003 Unable to read config file, %s", err)
		os.Exit(11)
	}
	return Config
}

func cloneInstallFiles(ConfigFile string) {
	cwd, _ := os.Getwd()
	utils.CreateFoldersPath("C:\\Program Files (x86)\\NaAV")
	args := os.Args
	err := utils.CopyFile(cwd+"\\"+args[1], "C:\\Program Files (x86)\\NaAV\\naav.exe")
	if err != nil {
		fmt.Printf("\t [!] ER-IN004 Unable to save %s , %s", "C:\\Program Files (x86)\\NaAV\\naav.exe", err)
	}
	err = utils.CopyFile(ConfigFile, "C:\\Program Files (x86)\\NaAV\\config.json")
	if err != nil {
		fmt.Printf("\t [!] ER-IN004 Unable to save %s , %s", "C:\\Program Files (x86)\\NaAV\\config.json", err)
	}
}
