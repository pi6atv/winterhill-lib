package commands

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParseWhShort(t *testing.T) {
	type args struct {
		in  string
		sep string
	}
	tests := []struct {
		name    string
		args    args
		want    WhShort
		wantErr bool
	}{
		{
			name: "happy path",
			args: args{in: "[to@wh],rcv=1,freq=144600,offset=0000000,srate=125,fplug=A", sep: ","},
			want: WhShort{
				Index:      1,
				Frequency:  144.6,
				Offset:     0,
				SymbolRate: 125,
				Antenna:    "A",
			},
			wantErr: false,
		},
		{
			name:    "error path: no whshort",
			args:    args{in: "this is no command"},
			wantErr: true,
		},
		{
			name:    "error path: invalid number",
			args:    args{in: "[to@wh],rcv=A,", sep: ","},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseWhShort(tt.args.in, tt.args.sep)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseWhShort() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if err != nil {
				return
			}
			if got == nil && !tt.wantErr {
				t.Fatal("got no struct, but didn't expect error")
				return
			}
			assert.Equal(t, tt.want, *got)
		})
	}
}
