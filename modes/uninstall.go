package modes

import (
	"../modules"
	"../utils"
)

func Uninstall(Config utils.Configuration) {	
	utils.PrintIfEnoughLevel("\n", utils.BASIC_INFORMATION_MESSAGE)
	utils.PrintIfEnoughLevel("=====> Virtual Box <=====\n", utils.BASIC_INFORMATION_MESSAGE)
	modules.UninstallFiles(Config.VirtualBox.Files, "guest files", "Virtual Box")
	modules.UninstallFiles(Config.VirtualBox.Drivers, "drivers", "Virtual Box")
	modules.UninstallRegkeys(Config.VirtualBox.Registry.KeysAndValueNames, "Virtual Box")
	modules.UninstallParentTrees(Config.VirtualBox.Registry.OrderedKeysParentTrees)
	modules.UninstallNetworkInterfaces(Config.VirtualBox.NetworkInterfaces, "Virtual Box")
	
	utils.PrintIfEnoughLevel("\n=====> VMware <=====\n", utils.BASIC_INFORMATION_MESSAGE)
	modules.UninstallFiles(Config.VMware.Files, "guest files", "VMware")
	modules.UninstallFiles(Config.VMware.Drivers, "drivers", "VMware")
	modules.UninstallRegkeys(Config.VMware.Registry.KeysAndValueNames, "VMware")
	modules.UninstallParentTrees(Config.VMware.Registry.OrderedKeysParentTrees)
	modules.UninstallNetworkInterfaces(Config.VMware.NetworkInterfaces, "VMware")
	
	utils.PrintIfEnoughLevel("\n=====> Hyper-V <=====\n", utils.BASIC_INFORMATION_MESSAGE)
	modules.UninstallNetworkInterfaces(Config.HyperV.NetworkInterfaces, "Hyper-V")
	
	utils.PrintIfEnoughLevel("\n=====> Parallels <=====\n", utils.BASIC_INFORMATION_MESSAGE)
	modules.UninstallNetworkInterfaces(Config.Parallels.NetworkInterfaces, "Parallels")

	utils.PrintIfEnoughLevel("=====> Sandboxie <=====\n", utils.BASIC_INFORMATION_MESSAGE)
	modules.UninstallFiles(Config.Sandboxie.Files, "guest files", "Sandboxie")
	
	utils.PrintIfEnoughLevel("\n=====> FakeProgramSpawner service <=====\n", utils.BASIC_INFORMATION_MESSAGE)	
	modules.UninstallFakeProgramService()
	utils.PrintIfEnoughLevel("\n", utils.BASIC_INFORMATION_MESSAGE)
}


