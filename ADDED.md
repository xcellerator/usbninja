# USB Ninja
## Newly Added Features

As things get crossed off the [TODO](TODO.md) list, they will be written about here.

### Post-Setup Scripts

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

### Alternative Backing File in `storage` mode

By default, the backing file used in `storage` mode is the one located at `/lib/usbninja/storage.img`. However, you can create another one, as detailed in [STORAGE.img](src/STORAGE.img). Once done, you can specify the location of your new backing file in `options.txt` with the `storage: ` line.
```
.
.
.
storage: /home/alarm/backing_file.img
```
