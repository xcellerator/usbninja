// USB Ninja - storage.go
// Author: Harvey Phillips
// github.com/xcellerator/usbninja

package main

// We expect 'storage.img' to be present in /lib/usbninja/

import (
  "fmt"
  "os"
  "os/exec"
)

func StorageGadgetSetup(config []string) {

  storageimg := "/lib/usbninja/storage.img"

  // Sprintf the directory locations
  function_loc := fmt.Sprintf("%sfunctions/mass_storage.usb0", configfs)
  function_ln_loc := fmt.Sprintf("%sconfigs/c.1/", configfs)

  // Make function directory
  MkdirCmd := exec.Command("mkdir", "-p", function_loc)
  MkdirCmdError := MkdirCmd.Run()

  // Check if it failed
  if (MkdirCmdError != nil) {
    fmt.Printf("LOG: (ERROR) Failed to make function directory!\nExited with: %v\n", MkdirCmdError)
  } else {
    fmt.Printf("LOG: (SUCCESS) Created function directories successfully!\n")
  }

  // Write descriptors
  fmt.Printf("LOG: Writing USB descriptors\n")

  stall_loc := fmt.Sprintf("%sfunctions/mass_storage.usb0/stall", configfs)
  cdrom_loc := fmt.Sprintf("%sfunctions/mass_storage.usb0/lun.0/cdrom", configfs)
  ro_loc := fmt.Sprintf("%sfunctions/mass_storage.usb0/lun.0/ro", configfs)
  nofua_loc := fmt.Sprintf("%sfunctions/mass_storage.usb0/lun.0/nofua", configfs)
  file_loc := fmt.Sprintf("%sfunctions/mass_storage.usb0/lun.0/file", configfs)

  WriteLine("1", stall_loc)
  WriteLine("0", cdrom_loc)
  WriteLine("0", ro_loc)
  WriteLine("0", nofua_loc)
  WriteLine(storageimg, file_loc)

  fmt.Printf("LOG: Done\n")

  // Link function to configurations
  LinkFunction := exec.Command("ln", "-sf", function_loc, function_ln_loc)
  LinkFunctionError := LinkFunction.Run()

  // Check to see if it failed
  if FileExist(function_ln_loc) {
    fmt.Printf("LOG: (SUCCESS) Function linked successfully!\n")
  } else {
    fmt.Printf("LOG: (ERROR) Function failed to link!\nExited with: %v", LinkFunctionError)
    os.Exit(1)
  }
}
