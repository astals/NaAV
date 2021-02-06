package modes

import (
	"../modules"
	"../utils"
)

func Uninstall(ConfigFile string) {
	Config := utils.ReadConfigFile(ConfigFile)
	modules.UninstallFiles(Config.VirtualBox.Files, "guest files", "Virtual Box")
	modules.UninstallFiles(Config.VMware.Files, "guest files", "VMware")
	modules.UninstallFiles(Config.VirtualBox.Drivers, "drivers", "Virtual Box")
	modules.UninstallFiles(Config.VMware.Drivers, "drivers", "VMware")
	modules.UninstallRegkeys(Config.VirtualBox.Registry.KeysAndValueNames, "Virtual Box")
	modules.UninstallParentTrees(Config.VirtualBox.Registry.OrderedKeysParentTrees)
	modules.UninstallRegkeys(Config.VMware.Registry.KeysAndValueNames, "VMware")
	modules.UninstallParentTrees(Config.VMware.Registry.OrderedKeysParentTrees)
	modules.UninstallNetworkInterfaces(Config.VirtualBox.NetworkInterfaces, "Virtual Box")
	modules.UninstallNetworkInterfaces(Config.VMware.NetworkInterfaces, "VMware")
	modules.UninstallNetworkInterfaces(Config.HyperV.NetworkInterfaces, "Hyper-V")
	modules.UninstallNetworkInterfaces(Config.Parallels.NetworkInterfaces, "Parallels")
}
