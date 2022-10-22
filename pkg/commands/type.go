package commands

// WhShort wraps a Winterhill short command
type WhShort struct {
	Index      int64   `json:"index"`       // RCV
	Frequency  float64 `json:"frequency"`   // FREQ
	Offset     float64 `json:"offset"`      // OFFSET
	SymbolRate int64   `json:"symbol_rate"` // SRATE
	Antenna    string  `json:"antenna"`     // FPLUG
	Prg        int64   // PRG
	VGX        string  // VGX
	VGY        string  // VGY
	Voltage    int64   // VOLTAGE
	Tone       bool    // 22KHZ
}
