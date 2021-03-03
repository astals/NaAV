package modules

// https://underc0de.org/foro/visual-studio-net/(vb-net)-anti-virtual-machine/
func CheckHardware(processes []string, VerbosePlatformName string) {
	/*
		echo %PROCESSOR_ARCHITECTURE% %PROCESSOR_IDENTIFIER% %PROCESSOR_LEVEL% %PROCESSOR_REVISION%
		AMD64 AMD64 Family 25 Model 33 Stepping 0, AuthenticAMD 25 2100
	*/

	/*
		wmic cpu get processorid
		ProcessorId
		178BFBFF00A20F10
	*/

	/*
		wmic cpu get socketdesignation
		SocketDesignation
		AM4
	*/
	/*
		Get-CimInstance -Query 'Select * from Win32_BIOS'
		SMBIOSBIOSVersion : F11q
		Manufacturer      : American Megatrends Inc.
		Name              : F11q
		SerialNumber      : Default string
		Version           : ALASKA - 1072009
	*/

	/*
		Get-CimInstance -Query 'Select * from Win32_ComputerSystem'

		Name             PrimaryOwnerName                                  Domain                                           TotalPhysicalMemory                              Model                                            Manufacturer
		----             ----------------                                  ------                                           -------------------                              -----                                            ------------
		DESKTOP-DG45PJS  cuckoo                                            WORKGROUP                                        17123803136                                      B550M DS3H                                       Gigabyte Technology Co., Ltd.
	*/
	/*
	   Win32_Account
	   Win32_AllocatedResource
	   Win32_BootConfiguration
	   Win32_ClassicCOMApplicationClasses
	   Win32_ClassicCOMClass
	   Win32_ClassicCOMClassSetting
	   Win32_ClassicCOMClassSettings
	   Win32_ClusterShare
	   Win32_BaseService
	   Win32_ClientApplicationSetting
	   Win32_CodecFile
	   Win32_COMApplication
	   Win32_COMApplicationClasses
	   Win32_COMApplicationSettings
	   Win32_COMClass
	   Win32_ComClassAutoEmulator
	   Win32_ComClassEmulator
	   Win32_ComponentCategory
	   Win32_COMSetting
	   Win32_ComputerSystem
	   Win32_ComputerSystem
	   Win32_ComputerSystem
	   Win32_DCOMApplication
	   Win32_DCOMApplicationSetting
	   Win32_DependentService
	   Win32_Desktop
	   Win32_ComputerSystemProduct
	   Win32_Directory
	   Win32_Environment
	   Win32_Group
	   Win32_GroupUser
	   Win32_ImplementedCategory
	   Win32_LoadOrderGroup
	   Win32_LoadOrderGroupServiceDependencies
	   Win32_LoadOrderGroupServiceMembers
	   Win32_LogonSessionMappedDisk
	   Win32_LoggedOnUser
	   Win32_LogonSession
	   Win32_OperatingSystem
	   Win32_OperatingSystemAutochkSetting
	   Win32_OperatingSystemQFE
	   Win32_OptionalFeature
	   Win32_OSRecoveryConfiguration
	   Win32_PageFile
	   Win32_PageFileElementSetting
	   Win32_PageFileSetting
	   Win32_PageFileUsage
	   Win32_PrivilegesStatus
	   Win32_Process
	   Win32_ProcessStartup
	   Win32_ProgramGroupContents
	   Win32_ProgramGroupOrItem
	   Win32_ProtocolBinding
	   Win32_QuickFixEngineering
	   Win32_Registry
	   Win32_ScheduledJob
	   Win32_Service
	   Win32_Session
	   Win32_SessionProcess
	   Win32_SessionResource
	   Win32_Share
	   Win32_ShareToDirectory
	   Win32_StartupCommand
	   Win32_SubDirectory
	   Win32_SubSession
	   Win32_SystemAccount
	   Win32_SystemBIOS
	   Win32_SystemBootConfiguration
	   Win32_SystemConfigurationChangeEvent
	   Win32_SystemDesktop
	   Win32_SystemDevices
	   Win32_SystemDriver
	   Win32_SystemDriverPNPEntity
	   Win32_SystemEnclosure
	   Win32_SystemLoadOrderGroups
	   Win32_SystemMemoryResource
	   Win32_SystemNetworkConnections
	   Win32_SystemOperatingSystem
	   Win32_SystemPartitions
	   Win32_SystemProcesses
	   Win32_SystemProgramGroups
	   Win32_SystemResources
	   Win32_SystemServices
	   Win32_SystemSetting
	   Win32_SystemSlot
	   Win32_SystemSystemDriver
	   Win32_SystemTimeZone
	   Win32_SystemUsers
	   Win32_Thread
	   Win32_TimeZone
	   Win32_UserAccount
	   Win32_UserDesktop
	   Win32_VideoConfiguration
	   Win32_VolumeChangeEvent
	   Win32_ShortcutFile
	*/
}
