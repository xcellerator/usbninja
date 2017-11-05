# USB Ninja
## Mass Storage Gadget

The mass storage gadget is activated by putting `storage` on the `gadget: ` line in `options.txt`. By default, there is a total of 1GB available, but this can be changed by quite easily. The "storage" part of mass storage arises from the "backing file" which must be located in `/lib/usbninja/storage.img`. You can create your own backing file by following the instructions below.

### Creating your own backing file
1. Startup the USB Ninja in any mode **other** than `storage` so that we can safely modify the storage.img file.
2. `cd` over to `/lib/usbninja/`, where `storage.img` is located.
3. It's advisable to rename the existing `storage.img` to something else so that you can rename it back if something goes wrong; `mv storage.img storage_old.img`.
4. You can use `dd` to create an empty file with the size you want.
* E.g. For a 1GB backing file, you would run: `dd if=/dev/zero of=storage.img bs=1M count=1000`.
5. Now you have just a raw unformatted backing file that will appear to the host OS as a block device.
6. Set `gadget: storage` in `/boot/usbninja/options.txt`, `sync` and `reboot now` the USB Ninja to restart in mass storage mode.
7. On the **Host OS**, you should now see the raw 1GB device show up either in Linux and MacOS under `/dev` or in Windows under "Device Manager".
8. You can format the device any way you are comfortable with, for example with a program like `gparted`.

**NOTE:** Make sure you format the device as `FAT32`. Other filesystems may work, but they have not been tested.
