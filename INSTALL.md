# USB Ninja
## Installing from Source

This guide assumes the following:
* You have flashed the basic [Arch Linux ARM](https://archlinuxarm.org/platforms/armv6/raspberry-pi) image to a MicroSD card
* You've configured WiFi using any one of the numerous guides online - see [here](https://archlinuxarm.org/forum/viewtopic.php?f=31&t=11529) for an example.
* You are using a Raspberry Pi Zero W.

### Install Required Packages
If you're unfamiliar with Arch Linux, then the only real difference (as far as we will be concerned) is the package manager, `pacman`.
1. First, `su` into root (remember, the default root password is **root**).
2. Next, update the entire system with `pacman -Syu`. This might take a little while as the RPi Zero is only small.
3. Now, you can install the required packages with `pacman -S sudo git dnsmasq golang`.

### Required Files
1. We need to create the working folder for USB Ninja: `mkdir /lib/usbninja`
* Optional Step: To make development easier, I prefer to change ownership of this file so I don't have to be root all the time: `chown -R alarm:alarm /lib/usbninja`
2. Create the USB Ninja config directory on partition one with: `mkdir /boot/usbninja`
* Again, you might want to `chown -R alarm:alarm /boot/usbninja`, but it isn't necessary.
3. Now you can create the Systemd service that will run at startup: `nano /etc/systemd/system/startup.service`:
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
4. We also want to add some lines to the bottom of `/etc/dnsmasq.conf`:
```
...
interface=usb0
dhcp-range=10.0.0.2,10.0.0.2
dhcp-option=3,10.0.0.1
```

### Building from Source
1. Now go back to `/lib/usbninja`. If you changed ownership to `alarm` then you can `exit` to leave the root account, otherwise stay as you are.
2. Clone this repo with `git clone https://github.com/xcellerator/usbninja.git build` and `cd build/src`.
3. Run the `./build_all.sh` script which will attempt to compile the binaries and copy them to `/lib/usbninja/`

### Enabling the Service
All that's left to do is enable `startup.service` so it runs at startup.
* `(sudo) systemctl enable startup.service`

### Setting your configurations
If you've followed these instructions and nothing more, then you don't have any options set! This is okay in the sense that everything will still work - the USB Ninja will startup in `serial` mode with all default settings. Most likely, you are going to want to change this, so you'll need to create an `options.txt` (and `hid.txt`) file in `/boot/usbninja/`. You can find a commented example in `/boot/usbninja/build/config/options.txt` or [here](config/options.txt). Alternatively, you can write your own by reading [OPTIONS.md](src/OPTIONS.md).
