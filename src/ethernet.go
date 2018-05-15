// USB Ninja - ethernet.go
// Author: Harvey Phillips
// github.com/xcellerator/usbninja

package main

import (
	"fmt"
	"os"
	"os/exec"
)

func EthernetGadgetSetup(config []string, driver string) {

	// Grab the options from the config array
	eth_hostaddr := GetOption(config, "eth_hostaddr")
	eth_devaddr := GetOption(config, "eth_devaddr")

	// Sprintf the locations to write to later on
	// Its safe to Sprintf driver into the path as it is NOT user controlled!
	eth_hostaddr_loc := fmt.Sprintf("%sfunctions/%s.usb0/host_addr", configfs, driver)
	eth_devaddr_loc := fmt.Sprintf("%sfunctions/%s.usb0/dev_addr", configfs, driver)

	// Create the function
	function_loc := fmt.Sprintf("%sfunctions/%s.usb0", configfs, driver)
	function_ln_loc := fmt.Sprintf("%sconfigs/c.1/", configfs)
	CreateFunction := exec.Command("mkdir", "-p", function_loc)
	CreateFunctionError := CreateFunction.Run()

	// Check that the function got created without error
	if CreateFunctionError == nil {
		// Write to the files
		WriteLine(eth_hostaddr, eth_hostaddr_loc)
		WriteLine(eth_devaddr, eth_devaddr_loc)
	} else {
		os.Exit(1)
	}

	LinkFunction := exec.Command("ln", "-s", function_loc, function_ln_loc)
	LinkFunctionError := LinkFunction.Run()
	if (FileExist(function_ln_loc)) {
		fmt.Printf("LOG: (SUCCESS) Function linked successfully!\n")
	} else {
		fmt.Printf("LOG: (ERROR) Linking functions failed!\nExited with error: %v\n", LinkFunctionError)
		os.Exit(1)
	}

}

func EthernetPostSetup(config []string) {
	// Assuming that succeeded then we can configure the network device
	IfconfigCmd := exec.Command("ifconfig", "usb0", "10.0.0.1", "netmask", "255.255.255.252", "up")
	IfconfigCmdError := IfconfigCmd.Run()

	RouteCmd := exec.Command("route", "add", "-net", "default", "gw", "10.0.0.2")
	RouteCmdError := RouteCmd.Run()

	DnsmasqKill := exec.Command("killall", "dnsmasq")
	DnsmasqKillError := DnsmasqKill.Run()

	DnsmasqCmd := exec.Command("dnsmasq", "-i", "usb0", "-F", "10.0.0.2,10.0.0.2", "-O", "3,10.0.0.1")
	DnsmasqCmdError := DnsmasqCmd.Run()

	if DnsmasqKillError != nil {
		fmt.Printf("LOG: Didn't kill dnsmasq.. Trying to continue anyway...\n")
	}

	if IfconfigCmdError != nil || RouteCmdError != nil || DnsmasqCmdError != nil {
		fmt.Printf("LOG: (ERROR) Something post-finalizing went wrong!\nIfconfigCmdError: %v\nRouteCmdError: %v\nDnsmasqCmdError: %v\n")
		os.Exit(1)
	}
}
