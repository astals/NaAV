package modes

import (
	"../modules"
	"../utils"
)

func Uninstall(ConfigFile string) {
	Config := utils.ReadConfigFile(ConfigFile)
	modules.UninstallFiles(Config.VMware.Files, "guest files", "VMware")
	modules.UninstallFiles(Config.VirtualBox.Files, "guest files", "Virtual Box")
}
