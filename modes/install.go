package modes

import (
	"fmt"
	"os"
	"io/ioutil"
	"../modules"
	"../utils"
	"embed"
)



func Install(Config utils.Configuration) {
	utils.PrintIfEnoughLevel("\n", utils.BASIC_INFORMATION_MESSAGE)
	utils.PrintIfEnoughLevel("=====> Virtual Box <=====\n", utils.BASIC_INFORMATION_MESSAGE)
	modules.InstallFiles(Config.VirtualBox.Files, "guest files", "Virtual Box")
	modules.InstallFiles(Config.VirtualBox.Drivers, "drivers", "Virtual Box")
	modules.InstallRegkeys(Config.VirtualBox.Registry.KeysAndValueNames, "Virtual Box")
	modules.InstallNetworkInterfaces(Config.VirtualBox.NetworkInterfaces, "Virtual Box")
	
	utils.PrintIfEnoughLevel("\n=====> VMware <=====\n", utils.BASIC_INFORMATION_MESSAGE)
	modules.InstallFiles(Config.VMware.Files, "guest files", "VMware")
	modules.InstallFiles(Config.VMware.Drivers, "drivers", "VMware")
	modules.InstallRegkeys(Config.VMware.Registry.KeysAndValueNames, "VMware")
	modules.InstallNetworkInterfaces(Config.VMware.NetworkInterfaces, "VMware")
	
	utils.PrintIfEnoughLevel("\n=====> Hyper-V <=====\n", utils.BASIC_INFORMATION_MESSAGE)
	modules.InstallNetworkInterfaces(Config.HyperV.NetworkInterfaces, "Hyper-V")
	
	utils.PrintIfEnoughLevel("\n=====> Parallels <=====\n", utils.BASIC_INFORMATION_MESSAGE)
	modules.InstallNetworkInterfaces(Config.Parallels.NetworkInterfaces, "Parallels")

	utils.PrintIfEnoughLevel("=====> Sandboxie <=====\n", utils.BASIC_INFORMATION_MESSAGE)
	modules.InstallFiles(Config.Sandboxie.Files, "guest files", "Sandboxie")
	
	utils.PrintIfEnoughLevel("\n=====> FakeProgramSpawner service <=====\n", utils.BASIC_INFORMATION_MESSAGE)
	all_processes := utils.JoinAllProgramNames(Config)
	if len(all_processes) == 0 {
		utils.PrintIfEnoughLevel("Skipped install of NaAVFakeProgramSpawner service: 0 processes provided\n", utils.OPERATION_SKIPPED_MESSAGE)
	} else {
		modules.InstallFakeProgramService()
	}
	utils.PrintIfEnoughLevel("\n", utils.BASIC_INFORMATION_MESSAGE)
}

func CloneInstallFiles(configfile string, resources embed.FS) {
	utils.PrintIfEnoughLevel("Saving files \n", utils.BASIC_INFORMATION_MESSAGE)
	utils.CreateFoldersPath("C:\\Program Files (x86)\\NaAV")
	if configfile !=""{
		cwd, _ := os.Getwd()
		err := utils.CopyFile(cwd+"\\"+os.Args[1], "C:\\Program Files (x86)\\NaAV\\naav.exe")
		if err != nil {
			fmt.Printf("\t [!] ER-IN004 Unable to save %s , %s \n", "C:\\Program Files (x86)\\NaAV\\naav.exe", err)
			os.Exit(5)
		}
	}else{
		bytesRead, err := resources.ReadFile("resources/defaultconfig.json")
		if err != nil {
			utils.PrintIfEnoughLevel(fmt.Sprintf("[!] ER-CO003 Unable to read config file, %s \n", err), utils.OPERATION_ERROR_MESSAGE)
			os.Exit(5)
		}
		err = ioutil.WriteFile("C:\\Program Files (x86)\\NaAV\\config.json",bytesRead,0600)
		if err != nil {
			fmt.Printf("\t [!] ER-IN004 Unable to save %s , %s \n", "C:\\Program Files (x86)\\NaAV\\config.json", err)
		}
	}
	files := map[string]string{
		"resources/NaAVFakeProgramSpawner/NaAVFakeProgramSpawner.exe":"C:\\Program Files (x86)\\NaAV\\NaAVFakeProgramSpawner.exe",
	 	"resources/dummyprogram/dummyprogram.exe":"C:\\Program Files (x86)\\NaAV\\dummyprogram.exe",
	}
	for source, destination := range files {
		bytesRead, err := resources.ReadFile(source)
		if err != nil {
			utils.PrintIfEnoughLevel(fmt.Sprintf("[!] ER-CO003 Unable to read file %s, %s \n", source, err), utils.OPERATION_ERROR_MESSAGE)
			os.Exit(5)
		}
		err = ioutil.WriteFile(destination,bytesRead,0600)
		if err != nil {
			fmt.Printf("\t [!] ER-IN004 Unable to save %s , %s \n", destination, err)
		}
	}
}
