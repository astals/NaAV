package main

import (
	"fmt"
	"os"

	"./modes"
	"./utils"
)

var NaAVVersion = "0.1"

func main() {
	// TODO: create uninstall $het (control panel)
	args := os.Args
	if len(args) == 1 || args[1] == "-h" || args[1] == "--help" {
		DisplayHelp()
		return
	}
	if args[1] == "-v" || args[1] == "--version" {
		fmt.Printf("NaAV %s\n", NaAVVersion)
		return
	}
	res := CheckIfRunningWithElevatedPrivileges()
	if res == false {
		fmt.Printf("!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!! \n")
		fmt.Printf("Most actions (creation of fake drivers, network interfaces, etc) require administration pivileges \n")
		fmt.Printf("!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!! \n")
		/*reader := bufio.NewReader(os.Stdin)
		fmt.Printf("Continue anyway [y/n]:")
		text, _ := reader.ReadString('\n')
		if text != "y\r\n" {
			fmt.Printf("dasfads \n")
			return
		}*/
	}
	if args[1] == "--check" {
		modes.Check()
	}
	if args[1] == "--uninstall" {
		modes.Uninstall()
	}
	if args[1] == "--install" {
		if len(args) < 3 {
			fmt.Printf("This action requires a configuration file\n")
			DisplayHelp()
			return
		}
		modes.Install(args[2])
	}
	if args[1] != "-v" && args[1] != "--version" && args[1] != "-h" && args[1] != "--help" && args[1] != "--check" && args[1] != "--uninstall" && args[1] != "--install" {
		fmt.Printf("Invalid option \n")
		DisplayHelp()
	}
}

func DisplayHelp() {
	fmt.Printf("---- NaAV %s ----\n", NaAVVersion)
	fmt.Printf("naav.exe --install [configFile] -> Install, this action requires a configuration file, you can see an example on https://github.com/astals/NaAV/blob/main/config.json\n")
	fmt.Printf("naav.exe --uninstall -> Uninstall, to uninstall is recommended running 'C:\\Program Files (x86)\\NaAV\\naav.exe -u' instead of using the downloaded file\n")
	fmt.Printf("naav.exe --check -> Check, this action checks your system in order to know how many Virtual Machine checks it passes\n")
	fmt.Printf("naav.exe -v/--version -> Versions (installed and current binary) \n")
	fmt.Printf("naav.exe -h/--help -> Help\n")

}

func CheckIfRunningWithElevatedPrivileges() bool {
	err := utils.CopyFile("resources\\dummy", "C:\\WINDOWS\\system32\\drivers\\NaAV_pigilege_check")
	if err == nil {
		return true
	}
	return false
}
