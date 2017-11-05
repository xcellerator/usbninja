# USB Ninja
## WiFi Setup

Setting up WiFi is very easy and also very useful if you're preparing a more advanced attack using the USB Ninja.

1. Start the USB Ninja in `serial` mode by ensuring that `/boot/usbninja/options.txt` contains the line `gadget: serial`.
2. Connect to the serial interface in your preferred manner. See the [serial](SERIAL.md) page for more details.
3. Once you're in, run `sudo wifi-menu` and follow the prompts to connect to your wireless access point.
4. Check your new local IP with `ifconfig wlp6s0`.
5. Now you can disconnect from the serial interface and `ssh` into the USB Ninja over your wireless LAN
