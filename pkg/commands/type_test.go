package commands

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestWhShort_String(t *testing.T) {
	type fields struct {
		Index      int64
		Frequency  int64
		Offset     int64
		SymbolRate int64
		Antenna    string
		Prg        int64
		VGX        string
		VGY        string
		Voltage    int64
		Tone       bool
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "happy path - minimal",
			fields: fields{
				Index:      1,
				Frequency:  2,
				Offset:     3,
				SymbolRate: 4,
				Antenna:    "A",
				Voltage:    6,
			},
			want: "",
		},
		{
			name: "happy path - all fields",
			fields: fields{
				Index:      1,
				Frequency:  2,
				Offset:     3,
				SymbolRate: 4,
				Antenna:    "A",
				Prg:        5,
				VGX:        "B",
				VGY:        "C",
				Voltage:    6,
				Tone:       true,
			},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			C := WhShort{
				Index:      tt.fields.Index,
				Frequency:  tt.fields.Frequency,
				Offset:     tt.fields.Offset,
				SymbolRate: tt.fields.SymbolRate,
				Antenna:    tt.fields.Antenna,
				Prg:        tt.fields.Prg,
				VGX:        tt.fields.VGX,
				VGY:        tt.fields.VGY,
				Voltage:    tt.fields.Voltage,
				Tone:       tt.fields.Tone,
			}
			assert.Equalf(t, tt.want, C.String(), "String()")
		})
	}
}
