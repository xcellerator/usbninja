# USB Ninja
## Ethernet Gadgets

There are currently two different ethernet gadgets available; `ethernet_ecm` and `ethernet_rndis` (although `ethernet` is also recognised as an alias to `ethernet_ecm`).

### Gadget-Specific `options.txt` entries
You have two extra configurations available to you in `options.txt` if you are using one of the `ethernet` gadgets.
* `eth_hostaddr` - The MAC address of the host on the "other side" of the ethernet adapter.
* `eth_devaddr` - The MAC address of the ethernet adapter itself.

### Usage

The difference is very slight and is to do with some funkiness on the part of Windows. RNDIS (or _**R**emote **N**etwork **D**river **I**nterface **S**pecification_ to it's friends) is a proprietary Microsoft protocol for handling USB-To-Ethernet adapters. If you want to use the USB Ninja with a Windows host, you're going to need to use `ethernet_rndis`.

While most (if not all!) Linux distros (and MacOS) also support _RNDIS_, some embedded devices may not. That's where `ethernet_ecm` comes in! ECM (or _**E**thernet **C**ontrol **M**odel_) predates _RNDIS_ but is supported by pretty much _**everything**_ except for Windows. In general, if you aren't targetting Windows, use `ethernet_ecm`.

Once an ethernet gadget is active, a small DHCP server is started that will assign `10.0.0.2` to the host (you may need to run something like `dhclient` on Linux depending on your distro). The name of the network interface will also vary depending on your OS, but shouldn't be too hard to figure out. Once you're connected, you should be able to ssh into the Pi with:
* `ssh alarm@10.0.0.1` (default password is **alarm**).

**NOTE:** The network interface used for the connection on the Pi is `usb0`.
