// Package events describes the events coming from the Winterhill on port 9901
// For definitions in Winterhill, see:  https://github.com/BritishAmateurTelevisionClub/winterhill/blob/main/whsource-3v20/whmain-3v20/main.h
package events

import (
	"fmt"
	"github.com/libp2p/go-reuseport"
	"github.com/pi6atv/winterhill-lib/pkg/ringbuffer"
	"github.com/pkg/errors"
	"strconv"
	"strings"
	"sync"
	"time"
)

type floatHistory struct {
	Time  time.Time `json:"time"`
	Value float64   `json:"value"`
}

type StatusEvent struct {
	lock                 sync.Mutex
	Index                int64            `json:"index"`              // $0
	State                string           `json:"state"`              // $1: StateType
	LNAGain              interface{}      `json:"lna_gain"`           // $2: unused?
	PunctureRate         interface{}      `json:"puncture_rate"`      // $3: unused?
	PowerI               interface{}      `json:"power_i"`            // $4: unused?
	PowerO               interface{}      `json:"power_o"`            // $5: unused?
	CarrierFrequency     float64          `json:"carrier_frequency"`  // $6: in MHz
	ConstellationI       interface{}      `json:"constellation_i"`    // $7: unused?
	ConstellationO       interface{}      `json:"constellation_o"`    // $8: unused?
	SymbolRate           int64            `json:"symbol_rate"`        // $9: in kS
	ViterbiErrorRate     interface{}      `json:"viterbi_error_rate"` // $10: unused?
	Ber                  interface{}      `json:"ber"`                // $11: unused?
	Mer                  float64          `json:"mer"`                // $12: in 0.1 dB
	MerHistory           *ringbuffer.Ring `json:"mer_history"`
	ServiceName          string           `json:"service_name"`           // $13: service name
	ServiceProviderName  string           `json:"service_provider_name"`  // $14: provider name
	TsNullPercentage     float64          `json:"ts_null_percentage"`     // $15: 99.9 * rcv[rx].nullpacketcountrx / (rcv[rx].packetcountrx + 1)
	EsPid                interface{}      `json:"es_pid"`                 // $16: unused
	EsType               interface{}      `json:"es_type"`                // $17: unused
	ModulationCode       string           `json:"modulation_code"`        // $18: modulation + fec
	FrameType            string           `json:"frame_type"`             // $19: 'L' ?
	Pilots               string           `json:"pilots"`                 // $20: 'N' ?
	ErrorsLDPCCount      interface{}      `json:"errors_ldpc_count"`      // $21: unused?
	ErrorsBCHCount       interface{}      `json:"errors_bch_count"`       // $22: unused?
	ErrorsBCHUncorrected interface{}      `json:"errors_bch_uncorrected"` // $23: unused?
	LNBSupply            interface{}      `json:"lnb_supply"`             // $24: unused?
	LNBPolarisationH     interface{}      `json:"lnb_polarisation_h"`     // $25: unused? - is referencedÏ
	MultiStream0         interface{}      `json:"multi_stream_0"`         // $26: unused?
	MultiStream1         interface{}      `json:"multi_stream_1"`         // $27: unused?
	Debug0               interface{}      `json:"debug_0"`                // $28: unused?
	Debug1               interface{}      `json:"debug_1"`                // $29: unused?
	DNumber              float64          `json:"d_number"`               // $30: MER value above minimum for decoding
	VideoType            string           `json:"video_type"`             // $31:
	RollOff              int64            `json:"roll_off"`               // $32:
	Antenna              string           `json:"antenna"`                // $33:
	AudioType            string           `json:"audio_type"`             // $34:
	TitleBar             string           `json:"title_bar"`              // $94:
	VlcStops             int64            `json:"vlc_stops"`              // $96:
	VlcExts              int64            `json:"vlc_exts"`               // $97:
	ModeChanges          int64            `json:"mode_changes"`           // $98:
	IPChanges            int64            `json:"ip_changes"`             // $99:
}

func StateToFloat(in string) float64 {
	switch in {
	case "idle":
		return 82
	case "timeout":
		return 81
	case "lost":
		return 80
	case "search":
		return 0
	case "header":
		return 1
	case "DVB-S2":
		return 2
	case "DVB-S":
		return 3

	default:
		return -1
	}
}

type Listener struct {
	Receivers map[int64]*StatusEvent `json:"receivers"`
}

func New(n int64, historySize int) *Listener {
	l := Listener{
		Receivers: make(map[int64]*StatusEvent, n),
	}
	for i := int64(0); i < n; i++ {
		l.Receivers[i] = &StatusEvent{
			MerHistory: ringbuffer.New(historySize),
		}
	}
	return &l
}

// Run will start listening on the network for events, and use them to update Receivers
func (L *Listener) Run(addr string, stop chan bool) error {
	const bufLen = 1500
	var buf [bufLen]byte

	sock, err := reuseport.ListenPacket("udp", addr)
	if err != nil {
		return errors.Wrapf(err, "listen on '%s'", addr)
	}
	for {
		select {
		case <-stop:
			return nil
		default:
			_ = sock.SetReadDeadline(time.Now().Add(time.Second))
			size, _, err := sock.ReadFrom(buf[:])
			if err != nil || size == 0 {
				time.Sleep(time.Millisecond * 10) // 'rate limit'
				continue                          // fixme, handle error
			}

			if strings.HasPrefix(string(buf[0:4]), "$0,0") {
				//fmt.Println("skipping winterhill global config")
				continue
			}
			go L.parse(string(buf[0:size]))
		}
	}
}

