package events

import (
	"github.com/pi6atv/winterhill-lib/pkg/fixtures"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParseEvent(t *testing.T) {
	type args struct {
		in string
	}
	tests := []struct {
		name    string
		args    args
		want    *StatusEvent
		wantErr bool
	}{
		{
			name: "single cycle, rx 1",
			args: args{in: fixtures.FullCycle[1]},
			want: &StatusEvent{
				Index:            1,
				State:            "search",
				CarrierFrequency: 10491.525,
				SymbolRate:       1500,
				TsNullPercentage: 0.0,
				Antenna:          "TOP",
				TitleBar:         "1: *search* SR1500 10491.525MHz T",
				VlcStops:         1,
				VlcExts:          0,
				IPChanges:        2,
			},
			wantErr: false,
		},
		{
			name: "single cycle, rx 3",
			args: args{in: fixtures.FullCycle[3]},
			want: &StatusEvent{
				Index:               3,
				State:               "DVB-S2",
				CarrierFrequency:    1963.332,
				SymbolRate:          500,
				TsNullPercentage:    0.0,
				Antenna:             "TOP",
				TitleBar:            "3: Digital M11.9 D7.9 S2H4 QP3/4 500 3.332M T",
				VlcStops:            1,
				VlcExts:             1,
				IPChanges:           3,
				Mer:                 11.9,
				ServiceName:         "Digital TV",
				ServiceProviderName: "PLUTO DVBS-(2) ",
				ModulationCode:      "QPSK 3/4",
				FrameType:           "L",
				Pilots:              "N",
				DNumber:             7.9,
				VideoType:           "H264",
				RollOff:             25,
				AudioType:           "MPA",
			},
			wantErr: false,
		},
		{
			name: "single cycle, rx 4",
			args: args{in: fixtures.FullCycle[4]},
			want: &StatusEvent{
				Index:               4,
				State:               "DVB-S2",
				CarrierFrequency:    1245.043,
				SymbolRate:          250,
				TsNullPercentage:    0.0,
				Antenna:             "BOT",
				TitleBar:            "4: SERVICE M11.7 D7.7 S2H4 QP3/4 250 5.043M B",
				VlcStops:            2,
				VlcExts:             2,
				IPChanges:           5,
				Mer:                 11.7,
				ServiceName:         "SERVICE",
				ServiceProviderName: "PROVIDER",
				ModulationCode:      "QPSK 3/4",
				FrameType:           "L",
				Pilots:              "N",
				DNumber:             7.7,
				VideoType:           "H264",
				RollOff:             25,
				AudioType:           "MPA",
			},
			wantErr: false,
		},
		{
			name:    "single cycle, rx 0", // winterhill common
			args:    args{in: fixtures.FullCycle[0]},
			want:    nil,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseEvent(tt.args.in)
			t.Logf("err=%v, got=%+v", err, got)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseEvent() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if tt.want == nil {
				if got != nil {
					t.Errorf("ParseEvent() got = %+v, want %+v", *got, tt.want)
				}
			} else {
				assert.Equal(t, tt.want, got)
			}

		})
	}
}
