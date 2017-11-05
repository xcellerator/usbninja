// USB Ninja - stopall.go
// Author: Harvey Phillips
// github.com/xcellerator/usbninja

package main

import (
  "fmt"
  "os"
  "os/exec"
)

// Include extra.go

func main() {
  // Before anything, we need to check to see if we're running as root!
  euid := os.Geteuid()
  if ( euid != 0 ) {
    fmt.Println("Please run as root!")
    os.Exit(1)
  }

  // First, we stop any services running via systemd that could have been activated!
  // Stop TTY for serial
  StopTTYCmd := exec.Command("systemctl", "stop", "getty@ttyS0.service")
  StopTTYCmdError := StopTTYCmd.Run()

  // Stop DNSMasq for ethernet adapter
  StopDNSMasq := exec.Command("systemctl", "stop", "dnsmasq.service")
  StopDNSMasqError := StopDNSMasq.Run()

  // Bring down usb0 interface
  BringDownIfCmd := exec.Command("ifconfig", "usb0", "down")
  BringDownIfCmdError := BringDownIfCmd.Run()

  // Check for errors
  if ( StopTTYCmdError != nil || StopDNSMasqError != nil || BringDownIfCmdError != nil ){
    fmt.Println("Failed to stop services before bringing down UDC!")
    fmt.Printf("StopTTYCmdError = %v\nStopDNSMasqError = %v\nBringDownIfCmdError = %v\n", StopTTYCmdError, StopDNSMasqError, BringDownIfCmdError)
    fmt.Printf("\nTrying to continue anyway...\n")
  }

  // Disable all gadgets
  fmt.Printf("LOG: Disabling gadgets\n")
  UDC_loc := fmt.Sprintf("%sUDC", configfs)
  WriteLine("", UDC_loc)

  // Remove linked functions from configurations
  fmt.Printf("LOG: Removing linked functions\n")
  RmLns := fmt.Sprintf("rm %sconfigs/c.1/*.*", configfs)
  ShellOut(RmLns)

  // Remove configuration strings and the configuration itself
  fmt.Printf("LOG: Removing configurations\n")
  RmConfs := fmt.Sprintf("rmdir %sconfigs/c.1/strings/0x409 %sconfigs/c.1", configfs, configfs)
  ShellOut(RmConfs)

  // Remove function directories and functions strings
  fmt.Printf("LOG: Removing functions\n")
  RmFuncCmd := fmt.Sprintf("rmdir %sfunctions/*.* %sstrings/0x409", configfs, configfs)
  ShellOut(RmFuncCmd)

  // Finally remove the gadget!
  fmt.Printf("LOG: Removing gadgets\n")
  RmGadgetCmd := fmt.Sprintf("rmdir /sys/kernel/config/usb_gadget/usb")
  ShellOut(RmGadgetCmd)

  fmt.Printf("LOG: Done!\n")
}
