package commands

import (
	"fmt"
	"github.com/pkg/errors"
	"net"
)

// WhShort wraps a Winterhill short command
type WhShort struct {
	Index      int64  `json:"index"`       // RCV
	Frequency  int64  `json:"frequency"`   // FREQ
	Offset     int64  `json:"offset"`      // OFFSET
	SymbolRate int64  `json:"symbol_rate"` // SRATE
	Antenna    string `json:"antenna"`     // FPLUG
	Prg        int64  `json:"prg"`         // PRG
	VGX        string `json:"vgx"`         // VGX
	VGY        string `json:"vgy"`         // VGY
	Voltage    int64  `json:"voltage"`     // VOLTAGE
	Tone       bool   `json:"22khz"`       // 22KHZ
}

// [to@wh],rcv=1,freq=144600,offset=0000000,srate=125,fplug=A
func (W WhShort) String() string {
	result := fmt.Sprintf("[to@wh] rcv=%d freq=%d offset=%d srate=%d fplug=%s voltage=%d",
		W.Index, W.Frequency, W.Offset, W.SymbolRate, W.Antenna, W.Voltage,
	)
	if W.Prg != 0 {
		result = fmt.Sprintf("%s prg=%d", result, W.Prg)
	}
	if W.VGX != "" {
		result = fmt.Sprintf("%s vgx=%s", result, W.VGX)
	}
	if W.VGY != "" {
		result = fmt.Sprintf("%s vgy=%s", result, W.VGY)
	}
	if W.Tone {
		result = fmt.Sprintf("%s 22khz=ON", result)
	}
	return result
}

// Send will send the command to remote
func (W WhShort) Send(remote string) error {
	conn, err := net.Dial("udp", remote)
	if err != nil {
		return errors.Wrapf(err, "connecting to %s", remote)
	}

	_, err = conn.Write([]byte(W.String()))
	return errors.Wrapf(err, "sending packet to %s", remote)
}
