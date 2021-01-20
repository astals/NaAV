package VMware

import (
	"../../utils"
)

var regKeysAndValueNames = map[string][]string{
	"HKEY_LOCAL_MACHINE\\SOFTWARE\\Classes\\Applications\\VMwareHostOpen.exe\\shell\\open\\command":                                                                                                     []string{},
	"HKEY_LOCAL_MACHINE\\SOFTWARE\\Classes\\VMwareHostOpen.AssocFile\\shell\\open\\command":                                                                                                             []string{},
	"HKEY_LOCAL_MACHINE\\SOFTWARE\\Classes\\VMwareHostOpen.AssocURL\\shell\\open\\command":                                                                                                              []string{},
	"HKEY_LOCAL_MACHINE\\SOFTWARE\\Clients\\StartMenuInternet\\VMWAREHOSTOPEN.EXE":                                                                                                                      []string{"LocalizedString"},
	"HKEY_LOCAL_MACHINE\\SOFTWARE\\Clients\\StartMenuInternet\\VMWAREHOSTOPEN.EXE\\DefaultIcon":                                                                                                         []string{},
	"HKEY_LOCAL_MACHINE\\SOFTWARE\\Clients\\StartMenuInternet\\VMWAREHOSTOPEN.EXE\\InstallInfo":                                                                                                         []string{"InstallInfo", "InstallInfo", "ReinstallCommand", "ShowIconsCommand"},
	"HKEY_LOCAL_MACHINE\\SOFTWARE\\Clients\\StartMenuInternet\\VMWAREHOSTOPEN.EXE\\shell\\open\\command":                                                                                                []string{},
	"HKEY_LOCAL_MACHINE\\SOFTWARE\\Microsoft\\Windows\\CurrentVersion\\Setup\\PnpResources\\Registry\\HKLM\\SOFTWARE\\Microsoft\\Windows\\CurrentVersion\\Run\\VMware VM3DService Process":              []string{"Owners"},
	"HKEY_LOCAL_MACHINE\\SOFTWARE\\VMware, Inc.\\VMware Drivers":                                                                                                                                        []string{"efifw.status", "pvscsi.status", "svga_wddm.status", "vmci.status", "VmciHostDevInst", "vmhgfs.status", "vmhgfs.status", "vmmouse.status", "vmrawdsk.status", "vmusbmouse.status", "vmxnet3.status", "vsock.status", "vsockDll.status", "vsockSys.status"},
	"HKEY_LOCAL_MACHINE\\SOFTWARE\\Clients\\WOW6432Node\\StartMenuInternet\\VMWAREHOSTOPEN.EXE\\DefaultIcon":                                                                                            []string{},
	"HKEY_LOCAL_MACHINE\\SOFTWARE\\Clients\\WOW6432Node\\StartMenuInternet\\VMWAREHOSTOPEN.EXE\\InstallInfo":                                                                                            []string{"InstallInfo", "InstallInfo", "ReinstallCommand", "ShowIconsCommand"},
	"HKEY_LOCAL_MACHINE\\SOFTWARE\\Clients\\WOW6432Node\\StartMenuInternet\\VMWAREHOSTOPEN.EXE\\shell\\open\\command":                                                                                   []string{},
	"HKEY_LOCAL_MACHINE\\SOFTWARE\\WOW6432Node\\Microsoft\\Windows\\CurrentVersion\\Setup\\PnpResources\\Registry\\HKLM\\SOFTWARE\\Microsoft\\Windows\\CurrentVersion\\Run\\VMware VM3DService Process": []string{"Owners"},
	"HKEY_USERS\\.DEFAULT\\Software\\VMware, Inc.\\VMware Tools\\Hgfs Usability\\MRU RootShare":                                                                                                         []string{},
	"HKEY_USERS\\S-1-5-18\\Software\\VMware, Inc.\\VMware Tools\\Hgfs Usability\\MRU RootShare":                                                                                                         []string{},
	"HKEY_USERS\\S-1-5-19\\Software\\VMware, Inc.\\VMware Tools\\Hgfs Usability\\MRU RootShare":                                                                                                         []string{},
	"HKEY_USERS\\S-1-5-20\\Software\\VMware, Inc.\\VMware Tools\\Hgfs Usability\\MRU RootShare":                                                                                                         []string{},
}

/*
HKEY_LOCAL_MACHINE\SYSTEM\ControlSet001\Enum\ACPI\VMW0001
HKEY_LOCAL_MACHINE\SYSTEM\ControlSet001\Enum\ACPI\VMW0003
HKEY_LOCAL_MACHINE\SYSTEM\ControlSet001\Enum\ROOT\VMWVMCIHOSTDEV
HKEY_LOCAL_MACHINE\SYSTEM\ControlSet001\Enum\SCSI\CdRom&Ven_NECVMWar&Prod_VMware_SATA_CD01
HKEY_LOCAL_MACHINE\SYSTEM\ControlSet001\Enum\SCSI\Disk&Ven_NVMe&Prod_VMware_Virtual_N
HKEY_LOCAL_MACHINE\SYSTEM\ControlSet001\Services\EventLog\Application\VMware Tools
HKEY_LOCAL_MACHINE\SYSTEM\ControlSet001\Services\vmwefifw
HKEY_LOCAL_MACHINE\SYSTEM\CurrentControlSet\Enum\ACPI\VMW0001
HKEY_LOCAL_MACHINE\SYSTEM\CurrentControlSet\Enum\ACPI\VMW0003
HKEY_LOCAL_MACHINE\SYSTEM\CurrentControlSet\Enum\ROOT\VMWVMCIHOSTDEV
HKEY_LOCAL_MACHINE\SYSTEM\CurrentControlSet\Enum\SCSI\CdRom&Ven_NECVMWar&Prod_VMware_SATA_CD01
HKEY_LOCAL_MACHINE\SYSTEM\CurrentControlSet\Enum\SCSI\Disk&Ven_NVMe&Prod_VMware_Virtual_N
HKEY_LOCAL_MACHINE\SYSTEM\CurrentControlSet\Services\EventLog\Application\VMware Tools
HKEY_LOCAL_MACHINE\SYSTEM\CurrentControlSet\Services\vmwefifw
HKEY_LOCAL_MACHINE\SYSTEM\DriverDatabase\DeviceIds\*VMW0003
HKEY_LOCAL_MACHINE\SYSTEM\DriverDatabase\DeviceIds\Root\VMWVMCIHOSTDEV
HKEY_LOCAL_MACHINE\SYSTEM\DriverDatabase\DriverPackages\vmci.inf_amd64_5e38a278d114b813\Descriptors\ROOT\VMWVMCIHOSTDEV
HKEY_LOCAL_MACHINE\SYSTEM\DriverDatabase\DriverPackages\vmmouse.inf_amd64_916101d3748847e7\Descriptors\*VMW0003
Computer\HKEY_USERS\S-1-5-21-6975932-3295827324-3868878605-1000(user)
*/
func InstallVirtualBoxRegistry() {
	utils.InstallRegkeys(regKeysAndValueNames, 2, "\\t")
}
