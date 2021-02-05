package modes

import (
	"../modules"
	"../utils"
)

func Uninstall(ConfigFile string) {
	Config := utils.ReadConfigFile(ConfigFile)
	modules.UninstallFiles(Config.VMware.Files, "guest files", "VMware")
	modules.UninstallFiles(Config.VMware.Drivers, "drivers", "VMware")
	modules.UninstallFiles(Config.VirtualBox.Files, "guest files", "Virtual Box")
	modules.UninstallFiles(Config.VirtualBox.Drivers, "drivers", "Virtual Box")
}
