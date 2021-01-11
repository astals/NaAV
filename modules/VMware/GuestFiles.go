package VMware

import (
	"fmt"

	"../../utils"
)

var guestfiles = map[string]string{
	"C:\\Windows\\System32\\vm3dservice.exe":                                                             "resources\\dummy",
	"C:\\Windows\\System32\\vm3dco.dll":                                                                  "resources\\dummy",
	"C:\\Windows\\System32\\vm3ddevapi64-debug.dll":                                                      "resources\\dummy",
	"C:\\Windows\\System32\\vm3ddevapi64-release.dll":                                                    "resources\\dummy",
	"C:\\Windows\\System32\\vm3ddevapi64-stats.dll":                                                      "resources\\dummy",
	"C:\\Windows\\System32\\vm3ddevapi64.dll":                                                            "resources\\dummy",
	"C:\\Windows\\System32\\vm3dgl64.dll":                                                                "resources\\dummy",
	"C:\\Windows\\System32\\vm3dglhelper64.dll":                                                          "resources\\dummy",
	"C:\\Windows\\System32\\vm3dum64-debug.dll":                                                          "resources\\dummy",
	"C:\\Windows\\System32\\vm3dum64-stats.dll":                                                          "resources\\dummy",
	"C:\\Windows\\System32\\vm3dum64.dll":                                                                "resources\\dummy",
	"C:\\Windows\\System32\\vm3dum64_10-debug.dll":                                                       "resources\\dummy",
	"C:\\Windows\\System32\\vm3dum64_10-stats.dll":                                                       "resources\\dummy",
	"C:\\Windows\\System32\\vm3dum64_10.dll":                                                             "resources\\dummy",
	"C:\\Windows\\System32\\vm3dum64_loader.dll":                                                         "resources\\dummy",
	"C:\\Windows\\System32\\vmGuestLib.dll":                                                              "resources\\dummy",
	"C:\\Windows\\System32\\vmGuestLibJava.dll":                                                          "resources\\dummy",
	"C:\\Windows\\System32\\vmhgfs.dll":                                                                  "resources\\dummy",
	"C:\\Windows\\System32\\VMWSU.DLL":                                                                   "resources\\dummy",
	"C:\\Program Files\\VMware\\VMware Tools\\7za.exe":                                                   "resources\\dummy",
	"C:\\Program Files\\VMware\\VMware Tools\\deployPkg.dll":                                             "resources\\dummy",
	"C:\\Program Files\\VMware\\VMware Tools\\gio-2.0.dll":                                               "resources\\dummy",
	"C:\\Program Files\\VMware\\VMware Tools\\glib-2.0.dll":                                              "resources\\dummy",
	"C:\\Program Files\\VMware\\VMware Tools\\glibmm-2.4.dll":                                            "resources\\dummy",
	"C:\\Program Files\\VMware\\VMware Tools\\gmodule-2.0.dll":                                           "resources\\dummy",
	"C:\\Program Files\\VMware\\VMware Tools\\gobject-2.0.dll":                                           "resources\\dummy",
	"C:\\Program Files\\VMware\\VMware Tools\\gthread-2.0.dll":                                           "resources\\dummy",
	"C:\\Program Files\\VMware\\VMware Tools\\guestStoreClient.dll":                                      "resources\\dummy",
	"C:\\Program Files\\VMware\\VMware Tools\\hgfs.dll":                                                  "resources\\dummy",
	"C:\\Program Files\\VMware\\VMware Tools\\iconv.dll":                                                 "resources\\dummy",
	"C:\\Program Files\\VMware\\VMware Tools\\icudt44l.dat":                                              "resources\\dummy",
	"C:\\Program Files\\VMware\\VMware Tools\\install-rvmSetup.cmd":                                      "resources\\dummy",
	"C:\\Program Files\\VMware\\VMware Tools\\intl.dll":                                                  "resources\\dummy",
	"C:\\Program Files\\VMware\\VMware Tools\\libeay32.dll":                                              "resources\\dummy",
	"C:\\Program Files\\VMware\\VMware Tools\\open_source_licenses.txt":                                  "resources\\dummy",
	"C:\\Program Files\\VMware\\VMware Tools\\pcre.dll":                                                  "resources\\dummy",
	"C:\\Program Files\\VMware\\VMware Tools\\poweroff-vm-default.bat":                                   "resources\\dummy",
	"C:\\Program Files\\VMware\\VMware Tools\\poweron-vm-default.bat":                                    "resources\\dummy",
	"C:\\Program Files\\VMware\\VMware Tools\\resume-vm-default.bat":                                     "resources\\dummy",
	"C:\\Program Files\\VMware\\VMware Tools\\rpctool.exe":                                               "resources\\dummy",
	"C:\\Program Files\\VMware\\VMware Tools\\rvmSetup.exe":                                              "resources\\dummy",
	"C:\\Program Files\\VMware\\VMware Tools\\sigc-2.0.dll":                                              "resources\\dummy",
	"C:\\Program Files\\VMware\\VMware Tools\\ssleay32.dll":                                              "resources\\dummy",
	"C:\\Program Files\\VMware\\VMware Tools\\suspend-vm-default.bat":                                    "resources\\dummy",
	"C:\\Program Files\\VMware\\VMware Tools\\vm-support.vbs":                                            "resources\\dummy",
	"C:\\Program Files\\VMware\\VMware Tools\\vmtools.dll":                                               "resources\\dummy",
	"C:\\Program Files\\VMware\\VMware Tools\\vmtoolsd.exe":                                              "resources\\dummy",
	"C:\\Program Files\\VMware\\VMware Tools\\VMToolsHook.dll":                                           "resources\\dummy",
	"C:\\Program Files\\VMware\\VMware Tools\\VMToolsHook64.dll":                                         "resources\\dummy",
	"C:\\Program Files\\VMware\\VMware Tools\\VMToolsHookProc.exe":                                       "resources\\dummy",
	"C:\\Program Files\\VMware\\VMware Tools\\VmUpgradeHelper.bat":                                       "resources\\dummy",
	"C:\\Program Files\\VMware\\VMware Tools\\VMwareHgfsClient.exe":                                      "resources\\dummy",
	"C:\\Program Files\\VMware\\VMware Tools\\VMwareHostOpen.exe":                                        "resources\\dummy",
	"C:\\Program Files\\VMware\\VMware Tools\\VMwareNamespaceCmd.exe":                                    "resources\\dummy",
	"C:\\Program Files\\VMware\\VMware Tools\\VMwareResolutionSet.exe":                                   "resources\\dummy",
	"C:\\Program Files\\VMware\\VMware Tools\\VMwareToolboxCmd.exe":                                      "resources\\dummy",
	"C:\\Program Files\\VMware\\VMware Tools\\VMwareXferlogs.exe":                                        "resources\\dummy",
	"C:\\Program Files\\VMware\\VMware Tools\\messages\\de\\desktopEvents.vmsg":                          "resources\\dummy",
	"C:\\Program Files\\VMware\\VMware Tools\\messages\\de\\hgfsUsability.vmsg":                          "resources\\dummy",
	"C:\\Program Files\\VMware\\VMware Tools\\messages\\de\\toolboxcmd.vmsg":                             "resources\\dummy",
	"C:\\Program Files\\VMware\\VMware Tools\\messages\\de\\vmtoolsd.vmsg":                               "resources\\dummy",
	"C:\\Program Files\\VMware\\VMware Tools\\messages\\es\\desktopEvents.vmsg":                          "resources\\dummy",
	"C:\\Program Files\\VMware\\VMware Tools\\messages\\es\\hgfsUsability.vmsg":                          "resources\\dummy",
	"C:\\Program Files\\VMware\\VMware Tools\\messages\\es\\toolboxcmd.vmsg":                             "resources\\dummy",
	"C:\\Program Files\\VMware\\VMware Tools\\messages\\es\\vmtoolsd.vmsg":                               "resources\\dummy",
	"C:\\Program Files\\VMware\\VMware Tools\\messages\\fr\\desktopEvents.vmsg":                          "resources\\dummy",
	"C:\\Program Files\\VMware\\VMware Tools\\messages\\fr\\hgfsUsability.vmsg":                          "resources\\dummy",
	"C:\\Program Files\\VMware\\VMware Tools\\messages\\fr\\toolboxcmd.vmsg":                             "resources\\dummy",
	"C:\\Program Files\\VMware\\VMware Tools\\messages\\fr\\vmtoolsd.vmsg":                               "resources\\dummy",
	"C:\\Program Files\\VMware\\VMware Tools\\messages\\it\\desktopEvents.vmsg":                          "resources\\dummy",
	"C:\\Program Files\\VMware\\VMware Tools\\messages\\it\\hgfsUsability.vmsg":                          "resources\\dummy",
	"C:\\Program Files\\VMware\\VMware Tools\\messages\\it\\toolboxcmd.vmsg":                             "resources\\dummy",
	"C:\\Program Files\\VMware\\VMware Tools\\messages\\it\\vmtoolsd.vmsg":                               "resources\\dummy",
	"C:\\Program Files\\VMware\\VMware Tools\\messages\\ja\\desktopEvents.vmsg":                          "resources\\dummy",
	"C:\\Program Files\\VMware\\VMware Tools\\messages\\ja\\hgfsUsability.vmsg":                          "resources\\dummy",
	"C:\\Program Files\\VMware\\VMware Tools\\messages\\ja\\toolboxcmd.vmsg":                             "resources\\dummy",
	"C:\\Program Files\\VMware\\VMware Tools\\messages\\ja\\vmtoolsd.vmsg":                               "resources\\dummy",
	"C:\\Program Files\\VMware\\VMware Tools\\messages\\ko\\desktopEvents.vmsg":                          "resources\\dummy",
	"C:\\Program Files\\VMware\\VMware Tools\\messages\\ko\\hgfsUsability.vmsg":                          "resources\\dummy",
	"C:\\Program Files\\VMware\\VMware Tools\\messages\\ko\\toolboxcmd.vmsg":                             "resources\\dummy",
	"C:\\Program Files\\VMware\\VMware Tools\\messages\\ko\\vmtoolsd.vmsg":                               "resources\\dummy",
	"C:\\Program Files\\VMware\\VMware Tools\\messages\\zh_CN\\desktopEvents.vmsg":                       "resources\\dummy",
	"C:\\Program Files\\VMware\\VMware Tools\\messages\\zh_CN\\hgfsUsability.vmsg":                       "resources\\dummy",
	"C:\\Program Files\\VMware\\VMware Tools\\messages\\zh_CN\\toolboxcmd.vmsg":                          "resources\\dummy",
	"C:\\Program Files\\VMware\\VMware Tools\\messages\\zh_CN\\vmtoolsd.vmsg":                            "resources\\dummy",
	"C:\\Program Files\\VMware\\VMware Tools\\messages\\zh_TW\\desktopEvents.vmsg":                       "resources\\dummy",
	"C:\\Program Files\\VMware\\VMware Tools\\messages\\zh_TW\\hgfsUsability.vmsg":                       "resources\\dummy",
	"C:\\Program Files\\VMware\\VMware Tools\\messages\\zh_TW\\toolboxcmd.vmsg":                          "resources\\dummy",
	"C:\\Program Files\\VMware\\VMware Tools\\messages\\zh_TW\\vmtoolsd.vmsg":                            "resources\\dummy",
	"C:\\Program Files\\VMware\\VMware Tools\\plugins\\common\\hgfsServer.dll":                           "resources\\dummy",
	"C:\\Program Files\\VMware\\VMware Tools\\plugins\\common\\hgfsUsability.dll":                        "resources\\dummy",
	"C:\\Program Files\\VMware\\VMware Tools\\plugins\\common\\vix.dll":                                  "resources\\dummy",
	"C:\\Program Files\\VMware\\VMware Tools\\plugins\\vmsvc\\appInfo.dll":                               "resources\\dummy",
	"C:\\Program Files\\VMware\\VMware Tools\\plugins\\vmsvc\\autoLogon.dll":                             "resources\\dummy",
	"C:\\Program Files\\VMware\\VMware Tools\\plugins\\vmsvc\\autoUpgrade.dll":                           "resources\\dummy",
	"C:\\Program Files\\VMware\\VMware Tools\\plugins\\vmsvc\\bitMapper.dll":                             "resources\\dummy",
	"C:\\Program Files\\VMware\\VMware Tools\\plugins\\vmsvc\\deployPkgPlugin.dll":                       "resources\\dummy",
	"C:\\Program Files\\VMware\\VMware Tools\\plugins\\vmsvc\\disableGuestHibernate.dll":                 "resources\\dummy",
	"C:\\Program Files\\VMware\\VMware Tools\\plugins\\vmsvc\\diskWiper.dll":                             "resources\\dummy",
	"C:\\Program Files\\VMware\\VMware Tools\\plugins\\vmsvc\\guestInfo.dll":                             "resources\\dummy",
	"C:\\Program Files\\VMware\\VMware Tools\\plugins\\vmsvc\\guestStore.dll":                            "resources\\dummy",
	"C:\\Program Files\\VMware\\VMware Tools\\plugins\\vmsvc\\guestStoreUpgrade.dll":                     "resources\\dummy",
	"C:\\Program Files\\VMware\\VMware Tools\\plugins\\vmsvc\\hwUpgradeHelper.dll":                       "resources\\dummy",
	"C:\\Program Files\\VMware\\VMware Tools\\plugins\\vmsvc\\powerOps.dll":                              "resources\\dummy",
	"C:\\Program Files\\VMware\\VMware Tools\\plugins\\vmsvc\\resolutionSet.dll":                         "resources\\dummy",
	"C:\\Program Files\\VMware\\VMware Tools\\plugins\\vmsvc\\timeSync.dll":                              "resources\\dummy",
	"C:\\Program Files\\VMware\\VMware Tools\\plugins\\vmsvc\\vmbackup.dll":                              "resources\\dummy",
	"C:\\Program Files\\VMware\\VMware Tools\\plugins\\vmusr\\darkModeSync.dll":                          "resources\\dummy",
	"C:\\Program Files\\VMware\\VMware Tools\\plugins\\vmusr\\desktopEvents.dll":                         "resources\\dummy",
	"C:\\Program Files\\VMware\\VMware Tools\\plugins\\vmusr\\dndcp.dll":                                 "resources\\dummy",
	"C:\\Program Files\\VMware\\VMware Tools\\plugins\\vmusr\\unity.dll":                                 "resources\\dummy",
	"C:\\Program Files\\VMware\\VMware Tools\\plugins\\vmusr\\vmtray.dll":                                "resources\\dummy",
	"C:\\Program Files\\VMware\\VMware Tools\\VMware VGAuth\\gio-2.0.dll":                                "resources\\dummy",
	"C:\\Program Files\\VMware\\VMware Tools\\VMware VGAuth\\glib-2.0.dll":                               "resources\\dummy",
	"C:\\Program Files\\VMware\\VMware Tools\\VMware VGAuth\\gmodule-2.0.dll":                            "resources\\dummy",
	"C:\\Program Files\\VMware\\VMware Tools\\VMware VGAuth\\gobject-2.0.dll":                            "resources\\dummy",
	"C:\\Program Files\\VMware\\VMware Tools\\VMware VGAuth\\gthread-2.0.dll":                            "resources\\dummy",
	"C:\\Program Files\\VMware\\VMware Tools\\VMware VGAuth\\iconv.dll":                                  "resources\\dummy",
	"C:\\Program Files\\VMware\\VMware Tools\\VMware VGAuth\\intl.dll":                                   "resources\\dummy",
	"C:\\Program Files\\VMware\\VMware Tools\\VMware VGAuth\\libeay32.dll":                               "resources\\dummy",
	"C:\\Program Files\\VMware\\VMware Tools\\VMware VGAuth\\libxml2.dll":                                "resources\\dummy",
	"C:\\Program Files\\VMware\\VMware Tools\\VMware VGAuth\\libxmlsec-openssl.dll":                      "resources\\dummy",
	"C:\\Program Files\\VMware\\VMware Tools\\VMware VGAuth\\libxmlsec.dll":                              "resources\\dummy",
	"C:\\Program Files\\VMware\\VMware Tools\\VMware VGAuth\\pcre.dll":                                   "resources\\dummy",
	"C:\\Program Files\\VMware\\VMware Tools\\VMware VGAuth\\ssleay32.dll":                               "resources\\dummy",
	"C:\\Program Files\\VMware\\VMware Tools\\VMware VGAuth\\VGAuth.dll":                                 "resources\\dummy",
	"C:\\Program Files\\VMware\\VMware Tools\\VMware VGAuth\\VGAuthCLI.exe":                              "resources\\dummy",
	"C:\\Program Files\\VMware\\VMware Tools\\VMware VGAuth\\VGAuthService.exe":                          "resources\\dummy",
	"C:\\Program Files\\VMware\\VMware Tools\\VMware VGAuth\\vmtools.dll":                                "resources\\dummy",
	"C:\\Program Files\\VMware\\VMware Tools\\VMware VGAuth\\VMwareAliasImport.exe":                      "resources\\dummy",
	"C:\\Program Files\\VMware\\VMware Tools\\VMware VGAuth\\VMWSU.dll":                                  "resources\\dummy",
	"C:\\Program Files\\VMware\\VMware Tools\\VMware VGAuth\\schemas\\catalog.xml":                       "resources\\dummy",
	"C:\\Program Files\\VMware\\VMware Tools\\VMware VGAuth\\schemas\\datatypes.dtd":                     "resources\\dummy",
	"C:\\Program Files\\VMware\\VMware Tools\\VMware VGAuth\\schemas\\saml-schema-assertion-2.0.xsd":     "resources\\dummy",
	"C:\\Program Files\\VMware\\VMware Tools\\VMware VGAuth\\schemas\\xenc-schema.xsd":                   "resources\\dummy",
	"C:\\Program Files\\VMware\\VMware Tools\\VMware VGAuth\\schemas\\xml.xsd":                           "resources\\dummy",
	"C:\\Program Files\\VMware\\VMware Tools\\VMware VGAuth\\schemas\\xmldsig-core-schema.xsd":           "resources\\dummy",
	"C:\\Program Files\\VMware\\VMware Tools\\VMware VGAuth\\schemas\\XMLSchema-hasFacetAndProperty.xsd": "resources\\dummy",
	"C:\\Program Files\\VMware\\VMware Tools\\VMware VGAuth\\schemas\\XMLSchema-instance.xsd":            "resources\\dummy",
	"C:\\Program Files\\VMware\\VMware Tools\\VMware VGAuth\\schemas\\XMLSchema.dtd":                     "resources\\dummy",
	"C:\\Program Files\\VMware\\VMware Tools\\VMware VGAuth\\schemas\\XMLSchema.xsd":                     "resources\\dummy",
	"C:\\Program Files\\VMware\\VMware Tools\\win32\\vmGuestLib.dll":                                     "resources\\dummy",
	"C:\\Program Files\\VMware\\VMware Tools\\win32\\vmGuestLibJava.dll":                                 "resources\\dummy",
	"C:\\Program Files\\VMware\\VMware Tools\\win64\\vmGuestLib.dll":                                     "resources\\dummy",
	"C:\\Program Files\\VMware\\VMware Tools\\win64\\vmGuestLibJava.dll":                                 "resources\\dummy",
}

func InstallVMwareGuestFiles() {
	fmt.Printf("Copying VMware Guest files on system:\n")
	okOperations := 0
	for destination, origin := range guestfiles {
		err := utils.SafeCopy(origin, destination, "\t")
		if err == nil {
			okOperations++
		}
	}
	fmt.Printf("\t [i] Performed %d of %d operations\n", okOperations, len(guestfiles))
}

func UninstallVMwareGuestFiles() {
	fmt.Printf("Removing VMware Guest files on system:\n")
	okOperations := 0
	for destination, _ := range guestfiles {
		res, _ := utils.FileExists(destination)
		if res {
			success, err := utils.DeleteIfIsNaAVFile(destination, "\t")
			if success && err == nil {
				okOperations++
			}
		} else {
			okOperations++
		}
	}
	fmt.Printf("\t [i] Performed %d of %d operations\n", okOperations, len(guestfiles))
}

func CheckVMwareGuestFiles() {
	fmt.Printf("Checking  VMware Guest files on system:\n")
	okOperations := 0
	for destination, _ := range guestfiles {
		res, _ := utils.FileExists(destination)
		if res {
			okOperations++
		}
	}
	fmt.Printf("\t [i] Found %d of %d files\n", okOperations, len(guestfiles))
}
