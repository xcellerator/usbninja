# USB Ninja
## HID Gadget

There is currently only one HID gadget available: `hid_payload`, although others are in the works. It emulates a keyboard in a similar way to the [USB Rubber Ducky](https://hakshop.com/products/usb-rubber-ducky-deluxe). The main file of importance is `/boot/usbninja/hid.txt` (accessible on partition one of the MicroSD card). The format is very similar to (and indeed heavily influenced by ) Hak5's `DuckyScript`. In fact, any DuckyScript file should also work with the USB Ninja, but some of the more advanced functions (like looping) are yet to be implemented and will simply be ignored.

### Example `hid.txt`
The following will open `notepad.exe` on Windows and type "Hello, World!". As you can see, it is very straightforward.
```
DELAY 5000
GUI r
DELAY 1000
PRINT notepad
ENTER
DELAY 1000
PRINT Hello, World!
```

### Supported Instructives
"Instructives" are the name given to the first word of each line in `hid.txt`. Anything not listed here will simply be ignored, so commenting your payloads becomes trivial!

Below are currently supported instructives:
```
DELAY <int>  # Sleep for <int> many milliseconds.
PRINT <str>  # Type out <str>, see further down for the supported characters.
ENTER        # Press the ENTER key. Aliased by "RETURN".
ESC          # Press the ESC key.
DEL          # Press the DEL key.
TAB          # Press the TAB key.
SPACE        # Press the SPACE bar.
CAPS         # Press the CAPS LOCK button.
F1-F12       # Press one of the function keys F1-F12.
PRNTSCRN     # Press the PRINTSCREEN key.
SCROLLOCK    # Press the SCROLL LOCK key.
PAUSE        # Press the PAUSE key.
INSERT       # Press the INSERT key.
HOME         # Press the HOME key.
PGUP         # Press the PAGE UP key.
PGDWN        # Press the PAGE DOWN key.
END          # Press the END key.
RIGHT        # Press the RIGHT arrow key.
LEFT         # Press the LEFT arrow key.
UP           # Press the UP arrow key.
DOWN         # Press the DOWN arrow key.
CLEAR        # Press the CLEAR key.
VOLUP        # Press the Volume Up button.
VOLDOWN      # Press the Volume Down button.

CTRL <key>   # Press CTRL + <key>
ALT <key>    # Press ALT + <key>
GUI <key>    # Press GUI + <key>
SHIFT <key>  # Press SHIFT + <key>
```

### Some important points about ``CTRL``, ``ALT``, ``GUI`` and ``SHIFT``
The ``CTRL``, ``ALT``, ``GUI`` and ``SHIFT`` keys are known as _modifiers_ and can be coupled together like ``CTRL ALT DEL`` and ``CTRL SHIFT ENTER``. In fact, you could even have all four as in ``CTRL ALT SHIFT GUI a``, if you really wanted to.

**NOTE:** The order of the modifiers is _**irrelevant**_. ``CTRL ALT DEL`` is exactly the same as ``ALT CTRL DEL``. The only requirement is that the _key_ (i.e. not one of the modifiers) is the last word on the line.

**E.g.** ``CTRL ALT DEL`` and ``ALT CTRL DEL`` are fine (and equivalent!) but ``CTRL DEL ALT`` is not!

### Supported characters
When using the `PRINT` instructive, it may be useful to have a list of the currently supported characters.
```
A-Z
a-z
0,1,2,3,4,5,6,7,8,9
! " # $ % ^ & * ( ) _ + - =
[ ] ; ' , . / { } : @ ~ < > ?
| \ ` ¬
```

**NOTE:** Support for multiple keyboard layouts is currently in development. Currently it defaults to `en_US`, so even though I have an `en_GB` keyboard, the key that reads "£" (`SHIFT + 3` usually) actually outputs "#".

### Technical Details
In [extra.go](../src/extra.go), a map `hex_codes` is defined which contains all the mappings for each key (bar a select few). It's here that all the "translation" is done when we parse through `hid.txt` and lookup each character in the string following a `PRINT` instructive. If a line in `hid.txt` doesn't start with the `PRINT` instructive, then the parser will check to see if the leading word exists in `hex_codes`.

Once the `hex_code` is looked up then the `hid_packet` is built as an array of strings. Eventually this HID packet will be echoed out to `/dev/hidg0`.
