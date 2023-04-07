package log

import (
	"github.com/pi6atv/winterhill-lib/pkg/log"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func Test_hub_Archive(t *testing.T) {
	now := time.Now()
	tests := []struct {
		name       string
		history    []log.Message
		newMessage log.Message
		want       []log.Message
	}{
		{
			name: "happy path: remove one",
			history: []log.Message{
				{Time: now},
				{Time: now.Add(time.Hour * -2)},
				{Time: now},
				{Time: now},
			},
			newMessage: log.Message{Time: now, Call: "new"},
			want: []log.Message{
				{Time: now},
				{Time: now},
				{Time: now},
				{Time: now, Call: "new"},
			},
		},
		{
			name: "happy path: remove none",
			history: []log.Message{
				{Time: now},
				{Time: now},
				{Time: now},
			},
			newMessage: log.Message{Time: now, Call: "new"},
			want: []log.Message{
				{Time: now},
				{Time: now},
				{Time: now},
				{Time: now, Call: "new"},
			},
		},
		{
			name: "happy path: remove all",
			history: []log.Message{
				{Time: now.Add(time.Hour * -2)},
				{Time: now.Add(time.Hour * -2)},
				{Time: now.Add(time.Hour * -2)},
			},
			newMessage: log.Message{Time: now, Call: "new"},
			want: []log.Message{
				{Time: now, Call: "new"},
			},
		},
		{
			name: "happy path: remove multiple",
			history: []log.Message{
				{Time: now.Add(time.Hour * -2)},
				{Time: now.Add(time.Hour * -2)},
				{Time: now},
				{Time: now},
				{Time: now.Add(time.Hour * -2)},
				{Time: now},
			},
			newMessage: log.Message{Time: now, Call: "new"},
			want: []log.Message{
				{Time: now},
				{Time: now},
				{Time: now},
				{Time: now, Call: "new"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &hub{
				History: tt.history,
			}
			h.Archive(tt.newMessage)
			assert.Equal(t, tt.want, h.History)
		})
	}
}