// parse will parse a UDP message and update the correct receiver
func (L *Listener) parse(in string) error {
	events := strings.Split(in, "\r\n")

	if !strings.HasPrefix(events[0], "$0,") {
		return errors.New("missing index header")
	}

	parts := strings.SplitN(events[0], ",", 2)
	value, err := strconv.ParseInt(parts[1], 10, 64)
	if err != nil {
		return errors.Wrapf(err, "parsing index: '%s'", parts[1])
	}

	index := value - 1 // 0 based
	//fmt.Printf("parse: RECEIVER: %d\n", value)
	L.Receivers[index].lock.Lock() // make sure we're the only one writing
	defer L.Receivers[index].lock.Unlock()
	L.Receivers[index].Index = value // set the human-readable receiver index
	promReceiverUpdate.WithLabelValues(fmt.Sprintf("%d", index+1)).Inc()

	// MER isn't reported when there's no signal. Default it to -100
	// FIXME: find absolute lowest possible MER
	L.Receivers[index].Mer = -100
	L.Receivers[index].DNumber = -100

	// loop over all events
	for _, line := range events[1:] {
		//fmt.Printf("LINE: %s\n", line)
		parts := strings.SplitN(line, ",", 2)
		switch parts[0] {
		case "$1":
			L.Receivers[index].State = parts[1]
			promReceiverState.WithLabelValues(fmt.Sprintf("%d", index+1)).Set(StateToFloat(L.Receivers[index].State))
		case "$6":
			value, err := strconv.ParseFloat(parts[1], 64)
			if err != nil {
				return errors.Wrapf(err, "parsing frequency: '%s'", parts[1])
			}
			L.Receivers[index].CarrierFrequency = value
			promReceiverFreq.WithLabelValues(fmt.Sprintf("%d", index+1)).Set(L.Receivers[index].CarrierFrequency)
		case "$9":
			value, err := strconv.ParseInt(parts[1], 10, 64)
			if err != nil {
				return errors.Wrapf(err, "parsing symbol rate: '%s'", parts[1])
			}
			L.Receivers[index].SymbolRate = value
			promReceiverSR.WithLabelValues(fmt.Sprintf("%d", index+1)).Set(float64(value))
		case "$12":
			value, err := strconv.ParseFloat(parts[1], 64)
			if err != nil {
				return errors.Wrapf(err, "parsing MER: '%s'", parts[1])
			}
			L.Receivers[index].Mer = value
		case "$13":
			L.Receivers[index].ServiceName = parts[1]
		case "$14":
			L.Receivers[index].ServiceProviderName = parts[1]
		case "$15":
			value, err := strconv.ParseFloat(parts[1], 64)
			if err != nil {
				return errors.Wrapf(err, "parsing TS null percentage: '%s'", parts[1])
			}
			L.Receivers[index].TsNullPercentage = value
		case "$18":
			L.Receivers[index].ModulationCode = parts[1]
		case "$19":
			L.Receivers[index].FrameType = parts[1]
		case "$20":
			L.Receivers[index].Pilots = parts[1]
		case "$30":
			value, err := strconv.ParseFloat(parts[1], 64)
			if err != nil {
				return errors.Wrapf(err, "parsing D number: '%s'", parts[1])
			}
			L.Receivers[index].DNumber = value
		case "$31":
			L.Receivers[index].VideoType = parts[1]
		case "$32":
			value, err := strconv.ParseInt(parts[1], 10, 64)
			if err != nil {
				return errors.Wrapf(err, "parsing roll off: '%s'", parts[1])
			}
			L.Receivers[index].RollOff = value
		case "$33":
			L.Receivers[index].Antenna = parts[1]
		case "$34":
			L.Receivers[index].AudioType = parts[1]
		case "$94":
			L.Receivers[index].TitleBar = parts[1]
		case "$96":
			value, err := strconv.ParseInt(parts[1], 10, 64)
			if err != nil {
				return errors.Wrapf(err, "parsing VLC stops: '%s'", parts[1])
			}
			L.Receivers[index].VlcStops = value
		case "$97":
			value, err := strconv.ParseInt(parts[1], 10, 64)
			if err != nil {
				return errors.Wrapf(err, "parsing VLC exits: '%s'", parts[1])
			}
			L.Receivers[index].VlcExts = value
		case "$98":
			value, err := strconv.ParseInt(parts[1], 10, 64)
			if err != nil {
				return errors.Wrapf(err, "parsing IP changes: '%s'", parts[1])
			}
			L.Receivers[index].IPChanges = value
		}
	}

	// because we set a default MER at the start, we always want to store it in the history here
	promReceiverMer.WithLabelValues(fmt.Sprintf("%d", index+1)).Set(L.Receivers[index].Mer)
	L.Receivers[index].MerHistory.Add(floatHistory{Time: time.Now(), Value: L.Receivers[index].Mer})
	promReceiverDnumber.WithLabelValues(fmt.Sprintf("%d", index+1)).Set(L.Receivers[index].DNumber)
	return nil
}
