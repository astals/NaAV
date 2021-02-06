package modes

import (
	"../modules"
	"../utils"
)

func Check(ConfigFile string) {
	Config := utils.ReadConfigFile(ConfigFile)
	modules.CheckFiles(Config.VirtualBox.Files, "guest files", "Virtual Box")
	modules.CheckFiles(Config.VMware.Files, "guest files", "VMware")
	modules.CheckFiles(Config.VirtualBox.Drivers, "drivers", "Virtual Box")
	modules.CheckFiles(Config.VMware.Drivers, "drivers", "VMware")
	modules.CheckRegkeys(Config.VirtualBox.Registry.KeysAndValueNames, "Virtual Box")
	modules.CheckTrees(Config.VirtualBox.Registry.OrderedKeysParentTrees)
	modules.CheckRegkeys(Config.VMware.Registry.KeysAndValueNames, "VMware")
	modules.CheckTrees(Config.VMware.Registry.OrderedKeysParentTrees)
	modules.CheckNetworkInterfaces(Config.VirtualBox.NetworkInterfaces, "Virtual Box")
	modules.CheckNetworkInterfaces(Config.VMware.NetworkInterfaces, "VMware")
	modules.CheckNetworkInterfaces(Config.HyperV.NetworkInterfaces, "Hyper-V")
	modules.CheckNetworkInterfaces(Config.Parallels.NetworkInterfaces, "Parallels")
}
