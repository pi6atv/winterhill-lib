package events

import (
	"github.com/pi6atv/winterhill-lib/pkg/fixtures"
	"reflect"
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
			} else if !reflect.DeepEqual(*got, *tt.want) {
				t.Errorf("ParseEvent() got = %+v, want %+v", *got, *tt.want)
			}
		})
	}
}
