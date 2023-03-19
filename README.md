# Аргонная сварка Юры Бондарева под arduino-nano на tinygo.org (golang).

### для прошивки в arduino nano под windows:
```
tinygo flash -target arduino-nano weld.go
```

### для прошивки в arduino nano под linux/mac:

```
tinygo flash -target arduino-nano -port /dev/cu.wchusbserial1410 weld.go

avrdude: AVR device initialized and ready to accept instructions
Reading | ################################################## | 100% 0.02s
avrdude: Device signature = 0x1e950f (probably m328p)
avrdude: NOTE: "flash" memory has been specified, an erase cycle will be performed
         To disable this feature, specify the -D option.
avrdude: erasing chip
avrdude: reading input file "/var/folders/gd/47xhprb972j4klgykgdv1p300000gp/T/tinygo1144818520/main.hex"
avrdude: writing flash (1808 bytes):
Writing | ################################################## | 100% 1.17s
avrdude: 1808 bytes of flash written
avrdude: verifying flash memory against /var/folders/gd/47xhprb972j4klgykgdv1p300000gp/T/tinygo1144818520/main.hex:
Reading | ################################################## | 100% 1.03s
avrdude: 1808 bytes of flash verified
avrdude done.  Thank you.
```
