package modes

import (
	"encoding/json"
	"fmt"
	"os"

	"../modules/Others"
	"../modules/VMware"
	"../modules/VirtualBox"
	"../utils"
)

type Configuration struct {
	WMware struct {
		FakeGuestDrivers           bool
		FakeGuestFiles             bool
		FakeGuestNetworkInterfaces bool
		FakeProcesses              []string
	}
	VirtualBox struct {
		FakeGuestDrivers                  bool
		FakeVirtualBoxGuestAdditionsFiles bool
		FakeGuestNetworkInterfaces        bool
		FakeProcesses                     []string
	}
	AnalysisTools struct {
		FakeProcesses []string
	}
	OtherRegistryKeys bool
}

func Install(ConfigFile string) {
	ConfigFile, err := utils.FindAbsolutePath(ConfigFile)
	if err != nil {
		fmt.Printf("[!] ER-IN001 Unable to find config file \n")
		os.Exit(2)
	}
	Config := ReadConfigFile(ConfigFile)
	cloneInstallFiles(ConfigFile)
	if Config.WMware.FakeGuestDrivers {
		VMware.InstallVMwareDrivers()
	}
	if Config.WMware.FakeGuestFiles {
		VMware.InstallVMwareGuestFiles()
	}
	if Config.VirtualBox.FakeGuestDrivers {
		VirtualBox.InstallVirtualBoxDrivers()
	}
	if Config.VirtualBox.FakeVirtualBoxGuestAdditionsFiles {
		VirtualBox.InstallVirtualBoxGuestAdditionsFiles()
	}
	if Config.OtherRegistryKeys {
		Others.InstallOtherRegistryKeys()
	}

}

func ReadConfigFile(ConfigFile string) Configuration {
	bytesRead, err := utils.ReadFile(ConfigFile)
	if err != nil {
		fmt.Printf("[!] ER-IN002 Unable to read config file, %s \n", err)
		os.Exit(5)
	}
	var Config Configuration
	err = json.Unmarshal(bytesRead, &Config)
	if err != nil {
		fmt.Printf("[!] ER-IN003 Unable to read config file, %s \n", err)
		os.Exit(11)
	}
	return Config
}

func cloneInstallFiles(ConfigFile string) {
	cwd, _ := os.Getwd()
	fmt.Printf("Saving files \n")
	utils.CreateFoldersPath("C:\\Program Files (x86)\\NaAV")
	args := os.Args
	err := utils.CopyFile(cwd+"\\"+args[1], "C:\\Program Files (x86)\\NaAV\\naav.exe")
	if err != nil {
		fmt.Printf("\t [!] ER-IN004 Unable to save %s , %s \n", "C:\\Program Files (x86)\\NaAV\\naav.exe", err)
	}
	err = utils.CopyFile(ConfigFile, "C:\\Program Files (x86)\\NaAV\\config.json")
	if err != nil {
		fmt.Printf("\t [!] ER-IN004 Unable to save %s , %s \n", "C:\\Program Files (x86)\\NaAV\\config.json", err)
	}
}
