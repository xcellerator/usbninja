# TODO

## Features
* Smart VID/PID decisions based on chosen gadget
* hid_remote gadget
* Add "one-stop-shop" mode that deletes options.txt after execution
* Multiple keyboard layouts
* ~~Post gadget-setup script execution~~ (See [ADDED.md](ADDED.md))
* Add some API functions for post-setup scripts to use
* ~~Add option to set alternative backing file in storage mode~~

## Code Changes
* **Find a better way to distribute releases other than an 8GB image!**
* Replace C style if IsTrue() == 0 statements with bools
* Replace config array with map for clarity - removes need for GetOption()
* Speed up boot time of Raspberry Pi
* Reduce "USB time" to under 15 seconds
* Impove WriteLine() in `extra.go`
* Write generic functions for `mkdir` and `ln`
* Merge sleep_str() and sleep_int()
* Replace WriteLine() with something sensible
