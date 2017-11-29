# USB Ninja
## Newly Added Features

As things get crossed off the [TODO](TODO.md) list, they will be written about here.

### Read-Only Storage Mode `Version: TBA`
A feature that I'm very excited to be releasing is being able to present a Read-Only Mass Storage device to a host via the USB Ninja. As usual, you can do this by adding a new line to `options.txt`.

```
.
.
.
storage_ro: yes
```

If you set this straight away, then you just won't be able to add anything to the mass storage. The idea is that you should leave the device in read/write mode (by simply omitting a `storage_ro` line in `options.txt`) and place a payload of some kind on the drive. You can then set the read-only flag in `options.txt` so ensure that an AV solution present on the target cannot delete or quarantine your payload if it is detected.

### Post-Setup Scripts `Version: TBA`

You can now specify a script or executable than will be ran after the gadgets are setup and running.
This done by adding a new line to `options.txt`.

```
.
.
.
post: /path/to/file
```

You can either pass it an absolute path as above, or a relative path from within `/boot/usbninja/`.

Note that you need to have the executable bit set on whatever you point to, else it will fail!
You can do this with a simple `chmod` as below:

```
$ sudo chmod +x /boot/usbninja/exec.sh
```

### Alternative Backing File in `storage` mode `Vesion: TBA`

By default, the backing file used in `storage` mode is the one located at `/lib/usbninja/storage.img`. However, you can create another one, as detailed in [STORAGE.img](src/STORAGE.img). Once done, you can specify the location of your new backing file in `options.txt` with the `storage: ` line.
```
.
.
.
storage: /home/alarm/backing_file.img
```
