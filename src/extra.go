// USB Ninja - extra.go
// Author: Harvey Phillips
// github.com/xcellerator/usbninja

package main

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"time"
)

var configfs string = "/sys/kernel/config/usb_gadget/usb/"

// Define all the hex codes to be available later on
//var hex_codes map[string]string = make(map[string]string)

var hex_codes = map[string]string{
	"a": "\\x4",
	"b": "\\x5",
	"c": "\\x6",
	"d": "\\x7",
	"e": "\\x8",
	"f": "\\x9",
	"g": "\\xa",
	"h": "\\xb",
	"i": "\\xc",
	"j": "\\xd",
	"k": "\\xe",
	"l": "\\xf",
	"m": "\\x10",
	"n": "\\x11",
	"o": "\\x12",
	"p": "\\x13",
	"q": "\\x14",
	"r": "\\x15",
	"s": "\\x16",
	"t": "\\x17",
	"u": "\\x18",
	"v": "\\x19",
	"w": "\\x1a",
	"x": "\\x1b",
	"y": "\\x1c",
	"z": "\\x1d",
	"1": "\\x1e",
	"!": "\\x1e",
	"2": "\\x1f",
	"@": "\\x1f",
	"3": "\\x20",
	//"#": "\\x20",
	"4":     "\\x21",
	"$":     "\\x21",
	"5":     "\\x22",
	"%":     "\\x22",
	"6":     "\\x23",
	"^":     "\\x23",
	"7":     "\\x24",
	"&":     "\\x24",
	"8":     "\\x25",
	"*":     "\\x25",
	"9":     "\\x26",
	"(":     "\\x26",
	"0":     "\\x27",
	")":     "\\x27",
	"ENTER": "\\x28",
	"ESC":   "\\x29",
	"DEL":   "\\x2a",
	"TAB":   "\\x2b",
	"SPACE": "\\x2c",
	"-":     "\\x2d",
	"_":     "\\x2d",
	"=":     "\\x2e",
	"+":     "\\x2e",
	"[":     "\\x2f",
	"{":     "\\x2f",
	"]":     "\\x30",
	"}":     "\\x30",
	"\\":    "\\x31",
	"|":     "\\x31",
	"#":     "\\x32",
	"~":     "\\x32",
	";":     "\\x33",
	":":     "\\x33",
	"'":     "\\x34",
	"\"":    "\\x34",
	"`":     "\\x35",
	//"~": "\\x35",
	",":         "\\x36",
	"<":         "\\x36",
	".":         "\\x37",
	">":         "\\x37",
	"/":         "\\x38",
	"?":         "\\x38",
	"CAPS":      "\\x39",
	"F1":        "\\x3a",
	"F2":        "\\x3b",
	"F3":        "\\x3c",
	"F4":        "\\x3d",
	"F5":        "\\x3e",
	"F6":        "\\x3f",
	"F7":        "\\x40",
	"F8":        "\\x41",
	"F9":        "\\x42",
	"F10":       "\\x43",
	"F11":       "\\x44",
	"F12":       "\\x45",
	"PRNTSCRN":  "\\x46",
	"SCROLLOCK": "\\x47",
	"PAUSE":     "\\x48",
	"INSERT":    "\\x49",
	"HOME":      "\\x4a",
	"PGUP":      "\\x4b",
	"DELFWD":    "\\x4c",
	"END":       "\\x4d",
	"PGDWN":     "\\x4e",
	"RIGHT":     "\\x4f",
	"LEFT":      "\\x50",
	"DOWN":      "\\x51",
	"UP":        "\\x52",
	"CLEAR":     "\\x53",
	"VOLUP":     "\\x80",
	"VOLDOWN":   "\\x81",
}

func ModifierBit(ctrl, shift, alt, gui int) string {
	// The modifier byte is built up from a bitfield with a 8-bit depth.
	// The high-end bits refer to the right modifier keys.
	// The low-end bits refer to the left modifier keys.
	// We only care about rh left ones, so only employ a 4-bit depth.
	var field int
	field = 0 + (ctrl) + (shift * 2) + (alt * 4) + (gui * 8)
	retstring := fmt.Sprintf("\\x%v", field)

	return retstring
}

func WriteLine(line string, file string) {
	// It's not pretty but it works - this could be much nicer!
	cmd := fmt.Sprintf("echo -ne %s > %s", line, file)

	ShellOut(cmd)
}

func sleep_str(delay_str string) {
	// Take a string, convert it to an integer and sleep for that many milliseconds.
	// This could/should be merged with the function below!
	delay, _ := strconv.Atoi(delay_str)
	time.Sleep(time.Duration(delay) * time.Millisecond)
}

func sleep_int(delay_int int) {
	time.Sleep(time.Duration(delay_int) * time.Millisecond)
}

func ExistInArray(str string, arr []string) int {
	// Returns 0 if and only if str is found in array
	returnval := 1

	for i := 0; i < len(arr); i++ {
		if arr[i] == str {
			returnval = 0
		}
	}

	return returnval
}

func EchoFormat(str string) string {
	// Returns str in the format "echo -ne <str> > /dev/hidg0"
	// For passing to ShellOut.
	// Would be nice to replace this, but not sure how to write to /dev/hidg0 properly
	prefix := "echo -ne \""
	suffix := "\" > /dev/hidg0"

	returnstring := fmt.Sprintf("%s%s%s", prefix, str, suffix)

	return returnstring
}

func HIDPacket2String(hid_packet [8]string) string {
	// Returns the contents of hid_packet as a string.
	var hid_string string
	for j := 0; j < len(hid_packet); j++ {
		hid_string += hid_packet[j]
	}

	return hid_string
}

func ShellOut(cmd string) bool {
	// Pass cmd to sh -c and check for an error
	// Very messy - only use if you really need to!
	// If you know the exact string to be executed, use exec.Command from os/exec

	ShellCommand := exec.Command("sh", "-c", cmd)
	ShellCommandError := ShellCommand.Run()

	if ShellCommandError != nil {
		return false
	}

	return true
}

func FileExist(file string) bool {
	// Determines if file exists on the local filesystem
	if _, err := os.Stat(file); err == nil {
		return true
	}

	return false
}

func GetHID() string {
	// Typically returns /boot/usbninja/hid.txt
	// Defaults to the one in /lib/ if not found
	var file_location string

	if FileExist("/boot/usbninja/hid.txt") {
		file_location = "/boot/usbninja/hid.txt"
	} else {
		file_location = "/lib/usbninja/config/hid.txt"
	}

	return file_location
}

func GetOptionLocation() string {
	// Typically returns /boot/usbninja/options.txt
	// Defaults to the one in /lib/ if not found
	file_location := "/boot/usbninja/options.txt"
	return file_location
}
