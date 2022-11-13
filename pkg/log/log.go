package log

import (
	"fmt"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/sirupsen/logrus"
	"net/http"
	"time"
)

var (
	messageMetric = promauto.NewCounterVec(
		prometheus.CounterOpts{
			Name: "winterhill_log_message",
			Help: "winterhill log messages",
		},
		[]string{"user", "receiver", "setting", "value"},
	)
	timeoutMetric = promauto.NewCounter(
		prometheus.CounterOpts{
			Name: "winterhill_log_timeout",
			Help: "winterhill log timeout",
		},
	)
)

type Stream struct {
	in  chan Message
	Out chan Message
}

type Message struct {
	Time     time.Time `json:"time"`
	Call     string    `json:"user"`
	Receiver int64     `json:"receiver"`
	Setting  string    `json:"setting"`
	Value    string    `json:"value"`
}

func New() *Stream {
	stream := Stream{
		in:  make(chan Message),
		Out: make(chan Message),
	}
	go func() {
		for {
			select {
			case msg := <-stream.in:
				stream.Out <- msg
			}
		}
	}()
	return &stream
}

func (S *Stream) Log(r *http.Request, receiver int64, setting, value string) {
	if S == nil {
		return // used during tests
	}

	user, ok := r.Context().Value("user").(string)
	if !ok {
		logrus.Warn("failed to get user from request context")
		user = "???"
	}

	msg := Message{
		Time:     time.Now(),
		Call:     user,
		Receiver: receiver,
		Setting:  setting,
		Value:    value,
	}
	messageMetric.WithLabelValues(user, fmt.Sprintf("%d", receiver), setting, value).Inc()

	// block with a timeout, while writing to the channel
	select {
	case <-time.After(100 * time.Millisecond):
		timeoutMetric.Inc()
		break
	case S.in <- msg:
		break
	}
}
