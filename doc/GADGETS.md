# USB Ninja
## Gadgets

Gadgets are configured by editing the `gadget: ` line in `/boot/usbninja/options.txt`. Currently, multiple gadget support is still in testing (see [FAQ](FAQ.md) number 6.), but it seems to work very well. You can specify multiple gadgets by separating them with commas.

**E.g.** `gadget: serial, storage, ethernet_ecm`

### USB-To-Serial Adapter
To use the USB-To-Serial gadget, set `gadget` to `serial` in `options.txt`. A TTY will be made available via `/dev/ttyGS0` on the USB Ninja which shows up as `/dev/ttyACM0` on the (Linux) host. If using Windows, you can use [PuTTY](https://www.chiark.greenend.org.uk/~sgtatham/putty/latest.html). Either way, ensure that the baud rate is set to `115200`. See [SERIAL.md](SERIAL.md) for more information.

### USB-To-Ethernet Adapter
To use the USB-To-Ethernet gadget, you can use either `ethernet_ecm` or `ethernet_rndis`. (Simply using `ethernet` will default to `ethernet_ecm`). To make a long story short, use `ethernet_rndis` for Windows and `ethernet_ecm` for everything else. See [ETHERNET.md](ETHERNET.md) for more information.

### HID Payload
To parse the `/boot/usbninja/hid.txt` and send the keystrokes contained within it as HID packets, use the `hid_payload` gadget. The format of `hid.txt` is very simple and straightforward - see [HID.md](HID.md) for more information.

### Mass Storage
The mass storage gadget is activated by including `storage` on the `gadget: ` line in `options.txt`. It makes the image in `/lib/usbninja/storage.img` available as a mass storage device to the host. For details on how to change the default image (only 64MB) and more, see [STORAGE.md](STORAGE.md).
