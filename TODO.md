# TODO

## Features
* ~~Mass Storage Support~~
* ~~Composite Device Support~~
* Smart VID/PID decisions based on chosen gadget
* ~~Multiple drivers for USB-To-Ethernet gadget (RNDIS, ECM, etc)~~
* hid_remote gadget
* ~~DHCP for ethernet gadget~~
* Add "one-stop-shop" mode that deletes options.txt after execution

## Code Changes
* Replace C style if IsTrue() == 0 statements with bools
* Replace config array with map for clarity - removes need for GetOption()
* Speed up boot time of Raspberry Pi
* Reduce "USB time" to under 15 seconds
* Impove WriteLine() in `extra.go`
* ~~Find out why LinkFunctionError != nil for serial gadget~~
* Write generic functions for `mkdir` and `ln`
* Merge sleep_str() and sleep_int()
* Replace WriteLine() with something sensible
