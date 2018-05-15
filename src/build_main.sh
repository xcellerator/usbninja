#!/bin/sh

go build main.go ethernet.go extra.go hid_payload.go hid_remote.go parser.go serial.go storage.go generic_gadget.go keyboard.go
#cp main /lib/usbninja/
