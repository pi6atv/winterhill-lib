package ringbuffer

import (
	"container/ring"
	"github.com/stretchr/testify/assert"
	"math"
	"testing"
)

func TestRing_MarshalJSON(t *testing.T) {
	type fields struct {
		ring *ring.Ring
	}
	tests := []struct {
		name    string
		in      []any
		want    []byte
		wantErr bool
	}{
		{
			name:    "happy flow",
			in:      []any{1, "a", 2.3},
			want:    []byte("[1,\"a\",2.3]"),
			wantErr: false,
		},
		{
			name:    "error flow",
			in:      []any{math.NaN(), "a", 2.3},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			R := &Ring{
				ring: ring.New(len(tt.in)),
			}
			for _, item := range tt.in {
				R.ring.Value = item
				R.ring = R.ring.Next()
			}
			got, err := R.MarshalJSON()
			if (err != nil) != tt.wantErr {
				t.Errorf("MarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			assert.Equal(t, string(tt.want), string(got))
		})
	}
}

func TestRing_Add(t *testing.T) {
	a := Ring{ring: ring.New(2)}
	a.Add(1)
	assert.Equal(t, 2, a.ring.Len())
	assert.Equal(t, 1, a.ring.Prev().Value)
}
