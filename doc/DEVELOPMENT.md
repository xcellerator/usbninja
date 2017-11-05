# USB Ninja
## Development information

If you're looking to add functionality to the USB Ninja or devise a more advanced payload than what the basic `options.txt` provides, then you'll probably find this document useful.

### Binaries
All the binaries are located in `/lib/usbninja/` and are as follows:
* `main` - The binary ran by `startup.service` that does all the work setting up the gadgets with the right configurations.
* `sendhid` - This will just send what's in `hid.txt` via `/dev/hidg0`. You need the `hid_payload` gadget to be already active on the host for this to work. Requires root.
* `stopall` - This will disable and remove all gadgets from `configfs`. Useful to change gadgets without having to reboot the USB Ninja. Note however, that occasionally the host can get confused (has happened to me a few times on Linux) so you might still have to unplug the device and plug it back in. **Can be temperamental at times!**

### Source Code
As well as being available in this repo, the source is included in `/lib/usbninja/src/`. You can use the build scripts to compile the source code and copy the binaries up a directory.
* `build_main.sh` - Builds `main` only
* `build_sendhid.sh` - Builds `sendhid` only
* `build_stopall.sh` - Builds `stopall` only
* `build_all.sh` - Builds everything

### Configuration
If you delete the `options.txt` (and `hid.txt`) from `/boot/usbninja/` then `main` will look to `/lib/usbninja/config/` for them instead. You might find it useful to have everything close by if you are constantly changing these files.

### Systemd
The `startup.service` lives in `/etc/systemd/system/startup.service` and contains the following:
```
[Unit]
Description=Startup Script for the USB Ninja
After=basic.target
Before=network.target

[Service]
ExecStart=/lib/usbninja/main

[Install]
WantedBy=basic.target
```
If you are devising an attack that incorporates some kind of script that is to be ran after `main` finishes executing, then currently you can just create your own `.service` file, but be sure to add the line `After=startup.service` in the `[Unit]` segment. Better support for this situation is on the way.
