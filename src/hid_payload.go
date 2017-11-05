// USB Ninja - hid_payload.go
// Author: Harvey Phillips
// github.com/xcellerator/usbninja

package main

import (
	"fmt"
	"os"
	"os/exec"
)

func HidPayloadSetup(config []string) {

	function_loc := fmt.Sprintf("%sfunctions/hid.usb0", configfs)
	function_ln_loc := fmt.Sprintf("%sconfigs/c.1/", configfs)

	MkdirCmd := exec.Command("mkdir", "-p", function_loc)
	MkdirCmdError := MkdirCmd.Run()

	if MkdirCmdError != nil {
		fmt.Printf("LOG: (ERROR) Unable to create folders for hid.usb0\nExited with error: %v\n", MkdirCmdError)
		os.Exit(1)
	}

	// Write protocols and class files for HID keyboard
	protocol_loc := fmt.Sprintf("%sfunctions/hid.usb0/protocol", configfs)
	subclass_loc := fmt.Sprintf("%sfunctions/hid.usb0/subclass", configfs)
	report_length_loc := fmt.Sprintf("%sfunctions/hid.usb0/report_length", configfs)
	report_desc_loc := fmt.Sprintf("%sfunctions/hid.usb0/report_desc", configfs)

	report_desc_content := "\"\\x05\\x01\\x09\\x06\\xa1\\x01\\x05\\x07\\x19\\xe0\\x29\\xe7\\x15\\x00\\x25\\x01\\x75\\x01\\x95\\x08\\x81\\x02\\x95\\x01\\x75\\x08\\x81\\x03\\x95\\x05\\x75\\x01\\x05\\x08\\x19\\x01\\x29\\x05\\x91\\x02\\x95\\x01\\x75\\x03\\x91\\x03\\x95\\x06\\x75\\x08\\x15\\x00\\x25\\x65\\x05\\x07\\x19\\x00\\x29\\x65\\x81\\x00\\xc0\""

	WriteLine("1", protocol_loc)
	WriteLine("1", subclass_loc)
	WriteLine("8", report_length_loc)
	WriteLine(report_desc_content, report_desc_loc)

	LinkFunctionCmd := exec.Command("ln", "-sf", function_loc, function_ln_loc)
	LinkFunctionCmdError := LinkFunctionCmd.Run()

	if FileExist(function_ln_loc) {
		fmt.Printf("LOG: (SUCCESS) Function linked successfully!\n")
	} else {
		fmt.Printf("LOG: (ERROR) Function linking failed!\nExited with status: %v\n", LinkFunctionCmdError)
	}
}

func HIDPostSetup(config []string) {
	fmt.Printf("LOG: Entered HIDPostSetup()!")
	// Get hid.txt location and delay from config
	hid_txt := GetHID()
	delay := GetOption(config, "delay")

	// Send payload!
	SendPayload(hid_txt, delay)
}
