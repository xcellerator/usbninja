# USB Ninja

The USB Ninja is an advanced USB attack development platform designed to be as simple as possible whilst leaving as many options open for development. Written in Golang, it was developed on a Raspberry Pi Zero W, but should work on any device that supports USB OTG - see the list of [Supported Devices](doc/SUPPORTED.md). It makes heavy use of [configfs](https://www.kernel.org/doc/Documentation/filesystems/configfs/configfs.txt) to configure and present the drivers to the host system. See the [FAQ](doc/FAQ.md) for more information.

Currently supported gadgets are:
* [USB-To-Serial](doc/SERIAL.md)
* [USB-To-Ethernet](doc/ETHERNET.md)
* [HID Emulation](doc/HID.md) (currently keyboard only)
* [Mass Storage](doc/STORAGE.md)

## Getting Started
I'm going to assume you've got a Raspberry Pi Zero (W) but it shouldn't be too difficult to adapt these instructions to a similar board - you'll just have to compile from source.

1. Download the image (or clone this repository) and write it to a MicroSD card in any way you're used to.
* `dd if=<FILE.IMG> of=/dev/mmcblk0`
2. Eject the card and pop it into the Raspberry Pi.
3. Make sure the micro USB end is plugged into the OTG port of the Pi (its labelled `USB` on the board itself)
4. By default the `USBNinja` will boot into `serial` mode.
5. A few seconds after plugging the Pi into a host machine, a USB serial device will be available.
* On Linux or MacOS, you can type `dmesg | tail` and should see something like `cdc_acm 1-2:1.0: ttyACM0: USB ACM device`.
* You can then use either `screen` or `minicom` to connect. The baud rate is `115200`.
* E.g. `minicom -b 115200 -D /dev/ttyACM0` or `screen /dev/ttyACM0 115200`
* On Windows, use [PuTTY](https://www.chiark.greenend.org.uk/~sgtatham/putty/latest.html) to connect.
* The default login is **alarm:alarm**.

## Using Different Gadgets
The configuration file resides in `/boot/usbninja/options.txt`, which corresponds to partition one on the MicroSD card - accessible on MacOS and Windows without setting up ext4 drivers. Open this up with any text editor and you will see the different options you can set along with some brief explanations. Each configuration is set by writing `option: value` and leaving a newline in between each.

**NOTE:** Any option can be omitted by this file and the defaults will automatically be chosen. If you wanted, you could just have `gadget: <YOUR_GADGET>` and nothing else.
Below are the currently supported options:
* `gadget`: The USB devices that are to be imitated, also known as "gadgets", separated by commas. See the [gadgets](doc/GADGETS.md) page for more information.
* `vendorid`: The VID that will be presented to the host machine. Along with the PID, this is very important as it will usually determine which drivers are loaded to handle the device by the host. Less of an issue in Linux, but is quite vital in Windows. Needs to be in the format: `0x????`. Defaults to **`0x1d6b`** (Linux Foundation).
* `productid`: The PID that will be presented to the host machine. Needs to be in the format `0x????`. Defaults to **`0x0104`** (Multifunction Composite Gadget).
* `serialnumber`: The serial number that will be presented to the host. Needs to be in the format `????????????????` (16 hex characters long). Defaults to **`fedcba9876543210`**.
* `manufacturer`: The manufacturer that will be presented to the host. Can be any string less than 32 characters long. Defaults to **`Generic Corp`**.
* `productname`: The product name that will be presented to the host. Can be any string less that 32 characters long. Defaults to **`USB Device`**.

There are also some gadget-specific options that will be ignored unless a certain gadget is used:
* `delay`: When using the `hid_payload` gadget, this is the number of milliseconds to leave in between individual keystrokes. You should only need to use this in very odd circumstances. See the [HID](doc/HID.md) page for more information. Needs to be an integer. Defaults to **`0`**.
* `eth_hostaddr`: When using one of the `ethernet` gadgets, this will be the MAC address presented to the system of the device _supposedly_ on the other end of the USB-To-Ethernet adapter. See the [Ethernet](doc/ETHERNET.md) page for more information. The only requirement here is that the first octet be an even number. Defaults to **`48:6f:73:74:50:43`**.
* `eth_devaddr`: When using one of the `ethernet` gadgets, this will be the MAC address presented to the system of the USB-To-Ethernet device itself. See the [Ethernet](doc/ETHERNET.md) page for more information. The only requirement here is that the first octet be an even number. Defaults to **`42:61:64:55:53:42`**.
