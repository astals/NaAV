package modes

import (
	"../modules"
	"../utils"
)

func Check(Config utils.Configuration) {
	utils.PrintIfEnoughLevel("\n", utils.BASIC_INFORMATION_MESSAGE)
	utils.PrintIfEnoughLevel("====> Virtual Box <====\n", utils.BASIC_INFORMATION_MESSAGE)
	modules.CheckFiles(Config.VirtualBox.Files, "guest files", "Virtual Box")
	modules.CheckFiles(Config.VirtualBox.Drivers, "drivers", "Virtual Box")
	modules.CheckRegkeys(Config.VirtualBox.Registry.KeysAndValueNames, "Virtual Box")
	modules.CheckTrees(Config.VirtualBox.Registry.OrderedKeysParentTrees)
	modules.CheckNetworkInterfaces(Config.VirtualBox.NetworkInterfaces, "Virtual Box")
	modules.CheckPrograms(Config.VirtualBox.Processes, "Virtual Box")

	utils.PrintIfEnoughLevel("\n====> VMware <====\n", utils.BASIC_INFORMATION_MESSAGE)
	modules.CheckFiles(Config.VMware.Drivers, "drivers", "VMware")
	modules.CheckFiles(Config.VMware.Files, "guest files", "VMware")
	modules.CheckRegkeys(Config.VMware.Registry.KeysAndValueNames, "VMware")
	modules.CheckTrees(Config.VMware.Registry.OrderedKeysParentTrees)
	modules.CheckNetworkInterfaces(Config.VMware.NetworkInterfaces, "VMware")
	modules.CheckPrograms(Config.VMware.Processes, "VMware")

	utils.PrintIfEnoughLevel("\n====> Hyper-V <====\n", utils.BASIC_INFORMATION_MESSAGE)
	modules.CheckNetworkInterfaces(Config.HyperV.NetworkInterfaces, "Hyper-V")

	utils.PrintIfEnoughLevel("\n====> Parallels <====\n", utils.BASIC_INFORMATION_MESSAGE)
	modules.CheckNetworkInterfaces(Config.Parallels.NetworkInterfaces, "Parallels")

	utils.PrintIfEnoughLevel("=====> Sandboxie <=====\n", utils.BASIC_INFORMATION_MESSAGE)
	modules.CheckFiles(Config.Sandboxie.Files, "guest files", "Sandboxie")

	utils.PrintIfEnoughLevel("\n=====> Runing analysis tools <=====\n", utils.BASIC_INFORMATION_MESSAGE)
	modules.CheckPrograms(Config.AnalysisTools, "Analysis Tools")

	utils.PrintIfEnoughLevel("\n====> Hardware & System Checks <====\n", utils.BASIC_INFORMATION_MESSAGE)
	modules.CheckHardware()
	utils.PrintIfEnoughLevel("\n", utils.BASIC_INFORMATION_MESSAGE)
}
