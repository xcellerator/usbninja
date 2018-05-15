// USB Ninja - generic_gadget.go
// Author: Harvey Phillips
// github.com/xcellerator/usbninja

package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func SetupGenericGadget(config []string) int {
	fmt.Printf("LOG: Entered SetupGenericGadget()!\n")
	// First we grab all the config options from the array
	vendorid := GetOption(config, "vendorid")
	productid := GetOption(config, "productid")
	serialnumber := GetOption(config, "serialnumber")
	manufacturer := GetOption(config, "manufacturer")
	productname := GetOption(config, "productname")

	// Now set all the file locations
	idVendor_loc := fmt.Sprintf("%sidVendor", configfs)
	idProduct_loc := fmt.Sprintf("%sidProduct", configfs)
	bcdDevice_loc := fmt.Sprintf("%sbcdDevice", configfs)
	bcdUSB_loc := fmt.Sprintf("%sbcdUSB", configfs)

	serialnumber_loc := fmt.Sprintf("%sstrings/0x409/serialnumber", configfs)
	manufacturer_loc := fmt.Sprintf("%sstrings/0x409/manufacturer", configfs)
	productname_loc := fmt.Sprintf("%sstrings/0x409/product", configfs)

	configuration_loc := fmt.Sprintf("%sconfigs/c.1/strings/0x409/configuration", configfs)
	maxpower_loc := fmt.Sprintf("%sconfigs/c.1/MaxPower", configfs)

	// Try and create the required directories
	fmt.Printf("LOG: Trying to mount configfs\n")
	MountConfigFS := exec.Command("mount", "-t", "configfs", "none", "/sys/kernel/config/")
	MountConfigFSError := MountConfigFS.Run()

	fmt.Printf("LOG: Trying to create configfs directory\n")
	MakeDirUSB := exec.Command("mkdir", "-p", configfs)
	MakeDirUSBError := MakeDirUSB.Run()

	fmt.Printf("LOG: Trying to create strings/0x409 directory\n")
	strings_loc := fmt.Sprintf("%sstrings/0x409", configfs)
	MakeDirStrings := exec.Command("mkdir", "-p", strings_loc)
	MakeDirStringsError := MakeDirStrings.Run()

	fmt.Printf("LOG: Trying to create configs/c.1/strings/0x409 directory\n")
	c1strings_loc := fmt.Sprintf("%sconfigs/c.1/strings/0x409/", configfs)
	MakeDirC1Strings := exec.Command("mkdir", "-p", c1strings_loc)
	MakeDirC1StringsError := MakeDirC1Strings.Run()

	// If creating the directories failed, exit
	if MountConfigFSError != nil {
		fmt.Printf("LOG: Failed to mount configfs! Trying to continue anyway...\n")
	}

	if MakeDirUSBError != nil || MakeDirStringsError != nil || MakeDirC1StringsError != nil {
		fmt.Printf("LOG: (FAIL) Unable to create one or more of the directories\n")
		fmt.Printf("%s\n", MakeDirUSBError)
		os.Exit(1)
	} else {
		fmt.Printf("LOG: (SUCCESS) Directories created successfully\n")
		fmt.Printf("LOG: Writing options from config.txt to device descriptors...\n")
		WriteLine(vendorid, idVendor_loc)
		WriteLine(productid, idProduct_loc)
		WriteLine("0x0100", bcdDevice_loc)
		WriteLine("0x0200", bcdUSB_loc)
		WriteLine(serialnumber, serialnumber_loc)
		WriteLine(manufacturer, manufacturer_loc)
		WriteLine(productname, productname_loc)
		WriteLine("Config 1: ECM Network", configuration_loc)
		WriteLine("250", maxpower_loc)
		fmt.Printf("LOG: Done\n")
	}

	return 0
}

func FinalizeGadgetSetup() {
	fmt.Printf("LOG: Finalizing by writing UDC to descriptor\n")
	output, _ := exec.Command("ls", "/sys/class/udc").Output()
	file := fmt.Sprintf("%sUDC", configfs)

	udc_result := strings.TrimSuffix(string(output), "\n")

	WriteLine(udc_result, file)
	fmt.Printf("LOG: USB Gadget now active\n")
}
