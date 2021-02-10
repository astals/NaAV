package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"time"

	"../../utils"

	"golang.org/x/sys/windows/svc"
	"golang.org/x/sys/windows/svc/debug"
	"golang.org/x/sys/windows/svc/eventlog"
)

var elog debug.Log
var ChildPids []int

const svcName = "NaAVFakeProgramSpawner"

type myservice struct{}

func (m *myservice) Execute(args []string, r <-chan svc.ChangeRequest, changes chan<- svc.Status) (ssec bool, errno uint32) {
	const cmdsAccepted = svc.AcceptStop | svc.AcceptShutdown | svc.AcceptPauseAndContinue
	changes <- svc.Status{State: svc.StartPending}
	fasttick := time.Tick(500 * time.Millisecond)
	slowtick := time.Tick(2 * time.Second)
	tick := fasttick
	changes <- svc.Status{State: svc.Running, Accepts: cmdsAccepted}
	SpawnChilds()
loop:
	for {
		select {
		case <-tick:

		case c := <-r:
			switch c.Cmd {
			case svc.Interrogate:
				changes <- c.CurrentStatus
				// Testing deadlock from https://code.google.com/p/winsvc/issues/detail?id=4
				time.Sleep(100 * time.Millisecond)
				changes <- c.CurrentStatus
			case svc.Stop, svc.Shutdown:
				// golang.org/x/sys/windows/svc.TestExample is verifying this output.
				KillChilds()
				break loop
			case svc.Pause:
				changes <- svc.Status{State: svc.Paused, Accepts: cmdsAccepted}
				tick = slowtick
			case svc.Continue:
				changes <- svc.Status{State: svc.Running, Accepts: cmdsAccepted}
				tick = fasttick
			default:
				elog.Error(1, fmt.Sprintf("unexpected control request #%d", c))
			}
		}
	}
	changes <- svc.Status{State: svc.StopPending}
	return
}

func main() {
	inService, err := svc.IsWindowsService()
	if err != nil {
		log.Fatalf("failed to determine if we are running in service: %v", err)
	}
	if inService {
		runService(svcName, false)
		return
	} else {
		fmt.Print("\n\tThis program must be run as a service \n\n")
	}

}

func runService(name string, isDebug bool) {
	var err error
	if isDebug {
		elog = debug.New(name)
	} else {
		elog, err = eventlog.Open(name)
		if err != nil {
			return
		}
	}
	defer elog.Close()

	elog.Info(1, fmt.Sprintf("starting %s service", name))
	run := svc.Run
	if isDebug {
		run = debug.Run
	}
	err = run(name, &myservice{})
	if err != nil {
		elog.Error(1, fmt.Sprintf("%s service failed: %v", name, err))
		return
	}
	elog.Info(1, fmt.Sprintf("%s service stopped", name))
}

func SpawnChilds() {
	Config := utils.ReadConfigFile("C:\\Program Files (x86)\\NaAV\\config.json")
	all_processes := append(Config.AnalysisTools, Config.HyperV.Processes...)
	all_processes = append(all_processes, Config.Other.Processes...)
	all_processes = append(all_processes, Config.Parallels.Processes...)
	all_processes = append(all_processes, Config.QEMU.Processes...)
	all_processes = append(all_processes, Config.VMware.Processes...)
	all_processes = append(all_processes, Config.VirtualBox.Processes...)
	utils.CreateFoldersPath("C:\\Program Files (x86)\\NaAV\\Temp")
	for _, process := range all_processes {
		target_file := fmt.Sprintf("C:\\Program Files (x86)\\NaAV\\Temp\\%s", process)
		exists, err := utils.FileExists(target_file)
		if err != nil {
			elog.Error(1, fmt.Sprintf("error checking if file %s exists: %s", target_file, err))
		}
		if exists {
			err := os.Remove(target_file)
			if err != nil {
				elog.Error(1, fmt.Sprintf("error removing file %s: %s", target_file, err))
			}
		}
		err = utils.CopyFile("C:\\Program Files (x86)\\NaAV\\dummyprogram.exe", target_file)
		if err != nil {
			elog.Error(1, fmt.Sprintf("error creating file %s: %s", target_file, err))
			continue
		}
		cmd := exec.Command(target_file)
		err = cmd.Start()
		if err != nil {
			elog.Error(1, fmt.Sprintf("error starting %s: %s", target_file, err))
			continue
		}
		ChildPids = append(ChildPids, cmd.Process.Pid)
	}
	elog.Info(1, fmt.Sprintf("spawned %d childs", len(ChildPids)))
}
func KillChilds() {
	for _, pid := range ChildPids {
		p, err := os.FindProcess(pid)
		if err != nil {
			elog.Error(1, fmt.Sprintf("error getting pid %d: %s", pid, err))
		}
		p.Kill()
		if err != nil {
			elog.Error(1, fmt.Sprintf("error killing pid %d: %s", pid, err))
		}
	}
}
