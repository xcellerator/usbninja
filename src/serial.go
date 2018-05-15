// USB Ninja - serial.go
// Author: Harvey Phillips
// github.com/xcellerator/usbninja

package main

import (
  "os"
  "os/exec"
  "fmt"
)


func SerialGadgetSetup(config []string) {
  fmt.Printf("LOG: Preparing serial specific configurations\n")

  function_loc := fmt.Sprintf("%sfunctions/acm.usb0", configfs)
  function_ln_loc := fmt.Sprintf("%sconfigs/c.1/", configfs)

  fmt.Printf("LOG: Trying to create functions/acm.usb0 directory\n")
  CreateFunction := exec.Command("mkdir", "-p", function_loc)
  CreateFunctionError := CreateFunction.Run()

  // Check to make sure it didn't fail
  if ( CreateFunctionError != nil ) {
    fmt.Printf("LOG: (ERROR) Unable to create directory\nExited with status: %v\n", CreateFunctionError)
    os.Exit(1)
  } else {
    fmt.Printf("LOG: (SUCCESS) Directory created successfully\n")
    LinkFunction := exec.Command("ln", "-sf", function_loc, function_ln_loc)
    LinkFunctionError := LinkFunction.Run()
    if ( FileExist(function_ln_loc) ) {
      fmt.Printf("LOG: (SUCCESS) Function linked successfully!\n")
    } else {
      fmt.Printf("LOG: (ERROR) Unable to link function\nExited with error: %v\n", LinkFunctionError)
      os.Exit(1)
    }
  }
}

func SerialPostSetup(config []string) {

  // Now the gadget is live, we can fire up a TTY on the device
  fmt.Printf("LOG: Starting a TTY on /dev/ttyGS0\n")
  StartService := exec.Command("agetty", "-L", "ttyGS0", "115200", "vt100")
  StartServiceError := StartService.Run()

  if ( StartServiceError != nil ) {
    fmt.Printf("LOG: (ERROR) Unable to start service\nExited with error: %v\n", StartServiceError)
    os.Exit(1)
  } else {
    fmt.Printf("LOG: (SUCCESS) Service started\n")
  }
}
