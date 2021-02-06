package modes

import (
	"fmt"
	"os"

	"../modules"
	"../utils"
)

func Install(ConfigFile string) {
	Config := utils.ReadConfigFile(ConfigFile)
	cloneInstallFiles(ConfigFile)
	modules.InstallFiles(Config.VirtualBox.Files, "guest files", "Virtual Box")
	modules.InstallFiles(Config.VMware.Files, "guest files", "VMware")
	modules.InstallFiles(Config.VirtualBox.Drivers, "drivers", "Virtual Box")
	modules.InstallFiles(Config.VMware.Drivers, "drivers", "VMware")
	modules.InstallRegkeys(Config.VirtualBox.Registry.KeysAndValueNames, "Virtual Box")
	modules.InstallRegkeys(Config.VMware.Registry.KeysAndValueNames, "VMware")
	modules.InstallNetworkInterfaces(Config.VirtualBox.NetworkInterfaces, "Virtual Box")
	modules.InstallNetworkInterfaces(Config.VMware.NetworkInterfaces, "VMware")
	modules.InstallNetworkInterfaces(Config.HyperV.NetworkInterfaces, "Hyper-V")
	modules.InstallNetworkInterfaces(Config.Parallels.NetworkInterfaces, "Parallels")
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
