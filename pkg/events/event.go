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
	Index                int64       `json:"index"`                  // $0
	State                string      `json:"state"`                  // $1:  StateType
	LNAGain              interface{} `json:"lna_gain"`               // $2: unused?
	PunctureRate         interface{} `json:"puncture_rate"`          // $3: unused?
	PowerI               interface{} `json:"power_i"`                // $4: unused?
	PowerO               interface{} `json:"power_o"`                // $5: unused?
	CarrierFrequency     float64     `json:"carrier_frequency"`      // $6: in MHz
	ConstellationI       interface{} `json:"constellation_i"`        // $7: unused?
	ConstellationO       interface{} `json:"constellation_o"`        // $8: unused?
	SymbolRate           int64       `json:"symbol_rate"`            // $9: in kS
	ViterbiErrorRate     interface{} `json:"viterbi_error_rate"`     // $10: unused?
	Ber                  interface{} `json:"ber"`                    // $11: unused?
	Mer                  float64     `json:"mer"`                    // $12: in 0.1 dB
	ServiceName          string      `json:"service_name"`           // $13: service name
	ServiceProviderName  string      `json:"service_provider_name"`  // $14: provider name
	TsNullPercentage     float64     `json:"ts_null_percentage"`     // $15: 99.9 * rcv[rx].nullpacketcountrx / (rcv[rx].packetcountrx + 1)
	EsPid                interface{} `json:"es_pid"`                 // $16:unused
	EsType               interface{} `json:"es_type"`                // $17: unused
	ModulationCode       string      `json:"modulation_code"`        // $18: modulation + fec
	FrameType            string      `json:"frame_type"`             // $19: 'L' ?
	Pilots               string      `json:"pilots"`                 // $20: 'N' ?
	ErrorsLDPCCount      interface{} `json:"errors_ldpc_count"`      // $21: unused?
	ErrorsBCHCount       interface{} `json:"errors_bch_count"`       // $22: unused?
	ErrorsBCHUncorrected interface{} `json:"errors_bch_uncorrected"` // $23: unused?
	LNBSupply            interface{} `json:"lnb_supply"`             // $24: unused?
	LNBPolarisationH     interface{} `json:"lnb_polarisation_h"`     // $25: unused? - is referencedÏ
	MultiStream0         interface{} `json:"multi_stream_0"`         // $26: unused?
	MultiStream1         interface{} `json:"multi_stream_1"`         // $27: unused?
	Debug0               interface{} `json:"debug_0"`                // $28: unused?
	Debug1               interface{} `json:"debug_1"`                // $29: unused?
	DNumber              float64     `json:"d_number"`               // $30: rcv[rx].rawinfos[STATUS_MER] - modinfo_S2[tempu].minmer
	VideoType            string      `json:"video_type"`             // $31:
	RollOff              int64       `json:"roll_off"`               // $32:
	Antenna              string      `json:"antenna"`                // $33:
	AudioType            string      `json:"audio_type"`             // $34:
	TitleBar             string      `json:"title_bar"`              // $94:
	VlcStops             int64       `json:"vlc_stops"`              // $96:
	VlcExts              int64       `json:"vlc_exts"`               // $97:
	ModeChanges          int64       `json:"mode_changes"`           // $98:
	IPChanges            int64       `json:"ip_changes"`             // $99:
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
		case "$12":
			value, err := strconv.ParseFloat(parts[1], 64)
			if err != nil {
				return nil, errors.Wrapf(err, "parsing MER: '%s'", parts[1])
			}
			result.Mer = value
		case "$13":
			result.ServiceName = parts[1]
		case "$14":
			result.ServiceProviderName = parts[1]
		case "$15":
			value, err := strconv.ParseFloat(parts[1], 64)
			if err != nil {
				return nil, errors.Wrapf(err, "parsing TS null percentage: '%s'", parts[1])
			}
			result.TsNullPercentage = value
		case "$18":
			result.ModulationCode = parts[1]
		case "$19":
			result.FrameType = parts[1]
		case "$20":
			result.Pilots = parts[1]
		case "$30":
			value, err := strconv.ParseFloat(parts[1], 64)
			if err != nil {
				return nil, errors.Wrapf(err, "parsing D number: '%s'", parts[1])
			}
			result.DNumber = value
		case "$31":
			result.VideoType = parts[1]
		case "$32":
			value, err := strconv.ParseInt(parts[1], 10, 64)
			if err != nil {
				return nil, errors.Wrapf(err, "parsing roll off: '%s'", parts[1])
			}
			result.RollOff = value
		case "$33":
			result.Antenna = parts[1]
		case "$34":
			result.AudioType = parts[1]
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
