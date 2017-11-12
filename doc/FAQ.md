# USB Ninja
## FAQ

**1. What are the login details to the default acccounts?**

The standard user account (and the only one accessible to `ssh`) is `alarm:alarm`. The root account is `root:root`.

**2. Why are you using Arch Linux ARM instead of Raspbian?**

Apart from the obvious meme points (btw, I use arch), I found that I could get the boot time down much lower on _ALARM_ than on Raspbian.

**3. I'm changing settings in `options.txt`, but they aren't making any difference when I reboot the Pi!**

First of all, make sure you're formatting the `options.txt` file in the way specified in [README.md](../README.md). If you're sure you aren't breaking any rules (if you do, it'll ignore your choices and use the defaults), then you might be having trouble actually saving the file. If you're accessing `options.txt` from the USB Ninja itself (i.e. through `ssh` or the serial gadget), then make sure you run `sync` before rebooting or power-cycling otherwise the cached data might not get written to disc properly. In general, its advisable to power off the USB Ninja, remove the MicroSD from the board and mount the first partition on another computer to edit `options.txt` on there instead. The same goes for `hid.txt`.

**4. I'm using the `hid_payload` gadget, but my strings aren't `PRINT`ing properly!**

This is most likely due to incompatible keyboard layouts. The issue lies in the fact that the scan codes sent as HID packets by the USB Ninja correspond to _physical keys_ and **not** _characters_. This is down to the way keyboards work and not something that can be directly controlled. The solution is to either wait for me or someone else to add multiple layout support, or to add it yourself (and send me a pull request!).

**5. I'm trying to use the `ethernet` gadget, but Windows won't recognise the adapter!**

As detailed on the [ethernet](ETHERNET.md) page, the `ethernet` gadget is just an alias for `ethernet_ecm` which is *not* supported by Windows (but is supported by everything else). To resolve the issue, use the `ethernet_rndis` gadget instead.

**6. I'm trying to use multiple gadgets at once, but it isn't working!**

Multiple gadget support is currently still in testing. As far as I have seen - it seems to work just fine, but if you've found a combination that doesn't work as expected (and you have reason to think it should), then please raise an issue on this repo.

**7. My `post` script isn't running even though the path is right in `options.txt`!**

If you're sure the path is right (note that you can specify either an absolute path or a path relative to `/boot/usbninja/`), then you might need to set the executable bit on the file. If your script is `/boot/usbninja/exec.sh`, then you can set the executable bit with the following command: `sudo chmod +x /boot/usbninja/exec.sh`
