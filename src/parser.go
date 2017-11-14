// USB Ninja - parser.go
// Author: Harvey Phillips
// github.com/xcellerator/usbninja

package main

import (
	"io/ioutil"
	"strconv"
	"strings"
	"fmt"
)

func GetGadgets() []string {
	// Parse out chosen gadgets from options.txt
	var gadgets []string

	// Check for options.txt locaitons
	file_location := GetOptionLocation()

	// Open config file from supplied locations
	dat, _ := ioutil.ReadFile(file_location)
	// If dat is empty, then nothing will be set and the only gadget setup will be Serial
	configs := string(dat)

	// Split each line into the array to be parsed
	array := strings.Split(configs, "\n")

	// Loop through each line until we find one starting with "gadget: "
	for i := 0 ; i < len(array)-1; i++ {
		key_val := strings.Split(array[i], ":")

		if key_val[0] == "gadget" {
			// Now split the rest of the line by ","
			chosen_gadgets_untrimmmed := strings.Split(key_val[1], ",")

			// We most likely have trailing and/or leading spaces in each element
			for j := 0 ; j < len(chosen_gadgets_untrimmmed) ; j++ {
				gadgets = append(gadgets, strings.TrimSpace(chosen_gadgets_untrimmmed[j]))
			}
		}
	}

	return gadgets
}

func GetConfig() []string {
	// Parse out options from options.txt and return as an array
	var delay, vendorid, productid, serialnumber, manufacturer, productname, eth_hostaddr, eth_devaddr, post string

	// Check for options.txt locations
	file_location := GetOptionLocation()

	// Open config file from supplied location
	dat, _ := ioutil.ReadFile(file_location)
	// If dat is empty, then nothing will be set and the variables will take their defaults
	configs := string(dat)

	// Split each line into the array to parse
	array := strings.Split(configs, "\n")

	// Begin looping over each line
	for i := 0; i < len(array)-1; i++ {
		key_val := strings.Split(array[i], ":")

		if key_val[0] == "delay" {
			// Needs to be an integer value, if not set empty to be dealt with later on
			val := strings.TrimSpace(key_val[1])
			_, err := strconv.Atoi(val)
			if err == nil {
				delay = val
			} else {
				delay = "" // Will get set to default!
			}

		} else if key_val[0] == "vendorid" {
			// Needs to be two bytes long in hex, i.e 0x1d6b
			val := strings.TrimSpace(key_val[1])
			if IsValidID(val) == 0 {
				vendorid = val
			} else {
				vendorid = ""
			}

		} else if key_val[0] == "productid" {
			// Same as above
			val := strings.TrimSpace(key_val[1])
			if IsValidID(val) == 0 {
				productid = val
			} else {
				productid = ""
			}

		} else if key_val[0] == "serialnumber" {
			// Needs to be 8 hex bytes
			val := strings.TrimSpace(key_val[1])
			if len(val) == 16 && IsValidHex(val) == 0 {
				serialnumber = val
			} else {
				serialnumber = ""
			}

		} else if key_val[0] == "manufacturer" {
			// Not sure of the real limit, but any sensible string under 32 characters should be fine.
			val := strings.TrimSpace(key_val[1])
			if len(val) < 32 {
				manufacturer = val
			} else {
				manufacturer = ""
			}

		} else if key_val[0] == "productname" {
			// Same as above
			val := strings.TrimSpace(key_val[1])
			if len(val) < 32 {
				productname = val
			} else {
				productname = ""
			}

		} else if key_val[0] == "eth_hostaddr" {
			// len(key_val) should equal 7 and each octet should be a hex byte.
			// We also need to check that the first byte is even!
			// Then we create a slice from the last 6 of key_val and join to the new strings
			var hex_check int = 0
			var even_check int = 0
			first_byte, _ := strconv.Atoi(strings.TrimSpace(string(key_val[1])))

			for i := 1; i < len(key_val); i++ {
				if IsValidHex(strings.TrimSpace(string(key_val[i]))) != 0 {
					hex_check = 1
				}
			}

			if first_byte%2 != 0 {
				even_check = 1
			}

			if len(key_val) != 7 || hex_check != 0 || even_check != 0 {
				eth_hostaddr = ""
			} else {
				slice := key_val[1:]
				eth_hostaddr = strings.TrimSpace(strings.Join(slice, ":"))
			}

		} else if key_val[0] == "eth_devaddr" {
			// len(key_val) should equal 7 and each octet should be a hex byte.
			// We also need to check that the first byte is even!
			// Then we create a slice from the last 6 of key_val and join to the new strings
			var hex_check int = 0
			var even_check int = 0
			first_byte, _ := strconv.Atoi(strings.TrimSpace(string(key_val[1])))

			for i := 1; i < len(key_val); i++ {
				if IsValidHex(strings.TrimSpace(string(key_val[i]))) != 0 {
					hex_check = 1
				}
			}

			if first_byte%2 != 0 {
				even_check = 1
			}

			if len(key_val) != 7 || hex_check != 0 || even_check != 0 {
				eth_devaddr = ""
			} else {
				slice := key_val[1:]
				eth_devaddr = strings.TrimSpace(strings.Join(slice, ":"))
			}

			} else if key_val[0] == "post" {
				// Check if the file exists, and set post to empty otherwise

				val := strings.TrimSpace(key_val[1])
				val2 := fmt.Sprintf("/boot/usbninja/%s", val)

				if FileExist(val) {
					post = val
				} else if FileExist(val2) {
					post = val2
				} else {
					post = ""
				}
			}
		}

	if delay == "" {
		delay = SetDefaults("delay")
	}
	if vendorid == "" {
		vendorid = SetDefaults("vendorid")
	}
	if productid == "" {
		productid = SetDefaults("productid")
	}
	if serialnumber == "" {
		serialnumber = SetDefaults("serialnumber")
	}
	if manufacturer == "" {
		manufacturer = SetDefaults("manufacturer")
	}
	if productname == "" {
		productname = SetDefaults("productname")
	}
	if eth_hostaddr == "" {
		eth_hostaddr = SetDefaults("eth_hostaddr")
	}
	if eth_devaddr == "" {
		eth_devaddr = SetDefaults("eth_devaddr")
	}
	if (post == "" || post == "/boot/usbninja" || post == "/boot/usbninja/") {
		post = SetDefaults("post")
	}

	return []string{delay, vendorid, productid, serialnumber, manufacturer, productname, eth_hostaddr, eth_devaddr, post}
}

