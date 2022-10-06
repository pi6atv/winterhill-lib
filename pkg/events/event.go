// Package events describes the events coming from the Winterhill on port 9901
// For definitions in Winterhill, see:  https://github.com/BritishAmateurTelevisionClub/winterhill/blob/main/whsource-3v20/whmain-3v20/main.h
package events

import (
	"github.com/pkg/errors"
	"strconv"
	"strings"
)

//type StateType int64
//
//const (
//	Unknown           StateType = 1
//	StatusSearch                = 0
//	StatusFoundHeader           = 1
//	StatusDemodS2               = 2
//	StatusDemodS                = 3
//	StatusLost                  = 0x80
//	StatusTimeout               = 0x81
//	StatusIdle                  = 0x82
//)

type StatusEvent struct {
	Index                int64
	State                string      // StateType
	LNAGain              interface{} // unused?
	PunctureRate         interface{} // unused?
	PowerI               interface{} // unused?
	PowerO               interface{} // unused?
	CarrierFrequency     float64     // in MHz
	ConstellationI       interface{} // unused?
	ConstellationO       interface{} // unused?
	SymbolRate           int64       // in kS
	ViterbiErrorRate     interface{} // unused?
	Ber                  interface{} // unused?
	Mer                  float64     // in 0.1 dB
	ServiceName          string
	ServiceProviderName  string
	TsNullPercentage     float64     // 99.9 * rcv[rx].nullpacketcountrx / (rcv[rx].packetcountrx + 1)
	EsPid                interface{} // unused
	EsType               interface{} // unused
	ModulationCode       int64       // fixme: MAP? https://github.com/BritishAmateurTelevisionClub/winterhill/blob/3e88b5cf9817f69ca69fa90fb5ef1864a1305a21/whsource-3v20/whmain-3v20/main.c#L322
	FrameType            bool
	Pilots               bool
	ErrorsLDPCCount      interface{} // unused?
	ErrorsBCHCount       interface{} // unused?
	ErrorsBCHUncorrected interface{} // unused?
	LNBSupply            interface{} // unused?
	LNBPolarisationH     interface{} // unused? - is referencedÏ
	MultiStream0         interface{} // unused?
	MultiStream1         interface{} // unused?
	Debug0               interface{} // unused?
	Debug1               interface{} // unused?
	DNumber              float64     // rcv[rx].rawinfos[STATUS_MER] - modinfo_S2[tempu].minmer
	VideoType            string
	RollOff              int64
	Antenna              string
	AudioType            string
	TitleBar             string
	VlcStops             int64
	VlcExts              int64
	ModeChanges          int64
	IPChanges            int64
}

func ParseEvent(in string) (*StatusEvent, error) {
	result := StatusEvent{}
	for _, line := range strings.Split(in, "\r\n") {
		parts := strings.SplitN(line, ",", 2)
		switch parts[0] {
		case "$0":
			if parts[1] == "0" {
				return nil, nil // winterhill common
			}
			value, err := strconv.ParseInt(parts[1], 10, 64)
			if err != nil {
				return nil, errors.Wrapf(err, "parsing index: '%s'", parts[1])
			}
			result.Index = value
		case "$1":
			result.State = parts[1]
		case "$6":
			value, err := strconv.ParseFloat(parts[1], 64)
			if err != nil {
				return nil, errors.Wrapf(err, "parsing frequency: '%s'", parts[1])
			}
			result.CarrierFrequency = value
		case "$9":
			value, err := strconv.ParseInt(parts[1], 10, 64)
			if err != nil {
				return nil, errors.Wrapf(err, "parsing symbol rate: '%s'", parts[1])
			}
			result.SymbolRate = value
		case "$15":
			value, err := strconv.ParseFloat(parts[1], 64)
			if err != nil {
				return nil, errors.Wrapf(err, "parsing TS null percentage: '%s'", parts[1])
			}
			result.TsNullPercentage = value
		case "$33":
			result.Antenna = parts[1]
		case "$94":
			result.TitleBar = parts[1]
		case "$96":
			value, err := strconv.ParseInt(parts[1], 10, 64)
			if err != nil {
				return nil, errors.Wrapf(err, "parsing VLC stops: '%s'", parts[1])
			}
			result.VlcStops = value
		case "$97":
			value, err := strconv.ParseInt(parts[1], 10, 64)
			if err != nil {
				return nil, errors.Wrapf(err, "parsing VLC exits: '%s'", parts[1])
			}
			result.VlcExts = value
		case "$98":
			value, err := strconv.ParseInt(parts[1], 10, 64)
			if err != nil {
				return nil, errors.Wrapf(err, "parsing IP changes: '%s'", parts[1])
			}
			result.IPChanges = value
		}
	}

	return &result, nil
}
