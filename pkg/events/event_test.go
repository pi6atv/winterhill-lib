package events

import (
	"github.com/pi6atv/winterhill-lib/pkg/fixtures"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestListener_parse(t *testing.T) {
	type args struct {
		in string
	}
	tests := []struct {
		name    string
		args    args
		want    map[int64]*StatusEvent
		wantErr bool
	}{
		{
			name: "happy flow: rx 1",
			args: args{in: fixtures.FullCycle[1]},
			want: map[int64]*StatusEvent{
				0: {
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
					Mer:              -100,
					DNumber:          -100,
				},
			},
		},
		{
			name: "happy flow: rx 3",
			args: args{in: fixtures.FullCycle[3]},
			want: map[int64]*StatusEvent{
				2: {
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
			},
		},
		{
			name: "happy flow: rx 4",
			args: args{in: fixtures.FullCycle[4]},
			want: map[int64]*StatusEvent{
				3: {
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
			},
			wantErr: false,
		},
		{
			name: "error flow: missing header",
			args: args{in: "foobar"},
			want: map[int64]*StatusEvent{
				0: {},
				1: {},
				2: {},
				3: {},
			},
			wantErr: true,
		},
		{
			name: "error flow: invalid header",
			args: args{in: "$0,a"},
			want: map[int64]*StatusEvent{
				0: {},
				1: {},
				2: {},
				3: {},
			},
			wantErr: true,
		},
		{
			name: "error flow: invalid freq",
			args: args{in: "$0,1\r\n$6,abcd"},
			want: map[int64]*StatusEvent{
				0: {Index: 1, Mer: -100, DNumber: -100},
				1: {},
				2: {},
				3: {},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			L := New(4, 1)
			if err := L.parse(tt.args.in); (err != nil) != tt.wantErr {
				t.Errorf("parse() error = %v, wantErr %v", err, tt.wantErr)
			}
			for index, receiver := range tt.want {
				// we're not testing the history here
				receiver.MerHistory = L.Receivers[index].MerHistory
				receiver.ServiceHistory = L.Receivers[index].ServiceHistory
				assert.Equal(t, tt.want[index], L.Receivers[index])
			}
		})
	}
}
