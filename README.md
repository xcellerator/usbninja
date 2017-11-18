# USB Ninja

The USB Ninja is an advanced USB attack development platform designed to be as simple as possible whilst leaving as many options open for development. Written in Golang, it was developed on a Raspberry Pi Zero W, but should work on any device that supports USB OTG - see the list of [Supported Devices](doc/SUPPORTED.md). It makes heavy use of [configfs](https://www.kernel.org/doc/Documentation/filesystems/configfs/configfs.txt) to configure and present the drivers to the host system. See the [FAQ](doc/FAQ.md) for more information.

Currently supported gadgets are:
* [USB-To-Serial](doc/SERIAL.md)
* [USB-To-Ethernet](doc/ETHERNET.md)
* [HID Emulation](doc/HID.md) (currently keyboard only)
* [Mass Storage](doc/STORAGE.md)

Please check out [ADDED.md](ADDED.md) for details about all the new features being added!

For examples of how to use the USB Ninja in different situations, check out my blog at [xcellerator.github.io](https://xcellerator.github.io)!

## Getting Started
Getting setup with the USBNinja is pretty simple, as long as you're happy with the Linux command-line.
1. First of all, head over to <LINK> to download the latest release files. You need both `root.tar.gz` and `boot.tar.gz`.
2. Create a new folder (doesn't matter where) to act as our working directory. You'll only need this during the first setup.
3. Create two folders within this called `boot` and `root` and then copy `boot.tar.gz` and `root.tar.gz` to the working folder.
* i.e. You should have `boot`, `boot.tar.gz`, `root`, and `root.tar.gz` sitting in your working folder.
4. Insert your MicroSD card into your computer and check whats its called under `/dev/` using a quick `lsblk`. I use an SD adapter, so its called `mmcblk0` for me.
5. Now fire up fdisk with `sudo fdisk /dev/mmcblk0`.
   * Type `o`. This will wipe any existing partitions on the card.
   * Type `n`, then `p` for a primary partition.
   * Type `1` for the first partition, then **ENTER** for the default starting sector, then `+100M` for a 100MB boot partition.
   * Type `t`, then `c` to set the first partition as "W95 FAT32 (LBA)".
   * Type `n`, then `p` for another primary partition.
   * Type `2` for the second partition, then **ENTER** twice to fill the rest of the card.
   * Now you can write the new partition table with `w`.
6. Next, we create the filesystems on the two partitions:
   * `sudo mkfs.vfat /dev/mmcblk0p1`
   * `sudo mkfs.ext4 /dev/mmcblk0p2`
7. Now, mount the two partitions to our `root` and `boot` folders:
   * `sudo mount /dev/mmcblk0p1 boot`
   * `sudo mount /dev/mmcblk0p2 root`
8. Extract the release to the mounted partitions
   * `sudo bsdtar -xpf boot.tar.gz`
   * `sudo bsdtar -xpf root.tar.gz`
9. Finally, you can `sync` (this will take a little while - don't worry!) and then `sudo umount boot root` to unmount the MicroSD card.
10. Pop the MicroSD card back into the Raspbery Pi and boot it up! Make sure you use the "USB OTG" port on the board (the one next to the mini-HDMI port).
11. After a few seconds, you should see a new serial device show up in `dmesg`. You can use either `screen` or `minicom` to access it. E.g. `sudo minicom -b


### Build from source
The other option is to just download and setup [Arch Linux ARM](https://archlinuxarm.org/platforms/armv6/raspberry-pi) by yourself and follow the instructions in [INSTALL.md](INSTALL.md) to compile the binaries from source and setup all the other services. The process is exactly the same as what was done to prepare the image.

## Using other gadgets
If you want to use gadgets other than plain old `serial` (and if you're here - you probably do), then all you need to do is edit `usbninja/options.txt` on the first partition of the MicroSD card. This corresponds to the `/boot` directory when the OS boots up.

See [OPTIONS.md](doc/OPTIONS.md) and [GADGETS.md](doc/GADGETS.md) for more information.
