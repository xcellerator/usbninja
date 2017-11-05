# USB Ninja
## Serial Gadget

The `serial` gadget is the simplest gadget on the USB Ninja platform. By default, a TTY is made available over the connection and is the quickest way to get started. The default login is `alarm:alarm` and the root account is `root:root`.

On Linux and MacOS, you can connect to the TTY with either `screen`, `minicom` or some other program. Personally, I prefer `minicom`.
* `screen /dev/ttyACM0 115200`
* `minicom -b 115200 -D /dev/ttyACM0`

**NOTE:** You'll probably need root to be able to open the `/dev/ttyACM0` file.

On Windows, you can use [PuTTY](https://www.chiark.greenend.org.uk/~sgtatham/putty/latest.html) to connect. Just make sure the baud-rate is set to `115200` and you should be good to go!

Once you're connected, you can set up [WiFi](WIFI.md) for easier SSH access or whatever you want! The image contains a full [Arch Linux](https://archlinuxarm.org/) install, so you can treat it like any other system.
