package log

import (
	"context"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
	"time"
)

func TestStream_Log(t *testing.T) {
	type args struct {
		module string
		what   string
		r      *http.Request
	}
	tests := []struct {
		name string
		args args
		user string
		want Message
	}{
		{
			name: "happy path",
			args: args{
				module: "module",
				what:   "what",
			},
			user: "user",
			want: Message{
				Call:    "user",
				Setting: "module",
				Value:   "what",
			},
		},
		{
			name: "error path: no user in context",
			args: args{
				module: "module",
				what:   "what",
			},
			user: "",
			want: Message{
				Call:    "???",
				Setting: "module",
				Value:   "what",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			S := &Stream{
				in: make(chan Message, 1),
			}
			r, _ := http.NewRequest("GET", "/", nil)
			if tt.user != "" {
				ctx := context.WithValue(r.Context(), "user", tt.user)
				r = r.WithContext(ctx)
			}
			S.Log(r, tt.args.module, tt.args.what)

			msg := <-S.in
			assert.Equal(t, tt.want.Call, msg.Call, "call")
			assert.Equal(t, tt.want.Setting, msg.Setting, "module")
			assert.Equal(t, tt.want.Value, msg.Value, "action")
			assert.True(t, time.Now().Sub(msg.Time).Seconds() < 1, "time is set")
		})
	}
}

func TestStream_Log_timeout(t *testing.T) {
	S := Stream{in: make(chan Message, 0)} // blocks on write
	start := time.Now()
	timeout := time.After(1 * time.Second)
	done := make(chan bool)

	go func() {
		S.Log(&http.Request{}, "", "")
		done <- true
	}()
	select {
	case <-timeout:
		t.Fatal("Test didn't finish in time")
	case <-done:
	}

	assert.Equal(t, int64(100), (time.Now().Sub(start).Milliseconds()/5)*5)
}
