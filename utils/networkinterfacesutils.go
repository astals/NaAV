package utils

import (
	"errors"
	"fmt"
	"net"
	"os/exec"
)

/* requisites:
- https://docs.microsoft.com/en-us/virtualization/hyper-v-on-windows/quick-start/enable-hyper-v
- Enable-WindowsOptionalFeature -Online -FeatureName Microsoft-Hyper-V-Management-PowerShell */

/* if you know a beter solution please let me know (other than setupapi please, that one is a nightmare) */
/*
New-VMSwitch -Name "NaAV VBox" -AllowManagementOS $True -NetAdapterName "Ethernet"
Rename-NetAdapter "vEthernet (NaAV VBox)" -NewName "NaAV VBox"
Set-NetAdapter -Name "NaAV VBox" -MacAddress "08:00:27:e1:ee:e7" -Confirm:$false (ipconfig /all)

Remove-VMSwitch "NaAV VBox" -Confirm:$false
*/
func CreateNetworkAdapter(name string, mac string, printPrepend string) error {
	targetinterface, err := GetTargetAdapterForVirtualSwitch()
	if err != nil {
		PrintIfEnoughLevel(fmt.Sprintf("%s [!] ER-NU001 Error getting target adapter: %s", printPrepend, err), OPERATION_ERROR_MESSAGE)
		return err
	}
	_, err = exec.Command("powershell.exe", "-command", fmt.Sprintf("New-VMSwitch -Name \"%s\" -AllowManagementOS $True -NetAdapterName \"%s\"", name, targetinterface)).Output()
	if err != nil {
		PrintIfEnoughLevel(fmt.Sprintf("%s [!] ER-NU002 Error creting VMSwitch: %s", printPrepend, err), OPERATION_ERROR_MESSAGE)
		return err
	}
	_, err = exec.Command("powershell.exe", "-command", fmt.Sprintf("Rename-NetAdapter \"vEthernet (%s)\" -NewName \"%s\"", name, name)).Output()
	if err != nil {
		PrintIfEnoughLevel(fmt.Sprintf("%s [!] ER-NU003 Error reanming VMSwitch: %s", printPrepend, err), OPERATION_ERROR_MESSAGE)
		return err
	}
	_, err = exec.Command("powershell.exe", "-command", fmt.Sprintf("Set-NetAdapter -Name \"%s\" -MacAddress \"%s\" -Confirm:$false", name, mac)).Output()
	if err != nil {
		PrintIfEnoughLevel(fmt.Sprintf("%s [!] ER-NU004 Error setting mac address: %s", printPrepend, err), OPERATION_ERROR_MESSAGE)
		return err
	}
	return nil
}
func DeleteNetworkAdapter(name string, printPrepend string) error {
	_, err := exec.Command("powershell.exe", "-command", fmt.Sprintf("Remove-VMSwitch \"%s\" -Force", name)).Output()
	if err != nil {
		PrintIfEnoughLevel(fmt.Sprintf("%s [!] ER-NU004 Error creting VMSwitch: %s", printPrepend, err), OPERATION_ERROR_MESSAGE)
		return err
	}
	return nil
}

func ExistsNetworkAdapter(mac string, printPrepend string) {
	addresses, err := GetMacAddresses()
	if err != nil {
		PrintIfEnoughLevel(fmt.Sprintf("%s [!] ER-NU004 Error getting mac addresses: %s\n", printPrepend, err), OPERATION_ERROR_MESSAGE)
	}
	if ElementInStringArray(mac, addresses) {
		PrintIfEnoughLevel(fmt.Sprintf("%s [i] Detected %d of %d Network Adapters \n", printPrepend, 1, 1), SUMMARY_MESSAGE)
	} else {
		PrintIfEnoughLevel(fmt.Sprintf("%s [i] Detected %d of %d Network Adapters \n", printPrepend, 0, 1), SUMMARY_MESSAGE)
	}

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

func GetTargetAdapterForVirtualSwitch() (string, error) {
	var ifacesnames []string
	ifaces, err := net.Interfaces()
	if err != nil {
		return "", err
	}
	for _, i := range ifaces {
		ifacename := i.Name
		ifacesnames = append(ifacesnames, ifacename)
	}
	candidateifaces := []string{"Ethernet", "ethernet", "WiFi", "wifi"}
	for _, name := range candidateifaces {
		if ElementInStringArray(name, ifacesnames) {
			return name, nil
		}
	}
	return "", errors.New("Unable to find a candicate interface for the new VMSwitch")
}
