package config

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestNew(t *testing.T) {
	tests := []struct {
		name              string
		configFileContent string
		want              WinterhillConfig
		wantErr           bool
	}{
		{
			name: "happy path",
			configFileContent: `COMMAND = [to@wh],rcv=1,freq=144600,offset=0000000,srate=125,fplug=A                                            
COMMAND = [to@wh],rcv=3,freq=1963260,offset=0000000,srate=500,fplug=A
COMMAND = 3             # a single digit causes a delay of that many seconds
COMMAND = [to@wh],rcv=2,freq=1964419,offset=0000000,srate=500,fplug=B
COMMAND = [to@wh],rcv=4,freq=1245000,offset=0000000,srate=250,fplug=B
`,
			want: WinterhillConfig{
				Receivers: map[int64]ReceiverConfig{
					0: {
						Index:      1,
						Frequency:  144.6,
						Offset:     0,
						SymbolRate: 125,
						Antenna:    "A",
					},
					1: {
						Index:      2,
						Frequency:  1964.419,
						Offset:     0,
						SymbolRate: 500,
						Antenna:    "B",
					},
					2: {
						Index:      3,
						Frequency:  1963.260,
						Offset:     0,
						SymbolRate: 500,
						Antenna:    "A",
					},
					3: {
						Index:      4,
						Frequency:  1245,
						Offset:     0,
						SymbolRate: 250,
						Antenna:    "B",
					},
				},
			},
			wantErr: false,
		},
		{
			name: "happy path - comments",
			configFileContent: `COMMAND = [to@wh],rcv=1,freq=144600,offset=0000000,srate=125,fplug=A                                            
COMMAND = [to@wh],rcv=3,freq=1963260,offset=0000000,srate=500,fplug=A
COMMAND = 3             # a single digit causes a delay of that many seconds
COMMAND = [to@wh],rcv=2,freq=1964419,offset=0000000,srate=500,fplug=B
COMMAND = [to@wh],rcv=4,freq=1245000,offset=0000000,srate=250,fplug=B
#COMMAND = [to@wh],rcv=1,freq=10491500,offset=9750000,srate=1500,fplug=A,vgx=hi         # (hi,lo,hit,lot,off)                                   
#COMMAND = [to@wh],rcv=3,freq=10491500,offset=9750000,srate=1500,fplug=A,vgx=hi
#COMMAND = 3            # a single digit causes a delay of that many seconds
#COMMAND = [to@wh],rcv=2,freq=10499250,offset=9750000,srate=333,fplug=A,vgx=hi
#COMMAND = [to@wh],rcv=4,freq=10498750,offset=9750000,srate=333,fplug=A,vgx=hi
`,
			want: WinterhillConfig{
				Receivers: map[int64]ReceiverConfig{
					0: {
						Index:      1,
						Frequency:  144.6,
						Offset:     0,
						SymbolRate: 125,
						Antenna:    "A",
					},
					1: {
						Index:      2,
						Frequency:  1964.419,
						Offset:     0,
						SymbolRate: 500,
						Antenna:    "B",
					},
					2: {
						Index:      3,
						Frequency:  1963.260,
						Offset:     0,
						SymbolRate: 500,
						Antenna:    "A",
					},
					3: {
						Index:      4,
						Frequency:  1245,
						Offset:     0,
						SymbolRate: 250,
						Antenna:    "B",
					},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			file, err := os.CreateTemp("", "winterhill.ini")
			if err != nil {
				t.Fatal(err)
			}
			defer os.Remove(file.Name())

			_, err = file.Write([]byte(tt.configFileContent))
			if err != nil {
				t.Fatal(err)
			}

			got, err := New(file.Name())
			if (err != nil) != tt.wantErr {
				t.Errorf("New() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if got == nil {
				t.Fatal("no struct returned: nil")
			}
			assert.Equal(t, tt.want, *got)
		})
	}
}
