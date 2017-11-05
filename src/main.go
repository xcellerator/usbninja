// USB Ninja - main.go
// Author: Harvey Phillips
// github.com/xcellerator/usbninja

package main

import (
	"fmt"
	//"strings"
)

func main() {
	config := GetConfig()

	// Check which gagdet is set and branch accordingly
	gadgets := GetGadgets()

	fmt.Printf("LOG: Options set as:\n=> delay: %s\n=> vendorid: %s\n=> productid: %s\n=> serialnumber: %s\n=>manufacturer: %s\n=>productname: %s\n\n", GetOption(config, "delay"), GetOption(config, "vendorid"), GetOption(config, "productid"), GetOption(config, "serialnumber"), GetOption(config, "manufacturer"), GetOption(config, "productname"))
	fmt.Printf("LOG: Gadgets set as: %v\n", gadgets)

	// SetupGenericGadget
	SetupGenericGadget(config)

	// Check that gadgets isn't empty
	if len(gadgets) == 0 {
		// If it is, just start the Serial gadget
		gadgets = append(gadgets, "serial")
	}

	// Loop through gadgets array
	for i := 0 ; i < len(gadgets) ; i++ {
		// Initial setup for each gadget
		switch gadgets[i] {
		case "serial":
			SerialGadgetSetup(config)
		case "ethernet", "ethernet_ecm":
			EthernetGadgetSetup(config, "ecm")
		case "ethernet_rndis":
			EthernetGadgetSetup(config, "rndis")
		case "hid_payload":
			HidPayloadSetup(config)
		case "hid_remote":
			HidRemoteSetup(config)
		case "storage":
			StorageGadgetSetup(config)
		default:
			fmt.Printf("LOG: (ERROR) No gagdet found to setup!\n")
		}
	}

	// Now we can finalize UDC
	FinalizeGadgetSetup()

	for j := 0 ; j < len(gadgets) ; j++ {
		// Post setup triggers for selected gadgets
		switch gadgets[j] {
		case "serial":
			SerialPostSetup(config)
		case "ethernet", "ethernet_ecm", "ethernet_rndis":
			EthernetPostSetup(config)
		case "hid_payload":
			HIDPostSetup(config)
		}
	}
}
