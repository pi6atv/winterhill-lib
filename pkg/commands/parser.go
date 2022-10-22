package commands

import (
	"github.com/pkg/errors"
	"strconv"
	"strings"
)

// [to@wh],rcv=1,freq=144600,offset=0000000,srate=125,fplug=A

// ParseWhShort parses a winterhill short command
func ParseWhShort(in string, sep string) (*WhShort, error) {
	if sep == "" {
		sep = ","
	}
	if !strings.HasPrefix(in, "[to@wh]") {
		return nil, errors.New("not in winterhill short format")
	}

	result := WhShort{}
	for _, segment := range strings.Split(in, sep) {
		keyValue := strings.SplitN(segment, "=", 2)
		switch strings.ToUpper(keyValue[0]) {
		case "RCV":
			value, err := strconv.ParseInt(keyValue[1], 10, 64)
			if err != nil {
				return nil, errors.Wrapf(err, "parsing RCV into number: '%s'", segment)
			}
			result.Index = value
		case "FREQ":
			value, err := strconv.ParseFloat(keyValue[1], 10)
			if err != nil {
				return nil, errors.Wrapf(err, "parsing FREQ into number: '%s'", segment)
			}
			result.Frequency = value / 1000
		case "OFFSET":
			value, err := strconv.ParseFloat(keyValue[1], 10)
			if err != nil {
				return nil, errors.Wrapf(err, "parsing OFFSET into number: '%s'", segment)
			}
			result.Offset = value
		case "SRATE":
			value, err := strconv.ParseInt(keyValue[1], 10, 64)
			if err != nil {
				return nil, errors.Wrapf(err, "parsing SRATE into number: '%s'", segment)
			}
			result.SymbolRate = value
		case "FPLUG":
			result.Antenna = keyValue[1]
		case "PRG":
			value, err := strconv.ParseInt(keyValue[1], 10, 64)
			if err != nil {
				return nil, errors.Wrapf(err, "parsing PRG into number: '%s'", segment)
			}
			result.Prg = value
		case "VGX":
			result.VGX = keyValue[1]
		case "VGY":
			result.VGY = keyValue[1]
		case "VOLTAGE":
			value, err := strconv.ParseInt(keyValue[1], 10, 64)
			if err != nil {
				return nil, errors.Wrapf(err, "parsing VOLTAGE into number: '%s'", segment)
			}
			result.Voltage = value
		case "22KHZ":
			result.Tone = keyValue[1] == "1"
		}
	}
	return &result, nil
}
