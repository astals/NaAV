package modes

import (
	"../modules"
	"../utils"
)

func Check(ConfigFile string) {
	Config := utils.ReadConfigFile(ConfigFile)
	modules.CheckFiles(Config.VMware.Files, "guest files", "VMware")
	modules.CheckFiles(Config.VMware.Drivers, "drivers", "VMware")
	modules.CheckFiles(Config.VirtualBox.Files, "guest files", "Virtual Box")
	modules.CheckFiles(Config.VirtualBox.Drivers, "drivers", "Virtual Box")
}
