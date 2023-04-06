package log

import (
	"github.com/gorilla/websocket"
	"github.com/pi6atv/winterhill-lib/pkg/log"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"time"
)

var (
	NextID              int
	clientSendSumMetric = promauto.NewCounter(
		prometheus.CounterOpts{
			Name: "winterhill_web_log_client_send_time_total",
			Help: "winterhill web log client send time total seconds",
		})
	clientSendCountMetric = promauto.NewCounter(
		prometheus.CounterOpts{
			Name: "winterhill_web_log_client_send_time_count",
			Help: "winterhill web log client send time count",
		})
)

const (
	// Time allowed to write a message to the peer.
	writeWait = 10 * time.Second
)

type Client struct {
	id   int
	conn *websocket.Conn
	in   chan log.Message
	hub  *hub
}

func NewClient(conn *websocket.Conn, hub *hub) *Client {
	NextID++
	return &Client{
		id:   NextID,
		conn: conn,
		in:   make(chan log.Message, 1000),
		hub:  hub,
	}
}

func (c *Client) handle() {
	// send history
	for _, message := range c.hub.History {
		_ = c.conn.SetWriteDeadline(time.Now().Add(writeWait))
		err := c.conn.WriteJSON(message)
		if err != nil {
			c.hub.deregister <- c
			return
		}
	}

	// handle incoming messages
	ticker := time.NewTicker(writeWait * 2)
	for {
		select {
		// a ping message to detect left clients early
		case <-ticker.C:
			err := c.conn.WriteControl(websocket.PingMessage, []byte{}, time.Now().Add(writeWait))
			if err != nil {
				c.hub.deregister <- c
				return
			}

		case message := <-c.in:
			start := time.Now()
			_ = c.conn.SetWriteDeadline(time.Now().Add(writeWait))
			err := c.conn.WriteJSON(message)
			if err != nil {
				c.hub.deregister <- c
				return
			}
			clientSendSumMetric.Add(float64(time.Now().Sub(start).Microseconds()) / 1000000)
			clientSendCountMetric.Inc()
		}
	}
}
