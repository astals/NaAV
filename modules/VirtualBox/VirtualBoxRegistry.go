package VirtualBox

import (
	"../../utils"
)

var regKeysAndValueNames = map[string][]string{
	"HKEY_LOCAL_MACHINE\\SYSTEM\\DriverDatabase\\DriverPackages\\vboxguest.inf_amd64_c4ea20f713fb03c4":                                        []string{"Version", "Provider", "InfName", "OemPath", "SignerName", "ImportDate", "Catalog", "SignerScore", "StatusFlags"},
	"HKEY_LOCAL_MACHINE\\SYSTEM\\CurrentControlSet\\Services\\VBoxGuest":                                                                      []string{"Type", "Start", "ErrorControl", "Tag", "ImagePath", "DisplayName", "Group", "Owners"},
	"HKEY_LOCAL_MACHINE\\SYSTEM\\CurrentControlSet\\Services\\VBoxGuest\\Enum":                                                                []string{"0", "Count", "NextInstance"},
	"HKEY_LOCAL_MACHINE\\SYSTEM\\CurrentControlSet\\Services\\VBoxMouse":                                                                      []string{"Type", "Start", "DisplayName", "ErrorControl", "ImagePath"},
	"HKEY_LOCAL_MACHINE\\SYSTEM\\CurrentControlSet\\Services\\VBoxMouse\\DriverInfo":                                                          []string{"RefCount"},
	"HKEY_LOCAL_MACHINE\\SYSTEM\\CurrentControlSet\\Services\\VBoxMouse\\Enum":                                                                []string{"0", "1", "Count", "NextInstance"},
	"HKEY_LOCAL_MACHINE\\SYSTEM\\CurrentControlSet\\Services\\VBoxService":                                                                    []string{"Type", "Start", "ErrorControl", "Tag", "ImagePath", "DisplayName", "Group", "ObjectName", "Description"},
	"HKEY_LOCAL_MACHINE\\SYSTEM\\CurrentControlSet\\Services\\VBoxSF":                                                                         []string{"Type", "Start", "ErrorControl", "Tag", "ImagePath", "DisplayName", "Group"},
	"HKEY_LOCAL_MACHINE\\SYSTEM\\CurrentControlSet\\Services\\VBoxSF\\NetworkProvider":                                                        []string{"DeviceName", "Name", "ProviderPath"},
	"HKEY_LOCAL_MACHINE\\SYSTEM\\CurrentControlSet\\Services\\VBoxWddm":                                                                       []string{"Type", "Start", "ErrorControl", "Tag", "ImagePath", "Group", "Owners"},
	"HKEY_LOCAL_MACHINE\\SYSTEM\\CurrentControlSet\\Services\\VBoxWddm\\Enum":                                                                 []string{"0", "Count", "NextInstance"},
	"HKEY_LOCAL_MACHINE\\SYSTEM\\CurrentControlSet\\Services\\VBoxWddm\\Video":                                                                []string{"Service", "DeviceDesc", "FeatureScore", "VideoID"},
	"HKEY_LOCAL_MACHINE\\SOFTWARE\\Microsoft\\Windows\\CurrentVersion\\Uninstall\\Oracle VM VirtualBox Guest Additions":                       []string{"DisplayName", "DisplayVersion", "Publisher", "UninstallString", "URLInfoAbout"},
	"HKEY_LOCAL_MACHINE\\SOFTWARE\\Microsoft\\Windows\\CurrentVersion\\DIFx\\DriverStore\\vboxguest.inf_amd64_c4ea20f713fb03c4":               []string{"DependentInstaller", "DependentInstallerName"},
	"HKEY_LOCAL_MACHINE\\SOFTWARE\\Microsoft\\Windows\\CurrentVersion\\DIFx\\DriverStore\\VBoxMouse_45447C6D9465DAA979E8D56AA27CB5ADCE72EB96": []string{"creation", "type", "INF", "Services", "ProductName", "ManufacturerName", "DisplayName", "DependentInstaller", "DependentInstallerName"},
	"HKEY_LOCAL_MACHINE\\SOFTWARE\\Microsoft\\Windows\\CurrentVersion\\DIFx\\DriverStore\\vboxwddm.inf_amd64_203cd8905c42de56":                []string{"DependentInstaller", "DependentInstallerName"},
	"HKEY_LOCAL_MACHINE\\DRIVERS\\DriverDatabase\\DriverPackages\\vboxwddm.inf_amd64_203cd8905c42de56":                                        []string{"Version", "Provider", "InfName", "OemPath", "SignerName", "ImportDate", "Catalog", "SignerScore", "StatusFlags"},
}

func InstallVirtualBoxRegistry() {
	utils.InstallRegkeys(regKeysAndValueNames, 2, "\t")
}
