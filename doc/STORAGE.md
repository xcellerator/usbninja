# USB Ninja
## Mass Storage Gadget

The mass storage gadget is activated by putting `storage` on the `gadget: ` line in `options.txt`. By default, there is a total of 1GB available, but this can be changed by quite easily. The "storage" part of mass storage arises from the "backing file", one of which can be located in `/lib/usbninja/storage.img`. You can create your own backing file by following the instructions below, and then set the location in `options.txt`.

### Gadget-Specific `options.txt` entries
You have two extra (optional) configurations available to you in `options.txt`
* `storage_ro` - Determines whether the storage device will be presented as read-only or not. Can be set to either `yes` or `no`.
* `storage` - An absolute path to another backing file if the default one is too small for you.

**Note: You do not need to do this. There is an existing empty `storage.img` file already in place in `/lib/usbninja/storage.img`.**

### Creating your own backing file
1. Startup the USB Ninja in any mode **other** than `storage`.
2. Choose a folder you want to store the backing file in and `cd` to it. For example, `/home/alarm/`.
3. You can use `dd` to create an empty file with the size you want.
* E.g. For a 1GB backing file, you would run: `dd if=/dev/zero of=storage.img bs=1M count=1000`.
4. Now you have just a raw unformatted backing file that will appear to the host OS as a block device.
5. Open up `/boot/usbninja/options.txt` in your editor of choice (That means vim). You will need root (or sudo).
* Set `gadget: storage`.
* Set `storage: /home/alarm/storage.img` (adjust to where you have saved your backing file)
6. Run `sync` and then `reboot now`.
7. On the **Host OS**, you should now see the raw 1GB device show up either in Linux and MacOS under `/dev` or in Windows under "Device Manager".
8. You can format the device any way you are comfortable with, for example with a program like `gparted`.

**NOTE:** Make sure you format the device as `FAT32`. Other filesystems may work, but they have not been tested.
