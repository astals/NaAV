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
	modules.InstallFiles(Config.VMware.Files, "guest files", "VMware")
	modules.InstallFiles(Config.VMware.Drivers, "drivers", "VMware")
	modules.InstallFiles(Config.VirtualBox.Files, "guest files", "Virtual Box")
	modules.InstallFiles(Config.VirtualBox.Drivers, "drivers", "Virtual Box")
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
