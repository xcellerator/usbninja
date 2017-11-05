// USB Ninja - keyboard.go
// Author: Harvey Phillips
// github.com/xcellerator/usbninja

package main

import (
	"io/ioutil"
	"strconv"
	"strings"
)

func Input2HID(char string) (string, string) {
	// Takes a line from hid_payload.txt without the instructive and returns a modifier and key1 string
	if len(char) == 1 {
		// Then we have a single letter and not another key
		if IsLower(char) == 0 {
			// We've got a lowercase character
			return ModifierBit(0, 0, 0, 0), hex_codes[char]

		} else if IsUpper(char) == 0 {
			// We've got an uppercase character
			return ModifierBit(0, 1, 0, 0), hex_codes[strings.ToLower(char)]

		} else if IsShiftedSpecial(char) == 0 {
			// We've got a "shifted" special character
			lookup := Special2Key(char)
			return ModifierBit(0, 1, 0, 0), hex_codes[lookup]

		} else if char != " " {
			// We must have a regular special character
			return ModifierBit(0, 0, 0, 0), hex_codes[char]

		} else {
			// We must have a space!
			return ModifierBit(0, 0, 0, 0), hex_codes["SPACE"]

		}
	} else if len(char) > 1 {
		// Then we have a named key and not a letter or punctuation
		return ModifierBit(0, 0, 0, 0), hex_codes[char]
	}
	return "", ""
}

func AssembleHIDPacket(modifier, key1 string) [8]string {
	// Takes the modifier and key1 string and returns an array of fields of the hid_packet
	var hid_packet [8]string

	// Loop through and set everything to "\0"
	for i := 0; i < 8; i++ {
		hid_packet[i] = "\\0"
	}

	// Now set the modifier and key1 bytes we received as input
	hid_packet[0] = modifier
	hid_packet[2] = key1

	return hid_packet
}

func SendPayload(payload_location, delay_str string) {
	// Takes the location of the hid_payload.txt (hardcoded?) and the delay between keystrokes. Prints the commands to be piped to bash.
	delay, _ := strconv.Atoi(delay_str)

	// Grab contents of hid_payload.txt
	dat, _ := ioutil.ReadFile(payload_location)
	payload_contents := string(dat)

	lines := strings.Split(payload_contents, "\n")

	// Grab the instructive from each line and branch accordingly
	for i := 0; i < len(lines); i++ {
		line := strings.Split(lines[i], " ")
		instructive := line[0]

		switch instructive {
		case "PRINT":
			str := strings.Join(line[1:], " ")
			SendString(str, delay)
		case "DELAY":
			sleep_str(line[1])
		case "ENTER", "RETURN":
			enter_key := []string{"ENTER"}
			SendKeys(enter_key, delay)
		case "CTRL", "ALT", "SHIFT", "GUI":
			SendKeys(line, delay)
		default:
			keytosend := []string{line[0]}
			SendKeys(keytosend, delay)
		}
	}

}

func SendString(str string, delay int) {
	// Convert each string into HID packets and send them separated by empty packets at the correct delay

	for i := 0; i < len(str); i++ {
		modifier, key1 := Input2HID(string(str[i]))
		hid_packet := AssembleHIDPacket(modifier, key1)

		// Now we need to join the hid_packet into a nice, neat string
		var hid_string string
		for j := 0; j < len(hid_packet); j++ {
			hid_string += hid_packet[j]
		}

		SendToUSB(hid_string, delay)
	}
}

func SendKeys(keys []string, delay int) {
	// Work out any needed modifier bits and format the named key as a hid_packet

	// Check is len(keys) is 1 - then there are no modifier keys (except maybe GUI?)
	if len(keys) == 1 {

		if keys[0] == "CTRL" || keys[0] == "SHIFT" || keys[0] == "ALT" || keys[0] == "GUI" {
			// Its just a modifier! No keys set!
			// NOTE: In most cases, this will do nothing, apart from GUI opening the start menu in Windows.

			var ctrlbit, shiftbit, altbit, guibit int

			if ExistInArray("CTRL", keys) == 0 {
				ctrlbit = 1
			}
			if ExistInArray("SHIFT", keys) == 0 {
				shiftbit = 1
			}
			if ExistInArray("ALT", keys) == 0 {
				altbit = 1
			}
			if ExistInArray("GUI", keys) == 0 {
				guibit = 1
			}

			modifier := ModifierBit(ctrlbit, shiftbit, altbit, guibit)
			hid_packet := AssembleHIDPacket(modifier, "\\0")
			hid_string := HIDPacket2String(hid_packet)

			SendToUSB(hid_string, delay)

		} else if keys[0] == "" {
			// Do nothing

		} else {
			modifier, key1 := Input2HID(keys[0])
			hid_packet := AssembleHIDPacket(modifier, key1)

			// We need to join the hid_packet into a string
			hid_string := HIDPacket2String(hid_packet)

			SendToUSB(hid_string, delay)
		}
	} else if len(keys) > 1 {
		// We have more than one key, so we have a modifier!
		var ctrlbit, shiftbit, altbit, guibit int

		// Check and set the modifier bits
		if ExistInArray("CTRL", keys) == 0 {
			ctrlbit = 1
		}

		if ExistInArray("SHIFT", keys) == 0 {
			shiftbit = 1
		}

		if ExistInArray("ALT", keys) == 0 {
			altbit = 1
		}

		if ExistInArray("GUI", keys) == 0 {
			guibit = 1
		}

		// Now we can set the modifier and key1 variables to be assembled into a hid packet
		modifier := ModifierBit(ctrlbit, shiftbit, altbit, guibit)
		_, key1 := Input2HID(keys[len(keys)-1])

		hid_packet := AssembleHIDPacket(modifier, key1)
		hid_string := HIDPacket2String(hid_packet)

		SendToUSB(hid_string, delay)
	}
}

func SendToUSB(hid_string string, delay int) {
	// Takes a hid_string and integer delay (in milliseconds) and sends it down the wire to the USB gadget

	empty_hid_string := "\\0\\0\\0\\0\\0\\0\\0\\0"

	presskey := EchoFormat(hid_string)
	releasekey := EchoFormat(empty_hid_string)

	ShellOut(presskey)
	sleep_int(delay)
	ShellOut(releasekey)
	sleep_int(delay)
}
