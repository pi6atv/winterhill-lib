package status_api

import (
	"encoding/json"
	"fmt"
	"github.com/pi6atv/winterhill-lib/pkg/config"
	"github.com/pi6atv/winterhill-lib/pkg/events"
	"github.com/pi6atv/winterhill-lib/pkg/summary"
	"github.com/pi6atv/winterhill-lib/pkg/web/metrics"
	"github.com/pkg/errors"
	"net/http"
)

type Api struct {
	eventListener   *events.Listener
	summaryListener *summary.Listener
	config          *config.WinterhillConfig
}

func New(ip string, port int) (*Api, error) {
	iniConfig, err := config.New("")
	if err != nil {
		return nil, errors.Wrap(err, "reading winterhill.init")
	}
	api := Api{
		eventListener:   events.New(4, 60),
		summaryListener: summary.New(),
		config:          iniConfig,
	}
	go api.eventListener.Run(fmt.Sprintf("%s:%d", ip, port), nil)
	go api.summaryListener.Run(fmt.Sprintf("%s:%d", ip, 9904), nil)
	return &api, nil
}

// StatusHandler returns the Receivers
func (A *Api) StatusHandler(w http.ResponseWriter, r *http.Request) {
	metrics.RequestMetrics.WithLabelValues("status").Inc()
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(A.eventListener.Receivers)
}

// SummaryHandler returns the Summary
func (A *Api) SummaryHandler(w http.ResponseWriter, r *http.Request) {
	metrics.RequestMetrics.WithLabelValues("status").Inc()
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(A.summaryListener.LastEvent)
}

func (A *Api) ConfigHandler(w http.ResponseWriter, r *http.Request) {
	metrics.RequestMetrics.WithLabelValues("config").Inc()
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(*A.config)
}
