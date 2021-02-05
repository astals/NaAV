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
func CreateNetworkAdapters(adapters map[string]string, printPrepend string) {
	okOperations := 0
	for name, mac := range adapters {
		err := CreateNetworkAdapter(name, mac, printPrepend)
		if err == nil {
			okOperations++
		}
	}
	PrintIfEnoughLevel(fmt.Sprintf("%s [i] Successfully performed %d of $d operations", printPrepend, okOperations, len(adapters)), SUMMARY_MESSAGE)

}
func CreateNetworkAdapter(name string, mac string, printPrepend string) error {
	// TODO random mac adress if address already exists
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
func DeleteNetworkAdapters(adapters map[string]string, printPrepend string) {
	okOperations := 0
	for name, _ := range adapters {
		err := DeleteNetworkAdapter(name, printPrepend)
		if err == nil {
			okOperations++
		}
	}
	PrintIfEnoughLevel(fmt.Sprintf("%s [i] Successfully performed %d of $d operations\n", printPrepend, okOperations, len(adapters)), SUMMARY_MESSAGE)
}

func DeleteNetworkAdapter(name string, printPrepend string) error {
	_, err := exec.Command("powershell.exe", "-command", fmt.Sprintf("Remove-VMSwitch \"%s\" -Force", name)).Output()
	if err != nil {
		PrintIfEnoughLevel(fmt.Sprintf("%s [!] ER-NU004 Error creting VMSwitch: %s", printPrepend, err), OPERATION_ERROR_MESSAGE)
		return err
	}
	return nil
}

func ExistsNetworkAdapters(adapters map[string]string, printPrepend string) {
	detected := 0
	for _, mac := range adapters {
		if ExistsNetworkAdapter(mac, printPrepend) {
			detected++
		}
	}
	PrintIfEnoughLevel(fmt.Sprintf("%s [i] Detected %d of %d mac addresses \n", printPrepend, detected, len(adapters)), 0)
}

func ExistsNetworkAdapter(mac string, printPrepend string) bool {
	addresses, err := GetMacAddresses()
	if err != nil {
		PrintIfEnoughLevel(fmt.Sprintf("%s [!] ER-NU004 Error getting mac addresses: %s", printPrepend, err), OPERATION_ERROR_MESSAGE)
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
		fmt.Print(address + "\n")
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
