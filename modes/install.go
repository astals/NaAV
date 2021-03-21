package modes

import (
	"fmt"
	"os"
	"strings"

	"../modules"
	"../utils"
)

func Install(ConfigFile string) {
	Config := utils.ReadConfigFile(ConfigFile)
	cloneInstallFiles(ConfigFile)
	utils.PrintIfEnoughLevel("=====> Virtual Box <=====\n", utils.BASIC_INFORMATION_MESSAGE)
	modules.InstallFiles(Config.VirtualBox.Files, "guest files", "Virtual Box")
	modules.InstallFiles(Config.VirtualBox.Drivers, "drivers", "Virtual Box")
	modules.InstallRegkeys(Config.VirtualBox.Registry.KeysAndValueNames, "Virtual Box")
	modules.InstallNetworkInterfaces(Config.VirtualBox.NetworkInterfaces, "Virtual Box")
	
	utils.PrintIfEnoughLevel("=====> VMware <=====\n", utils.BASIC_INFORMATION_MESSAGE)
	modules.InstallFiles(Config.VMware.Files, "guest files", "VMware")
	modules.InstallFiles(Config.VMware.Drivers, "drivers", "VMware")
	modules.InstallRegkeys(Config.VMware.Registry.KeysAndValueNames, "VMware")
	modules.InstallNetworkInterfaces(Config.VMware.NetworkInterfaces, "VMware")
	
	utils.PrintIfEnoughLevel("=====> Hyper-V <=====\n", utils.BASIC_INFORMATION_MESSAGE)
	modules.InstallNetworkInterfaces(Config.HyperV.NetworkInterfaces, "Hyper-V")
	
	utils.PrintIfEnoughLevel("=====> Parallels <=====\n", utils.BASIC_INFORMATION_MESSAGE)
	modules.InstallNetworkInterfaces(Config.Parallels.NetworkInterfaces, "Parallels")
	
	utils.PrintIfEnoughLevel("=====> FakeProgramSpawner service <=====\n", utils.BASIC_INFORMATION_MESSAGE)
	all_processes := utils.JoinAllProgramNames(Config)
	if len(all_processes) == 0 {
		utils.PrintIfEnoughLevel("Skipped install of NaAVFakeProgramSpawner service: 0 processes provided\n", utils.OPERATION_SKIPPED_MESSAGE)
	} else {
		modules.InstallFakeProgramService()
	}
}

func cloneInstallFiles(ConfigFile string) {
	//TODO use outpututils
	fmt.Printf("Saving files \n")
	cwd, _ := os.Getwd()
	utils.CreateFoldersPath("C:\\Program Files (x86)\\NaAV")
	err := utils.CopyFile(cwd+"\\"+os.Args[1], "C:\\Program Files (x86)\\NaAV\\naav.exe")
	if err != nil {
		fmt.Printf("\t [!] ER-IN004 Unable to save %s , %s \n", "C:\\Program Files (x86)\\NaAV\\naav.exe", err)
	}
	err = utils.CopyFile(ConfigFile, "C:\\Program Files (x86)\\NaAV\\config.json")
	if err != nil {
		fmt.Printf("\t [!] ER-IN004 Unable to save %s , %s \n", "C:\\Program Files (x86)\\NaAV\\config.json", err)
	}
	files := []string{"resources\\NaAVFakeProgramSpawner\\NaAVFakeProgramSpawner.exe", "resources\\dummyprogram\\dummyprogram.exe"}
	for _, file := range files {
		target_file := fmt.Sprintf("C:\\Program Files (x86)\\NaAV\\%s", strings.Split(file, "\\")[len(strings.Split(file, "\\"))-1])
		err = utils.CopyFile(file, target_file)
		if err != nil {
			fmt.Printf("\t [!] ER-IN004 Unable to save %s , %s \n", target_file, err)
		}
	}
}
