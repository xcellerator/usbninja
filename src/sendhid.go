// USB Ninja - sendhid.go
// Author: Harvey Phillips
// github.com/xcellerator/usbninja

package main

import (
	"fmt"
	"time"
)

func main() {
	// Just send the hid_payload, no SetupGenericGadget()

	// First, get delay
	delay := GetConfig()[1]

	// Now the hid.txt location
	file_location := GetHID()

	// Now send payload!
	fmt.Println("Make sure you are running as root!")
	fmt.Printf("Sending HID Packet in...")
	fmt.Printf("3...")
	time.Sleep(time.Duration(1000) * time.Millisecond)
	fmt.Printf("2...")
	time.Sleep(time.Duration(1000) * time.Millisecond)
	fmt.Printf("1...\n")
	time.Sleep(time.Duration(1000) * time.Millisecond)

	SendPayload(file_location, delay)
}