func SetDefaults(options string) string {
	if options == "delay" || options == "all" {
		return "0"
	}
	if options == "vendorid" || options == "all" {
		return "0x1d6b"
	}
	if options == "productid" || options == "all" {
		return "0x0104"
	}
	if options == "serialnumber" || options == "all" {
		return "fedcba9876543210"
	}
	if options == "manufacturer" || options == "all" {
		return "Generic Corp"
	}
	if options == "productname" || options == "all" {
		return "USB Device"
	}
	if options == "eth_hostaddr" || options == "all" {
		return "48:6f:73:74:50:43"
	}
	if options == "eth_devaddr" || options == "all" {
		return "42:61:64:55:53:42"
	}
	if options == "post" || options == "all" {
		return "echo"
	}
	return "error"
}

func IsValidHex(str string) int {
	// Returns 0 if and only if str contains only hex characters
	var returnval int = 0

	for _, r := range str {
		if (r < 'a' || r > 'f') && (r < '0' || r > '9') {
			returnval = 1
		}
	}

	return returnval
}

func IsValidID(str string) int {
	// Returns 0 if and only if str is of the form 0x???? where the ?'s are valid hex characters
	var returnval int

	if len(str) == 6 && string(str[0]) == "0" && string(str[1]) == "x" && IsValidHex(string(str[2:])) == 0 {
		returnval = 0
	} else {
		returnval = 1
	}

	return returnval
}

func IsAlpha(str string) int {
	// Returns 0 if and only if str is alpha (but "_" is also allowed)
	var returnval int = 0

	for _, r := range str {
		if (r < 'a' || r > 'z') && (r < 'A' || r > 'Z') && (r != '_') {
			returnval = 1
		}
	}

	return returnval
}

func IsLower(str string) int {
	// Returns 0 if and only if str is lowercase alpha
	var returnval int = 0

	for _, r := range str {
		if r < 'a' || r > 'z' {
			returnval = 1
		}
	}

	return returnval
}

func IsUpper(str string) int {
	// Returns 0 if and only if str is uppercase alpha
	var returnval int = 0

	for _, r := range str {
		if r < 'A' || r > 'Z' {
			returnval = 1
		}
	}

	return returnval
}

func IsShiftedSpecial(char string) int {
	// Returns 0 if and only if str is a special "shifted" character
	var returnval int

	if len(char) == 1 {
		switch char {
		case "¬", "!", "\"", "£", "$", "%", "^", "&", "*", "(", ")", "_", "+", "{", "}", ":", "@", "~", "<", ">", "?", "|":
			returnval = 0
		default:
			returnval = 1
		}
	}

	return returnval
}

func Special2Key(char string) string {
	// Takes a "shifted" special character and returns the key it sits on
	var lookup string

	if len(char) == 1 {
		switch char {
		case "!":
			lookup = "1"
		case "\"":
			lookup = "2"
		case "£":
			lookup = "3"
		case "$":
			lookup = "4"
		case "%":
			lookup = "5"
		case "^":
			lookup = "6"
		case "&":
			lookup = "7"
		case "*":
			lookup = "8"
		case "(":
			lookup = "9"
		case ")":
			lookup = "0"
		case "_":
			lookup = "-"
		case "+":
			lookup = "="
		case "{":
			lookup = "["
		case "}":
			lookup = "]"
		case ":":
			lookup = ";"
		case "@":
			lookup = "'"
		case "~":
			lookup = "#"
		case "<":
			lookup = ","
		case ">":
			lookup = "."
		case "?":
			lookup = "/"
		case "¬":
			lookup = "`"
		case "|":
			lookup = "\\"
		}
	}

	return lookup
}

func GetOption(config []string, option string) string {
	switch option {
	case "delay":
		return string(config[0])
	case "vendorid":
		return string(config[1])
	case "productid":
		return string(config[2])
	case "serialnumber":
		return string(config[3])
	case "manufacturer":
		return string(config[4])
	case "productname":
		return string(config[5])
	case "eth_hostaddr":
		return string(config[6])
	case "eth_devaddr":
		return string(config[7])
	case "post":
		return string(config[8])
	}
	return "error"
}
