package utils

import (
	"fmt"
	"net"
	"os/exec"
)

/* requisites:
- https://docs.microsoft.com/en-us/virtualization/hyper-v-on-windows/quick-start/enable-hyper-v
- Enable-WindowsOptionalFeature -Online -FeatureName Microsoft-Hyper-V-Management-PowerShell */

/* if you know a beter solution please let me know (other than setupapi please, that one is a nightmare) */
/*

New-VMSwitch  -Name "NaAV VBoxssss" -SwitchType "Internal"
Rename-NetAdapter "vEthernet (NaAV VBox)" -NewName "NaAV VBox"
Set-NetAdapter -Name "NaAV VBox" -MacAddress "08:00:27:e1:ee:e7" -Confirm:$false (ipconfig /all)

Remove-VMSwitch "NaAV VBox" -Confirm:$false
*/

func CreateNetworkAdapter(name string, mac string, printPrepend string) error {
	// TODO random mac adress if address already exists
	// TODO sanitaze input before calling powershell
	out, err := exec.Command("powershell.exe", "-command", fmt.Sprintf("New-VMSwitch  -Name \"%s\" -SwitchType \"Internal\"", name)).Output()
	if err != nil {
		PrintIfEnoughLevel(fmt.Sprintf("%s [!] ER-NU002 Error creting VMSwitch: %s\n", printPrepend, err), OPERATION_ERROR_MESSAGE)
		PrintIfEnoughLevel(fmt.Sprintf("%s%s Output: %s\n", printPrepend, printPrepend, out), OPERATION_ERROR_MESSAGE)
		return err
	}
	out, err = exec.Command("powershell.exe", "-command", fmt.Sprintf("Rename-NetAdapter \"vEthernet (%s)\" -NewName \"%s\"", name, name)).Output()
	if err != nil {
		PrintIfEnoughLevel(fmt.Sprintf("%s [!] ER-NU003 Error reanming VMSwitch: %s\n", printPrepend, err), OPERATION_ERROR_MESSAGE)
		PrintIfEnoughLevel(fmt.Sprintf("%s%s Output: %s\n", printPrepend, printPrepend, out), OPERATION_ERROR_MESSAGE)
		return err
	}
	out, err = exec.Command("powershell.exe", "-command", fmt.Sprintf("Set-NetAdapter -Name \"%s\" -MacAddress \"%s\" -Confirm:$false", name, mac)).Output()
	if err != nil {
		PrintIfEnoughLevel(fmt.Sprintf("%s [!] ER-NU004 Error setting mac address: %s\n", printPrepend, err), OPERATION_ERROR_MESSAGE)
		PrintIfEnoughLevel(fmt.Sprintf("%s%s Output: %s\n", printPrepend, printPrepend, out), OPERATION_ERROR_MESSAGE)
		return err
	}
	return nil
}

func DeleteNetworkAdapter(name string, printPrepend string) error {
	_, err := exec.Command("powershell.exe", "-command", fmt.Sprintf("Remove-VMSwitch \"%s\" -Force", name)).Output()
	if err != nil {
		PrintIfEnoughLevel(fmt.Sprintf("%s [!] ER-NU004 Error deleting VMSwitch: %s\n", printPrepend, err), OPERATION_ERROR_MESSAGE)
		return err
	}
	return nil
}

func ExistsNetworkAdapter(mac string, printPrepend string) bool {
	addresses, err := GetMacAddresses()
	if err != nil {
		PrintIfEnoughLevel(fmt.Sprintf("%s [!] ER-NU004 Error getting mac addresses: %s\n", printPrepend, err), OPERATION_ERROR_MESSAGE)
		return false
	}
	if ElementInStringArray(mac, addresses) {
		return true
	}
	return false
}

func GetMacAddresses() ([]string, error) {
	ifaces, err := net.Interfaces()
	if err != nil {
		return nil, err
	}
	var addresses []string
	for _, ifa := range ifaces {
		address := ifa.HardwareAddr.String()
		if address != "" {
			addresses = append(addresses, address)
		}
	}
	return addresses, nil
}

/*
New-VMSwitch -Name "NaAV VBox" -AllowManagementOS $True -NetAdapterName "Ethernet"
func GetTargetAdapterForVirtualSwitch() (string, error) {
	if TargetAdapterForVirtualSwitch != "" {
		return TargetAdapterForVirtualSwitch, nil
	}
	var ifacesnames []string
	ifaces, err := net.Interfaces()
	if err != nil {
		return "", err
	}
	for _, i := range ifaces {
		ifacename := i.Name
		fmt.Printf("i: %s\n", i.Name)
		ifacesnames = append(ifacesnames, ifacename)
	}
	candidateifaces := []string{"Ethernet", "ethernet", "WiFi", "wifi"}
	for _, name := range candidateifaces {
		if ElementInStringArray(name, ifacesnames) {
			TargetAdapterForVirtualSwitch = name
			return name, nil
		}
	}
	return "", errors.New("Unable to find a candicate interface for the new VMSwitch")
}*/
