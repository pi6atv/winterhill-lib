## Commands
[Quicktune](https://github.com/BritishAmateurTelevisionClub/winterhill/blob/main/whsource-3v20/whmain-3v20/main.c#L1317)
```bash
22:30:00.154606 IP 44.137.11.146.55542 > 44.137.11.145.9922: UDP, length 112
E...?1......,...,.....&..x8.[GlobalMsg],Freq=10492745,Offset=0,Doppler=0,Srate=500,WideScan=-,LowSR=-,DVBmode=-,FPlug=B,Voltage=0,22kHz=Off

22:30:00.904094 IP 44.137.11.146.55542 > 44.137.11.145.9923: UDP, length 112
E...?2......,...,.....&..x9.[GlobalMsg],Freq=10492745,Offset=0,Doppler=0,Srate=500,WideScan=-,LowSR=-,DVBmode=-,FPlug=A,Voltage=0,22kHz=Off

22:30:01.537988 IP 44.137.11.146.55542 > 44.137.11.145.9924: UDP, length 112
E...?3......,...,.....&..x8.[GlobalMsg],Freq=10492745,Offset=0,Doppler=0,Srate=500,WideScan=-,LowSR=-,DVBmode=-,FPlug=B,Voltage=0,22kHz=Off
```
Supported args:
* FREQ (integer)
* OFFSET (integer)
* SRATE (integer)
* FPLUG (A|B)
* PRG (integer)
* VGX (OFF|LO|HI|LOT|HIT)
* VGY (OFF|LO|HI|LOT|HIT)
* VOLTAGE (0|13|18)
* 22KHZ (OFF|ON)

[winterhill short](https://github.com/BritishAmateurTelevisionClub/winterhill/blob/main/whsource-3v20/whmain-3v20/main.c#L1294)
```bash
22:30:50.307958 IP 44.137.11.146.56156 > 44.137.11.145.9921: UDP, length 73
E..e?4......,...,....\&..Q..[to@wh] rcv=1 freq=1240000000 srate=1500 offset=9000000 fplug=B vgx=LOT

22:31:23.983410 IP 44.137.11.146.51936 > 44.137.11.145.9921: UDP, length 73
E..e?5......,...,.....&..Q	.[to@wh] rcv=1 freq=1240000000 srate=1500 offset=9000000 fplug=B vgy=OFF
```
Supported args:
* RCV (1..4)
* FREQ (integer)
* OFFSET (integer)
* SRATE (integer)
* FPLUG (A|B)
* PRG (integer)
* VGX (OFF|LO|HI|LOT|HIT)
* VGY (OFF|LO|HI|LOT|HIT)
* VOLTAGE (0|13|18)
* 22KHZ (OFF|ON)

## Status Format (port 9901)
* `$`
* `<index>`
* `,`
* `<value(s)>`
* `<cr><nl>`

End of section: `<nul>`