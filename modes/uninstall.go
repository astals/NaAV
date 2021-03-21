package modes

import (
	"../modules"
	"../utils"
)

func Uninstall(ConfigFile string) {
	Config := utils.ReadConfigFile(ConfigFile)
	utils.PrintIfEnoughLevel("=====> Virtual Box <=====\n", utils.BASIC_INFORMATION_MESSAGE)
	modules.UninstallFiles(Config.VirtualBox.Files, "guest files", "Virtual Box")
	modules.UninstallFiles(Config.VirtualBox.Drivers, "drivers", "Virtual Box")
	modules.UninstallRegkeys(Config.VirtualBox.Registry.KeysAndValueNames, "Virtual Box")
	modules.UninstallParentTrees(Config.VirtualBox.Registry.OrderedKeysParentTrees)
	modules.UninstallNetworkInterfaces(Config.VirtualBox.NetworkInterfaces, "Virtual Box")
	
	utils.PrintIfEnoughLevel("=====> VMware <=====\n", utils.BASIC_INFORMATION_MESSAGE)
	modules.UninstallFiles(Config.VMware.Files, "guest files", "VMware")
	modules.UninstallFiles(Config.VMware.Drivers, "drivers", "VMware")
	modules.UninstallRegkeys(Config.VMware.Registry.KeysAndValueNames, "VMware")
	modules.UninstallParentTrees(Config.VMware.Registry.OrderedKeysParentTrees)
	modules.UninstallNetworkInterfaces(Config.VMware.NetworkInterfaces, "VMware")
	
	utils.PrintIfEnoughLevel("=====> Hyper-V <=====\n", utils.BASIC_INFORMATION_MESSAGE)
	modules.UninstallNetworkInterfaces(Config.HyperV.NetworkInterfaces, "Hyper-V")
	
	utils.PrintIfEnoughLevel("=====> Parallels <=====\n", utils.BASIC_INFORMATION_MESSAGE)
	modules.UninstallNetworkInterfaces(Config.Parallels.NetworkInterfaces, "Parallels")
	
	utils.PrintIfEnoughLevel("=====> FakeProgramSpawner service <=====\n", utils.BASIC_INFORMATION_MESSAGE)	
	modules.UninstallFakeProgramService()
}


