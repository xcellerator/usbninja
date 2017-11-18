# USB Ninja - Testing

The USB Ninja is an advanced USB attack development platform designed to be as simple as possible whilst leaving as many options open for development. Written in Golang, it was developed on a Raspberry Pi Zero W, but should work on any device that supports USB OTG - see the list of [Supported Devices](doc/SUPPORTED.md). It makes heavy use of [configfs](https://www.kernel.org/doc/Documentation/filesystems/configfs/configfs.txt) to configure and present the drivers to the host system. See the [FAQ](doc/FAQ.md) for more information.

**Post-Setup Scripts can now be set! See [ADDED.md](ADDED.md) or [OPTIONS.md](doc/OPTIONS.md) for more information!**

**A new release with this new update is on the way! Stay tuned!**

Currently supported gadgets are:
* [USB-To-Serial](doc/SERIAL.md)
* [USB-To-Ethernet](doc/ETHERNET.md)
* [HID Emulation](doc/HID.md) (currently keyboard only)
* [Mass Storage](doc/STORAGE.md)

Download USB Ninja v1.0 from [here](https://mega.nz/#!H8E2gRrY!k2cv3XcMCp6nspSMlGr8vM6tULEOGLmcGLj2zAcT63w)!

|Filename|Size|Sha1sum|
|-|-|-|
|USBNinja_v1.7z|2.1G|495a180f6c92dfac0f8efb095bdf2f61f83c2dd0|
|USBNinja_v1.img|7.5G|57828c5b8ae61da53e110799c46cfdfa6eb21a50|

Please check out [ADDED.md](ADDED.md) for details about all the new features being added!

For examples of how to use the USB Ninja in different situations, check out my blog at [xcellerator.github.io](https://xcellerator.github.io)!

## Getting Started
### Flash the image
If you've got a Raspberry Pi Zero (W), then its as simple as flashing the image to an 8GB MicroSD card popping it in.

A simple `dd if=USBNinja_v1.img of=/dev/mmcblk0` should do the job.

By default, the USB Ninja will start up in `serial` mode with a baud rate of `115200`.

* Make sure the Micro USB end is plugged into the OTG port of the Pi (its labelled `USB` on the board)

* On Linux or MacOS, you can type `dmesg | tail` and should see something like `cdc_acm 1-2:1.0: ttyACM0: USB ACM device`.
..* E.g. `minicom -b 115200 -D /dev/ttyACM0` or `screen /dev/ttyACM0 115200`

* On Windows, use [PuTTY](https://www.chiark.greenend.org.uk/~sgtatham/putty/latest.html) to connect.
* The default login is **alarm:alarm**.

### Build from source
The other option is to just download and setup [Arch Linux ARM](https://archlinuxarm.org/platforms/armv6/raspberry-pi) by yourself and follow the instructions in [INSTALL.md](INSTALL.md) to compile the binaries from source and setup all the other services. The process is exactly the same as what was done to prepare the image.

## Using other gadgets
If you want to use gadgets other than plain old `serial` (and if you're here - you probably do), then all you need to do is edit `usbninja/options.txt` on the first partition of the MicroSD card. This corresponds to the `/boot` directory when the OS boots up.

See [OPTIONS.md](doc/OPTIONS.md) and [GADGETS.md](doc/GADGETS.md) for more information.
