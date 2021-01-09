package main

import (
	"bufio"
	"fmt"
	"os"

	"./modes"
	"./utils"
)

func main() {
	// TODO: create uninstall $het (control panel)
	args := os.Args
	if len(args) == 1 || args[1] == "-h" {
		DisplayHelp()
		return
	} else {
		res := CheckIfRunningWithElevatedPrivileges()
		if res == false {
			fmt.Printf("!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!! \n")
			fmt.Printf("Most actions like creation of fake drivers, network interfaces, registry keys, etc require administration pivileges \n")
			fmt.Printf("!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!! \n")
			reader := bufio.NewReader(os.Stdin)
			fmt.Printf("Continue anyway [y/n]:")
			text, _ := reader.ReadString('\n')
			if text != "y\r\n" {
				fmt.Printf("dasfads \n")
				return
			}
		}
	}
	if args[1] == "-c" {
		modes.Check()
	}
	if args[1] == "-u" {
		modes.Uninstall()
	}
	if args[1] == "-i" {
		//read config file args[2]
		modes.Install()
	}

}

func DisplayHelp() {
	fmt.Printf("---- NaAV 0.0 ----\n")
	fmt.Printf("naav.exe -i [configFile] -> Install, this acction requires a configuration file, you can see an example on https://github.com/astals/NaAV\n")
	fmt.Printf("naav.exe -u -> Uninstall\n")
	fmt.Printf("naav.exe -c -> Check, this action checks your system in order to know how many Virtual Machine checks it passes\n")
	fmt.Printf("naav.exe -h -> Help\n")

}

func CheckIfRunningWithElevatedPrivileges() bool {
	err := utils.CopyFile("resources\\dummy", "C:\\WINDOWS\\system32\\drivers\\NaAV_pigilege_check")
	if err == nil {
		return true
	}
	return false
}
