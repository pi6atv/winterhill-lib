package status_api

import (
	"encoding/json"
	"fmt"
	"github.com/pi6atv/winterhill-lib/internal/web/metrics"
	"github.com/pi6atv/winterhill-lib/pkg/config"
	"github.com/pi6atv/winterhill-lib/pkg/events"
	"github.com/pkg/errors"
	"net/http"
)

type Api struct {
	listener *events.Listener
	config   *config.WinterhillConfig
}

func New(ip string, port int) (*Api, error) {
	iniConfig, err := config.New("")
	if err != nil {
		return nil, errors.Wrap(err, "reading winterhill.init")
	}
	api := Api{
		listener: events.New(4, 60),
		config:   iniConfig,
	}
	go api.listener.Run(fmt.Sprintf("%s:%d", ip, port), nil)
	return &api, nil
}

// StatusHandler returns the Receivers
func (A *Api) StatusHandler(w http.ResponseWriter, r *http.Request) {
	metrics.RequestMetrics.WithLabelValues("status").Inc()
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(A.listener.Receivers)
}

func (A *Api) ConfigHandler(w http.ResponseWriter, r *http.Request) {
	metrics.RequestMetrics.WithLabelValues("config").Inc()
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(*A.config)
}
