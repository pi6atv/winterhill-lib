package config

import (
	"github.com/pi6atv/winterhill-lib/pkg/commands"
	"github.com/pkg/errors"
	"os"
	"strings"
)

// WinterhillConfig wraps the startup config as provided by winterhill.ini
type WinterhillConfig struct {
	Receivers map[int64]ReceiverConfig `json:"receivers"`
}

// ReceiverConfig holds the startup config for a single receiver
type ReceiverConfig struct {
	Index      int64   `json:"index"`
	Frequency  float64 `json:"frequency"`
	Offset     float64 `json:"offset"`
	SymbolRate int64   `json:"symbol_rate"`
	Antenna    string  `json:"antenna"`
}

// New parses the winterhill.ini file. Default path is ~pi/winterhill/winterhill.ini
func New(path string) (*WinterhillConfig, error) {
	content, err := os.ReadFile(path)
	if err != nil {
		return nil, errors.Wrapf(err, "reading config file '%s'", path)
	}

	result := WinterhillConfig{
		Receivers: make(map[int64]ReceiverConfig, 4),
	}

	for _, line := range strings.Split(string(content), "\n") {
		if strings.HasPrefix(line, "COMMAND = [to@wh]") {
			parts := strings.SplitN(strings.TrimRight(line, " "), "COMMAND = ", 2)
			command, err := commands.ParseWhShort(parts[1], ",")
			if err != nil {
				return nil, errors.Wrapf(err, "failed to parse: '%s", parts[1])
			}
			result.Receivers[command.Index-1] = ReceiverConfig{
				Index:      command.Index,
				Frequency:  command.Frequency,
				Offset:     command.Offset,
				SymbolRate: command.SymbolRate,
				Antenna:    command.Antenna,
			}
		}
	}
	return &result, nil
}
