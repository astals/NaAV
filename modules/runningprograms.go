package modules

import (
	"fmt"
	"syscall"
	"unsafe"

	"../utils"
	"golang.org/x/sys/windows"
	"golang.org/x/sys/windows/svc"
	"golang.org/x/sys/windows/svc/mgr"
)

type WindowsProcess struct {
	ProcessID       int
	ParentProcessID int
	Exe             string
}

var servicename = "NaAVFakeProgramSpawner"

func InstallFakeProgramService() {
	//TODO evaluate startup folder
	utils.PrintIfEnoughLevel("Installing NaAVFakeProgramSpawner service\n", utils.BASIC_INFORMATION_MESSAGE)
	servicemanager, err := mgr.Connect()
	defer servicemanager.Disconnect()
	if err != nil {
		utils.PrintIfEnoughLevel(fmt.Sprintf("\t [!] ER-FP001 Error conecting with the service manager: %s\n", err), utils.OPERATION_ERROR_MESSAGE)
		return
	}
	service, err := servicemanager.CreateService(servicename, "C:\\Program Files (x86)\\NaAV\\NaAVFakeProgramSpawner.exe", mgr.Config{DisplayName: servicename, StartType: mgr.StartAutomatic})
	if err != nil {
		utils.PrintIfEnoughLevel(fmt.Sprintf("\t [!] ER-FP002 Error creating service %s: %s\n", servicename, err), utils.OPERATION_ERROR_MESSAGE)
		return
	}
	utils.PrintIfEnoughLevel(fmt.Sprintf("%s [i] Successfully created %s service\n", "\t", servicename), utils.SUMMARY_MESSAGE)
	err = service.Start()
	if err != nil {
		utils.PrintIfEnoughLevel(fmt.Sprintf("\t [!] ER-FP002 Error starting service %s: %s\n", servicename, err), utils.OPERATION_ERROR_MESSAGE)
		return
	}
	utils.PrintIfEnoughLevel(fmt.Sprintf("%s [i] Successfully started %s service\n", "\t", servicename), utils.SUMMARY_MESSAGE)
}

func UninstallFakeProgramService() {
	config := utils.LoadConfigFromFile("C:\\Program Files (x86)\\NaAV\\config.json")
	all_processes := utils.JoinAllProgramNames(config)
	if len(all_processes) == 0 {
		utils.PrintIfEnoughLevel("Skipped uninstall of NaAVFakeProgramSpawner service: 0 processes found\n", utils.OPERATION_SKIPPED_MESSAGE)
		return
	}
	utils.PrintIfEnoughLevel("Removing NaAVFakeProgramSpawner service\n", utils.BASIC_INFORMATION_MESSAGE)
	servicemanager, err := mgr.Connect()
	defer servicemanager.Disconnect()
	if err != nil {
		utils.PrintIfEnoughLevel(fmt.Sprintf("\t [!] ER-FP003 Error conecting with the service manager: %s\n", err), utils.OPERATION_ERROR_MESSAGE)
		return
	}
	service, err := servicemanager.OpenService(servicename)
	if err != nil {
		utils.PrintIfEnoughLevel(fmt.Sprintf("\t [!] ER-FP004 Error conecting with the service %s: %s\n", servicename, err), utils.OPERATION_ERROR_MESSAGE)
		return
	}
	_, err = service.Control(svc.Stop)
	if err != nil {
		utils.PrintIfEnoughLevel(fmt.Sprintf("\t [!] ER-FP005 Error stopping the service %s: %s\n", servicename, err), utils.OPERATION_ERROR_MESSAGE)
	}
	err = service.Delete()
	if err != nil {
		utils.PrintIfEnoughLevel(fmt.Sprintf("\t [!] ER-FP006 Error deleting the service %s: %s\n", servicename, err), utils.OPERATION_ERROR_MESSAGE)
		return
	}
	utils.PrintIfEnoughLevel(fmt.Sprintf("%s [i] Successfully removed %s service\n", "\t", servicename), utils.SUMMARY_MESSAGE)
}

func CheckPrograms(processes []string, VerbosePlatformName string) {
	utils.PrintIfEnoughLevel(fmt.Sprintf("Checking %s processes:\n", VerbosePlatformName), utils.BASIC_INFORMATION_MESSAGE)
	var runningProcessesNames []string
	runningProcesses, err := GetRunningProcesses()
	if err != nil {
		utils.PrintIfEnoughLevel(fmt.Sprintf("\t [!] ER-FP006 Error getting running processes: %s\n", err), utils.OPERATION_ERROR_MESSAGE)
		return
	}
	for _, p := range runningProcesses {
		runningProcessesNames = append(runningProcessesNames, GetProcessName(p))
	}
	detected := 0
	for _, p := range processes {
		if utils.ElementInStringArray(p, runningProcessesNames) {
			detected++
		}
	}
	utils.PrintIfEnoughLevel(fmt.Sprintf("%s [i] Detected %d of %d processes \n", "\t", detected, len(processes)), utils.SUMMARY_MESSAGE)
}

func GetRunningProcesses() ([]windows.ProcessEntry32, error) {
	// https://stackoverflow.com/a/46849516
	var res []windows.ProcessEntry32
	handle, err := windows.CreateToolhelp32Snapshot(0x00000002, 0)
	if err != nil {
		return nil, err
	}
	defer windows.CloseHandle(handle)
	var entry windows.ProcessEntry32
	entry.Size = uint32(unsafe.Sizeof(entry))
	err = windows.Process32First(handle, &entry)
	res = append(res, entry)
	if err != nil {
		return nil, err
	}
	for {
		err = windows.Process32Next(handle, &entry)
		if err != nil {
			if err == syscall.ERROR_NO_MORE_FILES {
				return res, nil
			}
			return nil, err
		}
		res = append(res, entry)
	}
}

func GetProcessName(process windows.ProcessEntry32) string {
	end := 0
	for {
		if process.ExeFile[end] == 0 {
			break
		}
		end++
	}
	return syscall.UTF16ToString(process.ExeFile[:end])
}
