package log

import (
	"github.com/pi6atv/winterhill-lib/pkg/log"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"time"
)

var (
	clientsMetric = promauto.NewGauge(
		prometheus.GaugeOpts{
			Name: "winterhill_web_log_clients",
			Help: "winterhill web log clients",
		},
	)

	clientStateMetric = promauto.NewCounterVec(
		prometheus.CounterOpts{
			Name: "winterhill_web_log_client_states_total",
		},
		[]string{"action"},
	)

	archiveSizeMetric = promauto.NewGauge(
		prometheus.GaugeOpts{
			Name: "winterhill_web_log_archive_size",
			Help: "winterhill web log archive size",
		},
	)
)

type hub struct {
	logStream  *log.Stream
	History    []log.Message
	clients    map[int]*client
	register   chan *client
	deregister chan *client
}

// NewHub creates a new hub with initialised channels
func NewHub(log *log.Stream) *hub {
	return &hub{
		logStream:  log,
		clients:    make(map[int]*client),
		register:   make(chan *client),
		deregister: make(chan *client),
	}
}

// Start will run the hub that handles the broadcasting of log messages
func (h *hub) Start() {
	go func() {
		ticker := time.NewTicker(5 * time.Second)
		for {
			select {
			case <-ticker.C:
				h.logStream.Out <- log.Message{Setting: "ping"} // should filter itself out, as it happens on epoch
			}
		}
	}()

	for {
		select {
		case message := <-h.logStream.Out:
			for _, client := range h.clients {
				client.in <- message
			}

			h.Archive(message)
		case client := <-h.register:
			clientStateMetric.WithLabelValues("register").Inc()
			h.clients[client.id] = client
			clientsMetric.Set(float64(len(h.clients)))
		case client := <-h.deregister:
			clientStateMetric.WithLabelValues("deregister").Inc()
			delete(h.clients, client.id)
			_ = client.conn.Close()
			clientsMetric.Set(float64(len(h.clients)))
		}
	}
}

// Archive will store the message in the history, and remove messages older than 4 hours
func (h *hub) Archive(message log.Message) {
	// store in memory
	h.History = append(h.History, message)

	// remove older messages
	var hNew []log.Message
	for index, item := range h.History {
		if item.Time.After(time.Now().Add(-4 * time.Hour)) {
			hNew = append(hNew, h.History[index])
		}
	}
	h.History = hNew
	archiveSizeMetric.Set(float64(len(h.History)))
}
