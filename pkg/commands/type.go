package commands

import "fmt"

// WhShort wraps a Winterhill short command
type WhShort struct {
	Index      int64  `json:"index"`       // RCV
	Frequency  int64  `json:"frequency"`   // FREQ
	Offset     int64  `json:"offset"`      // OFFSET
	SymbolRate int64  `json:"symbol_rate"` // SRATE
	Antenna    string `json:"antenna"`     // FPLUG
	Prg        int64  // PRG
	VGX        string // VGX
	VGY        string // VGY
	Voltage    int64  // VOLTAGE
	Tone       bool   // 22KHZ
}

// [to@wh],rcv=1,freq=144600,offset=0000000,srate=125,fplug=A
func (C WhShort) String() string {
	result := fmt.Sprintf("[to@wh] rcv=%d freq=%d offset=%d srate=%d antenna=%s",
		C.Index, C.Frequency, C.Offset, C.SymbolRate, C.Antenna,
	)
	return result
}
