package status_api

import (
	"encoding/json"
	"fmt"
	"github.com/pi6atv/winterhill-lib/pkg/events"
	"net/http"
)

type Api struct {
	listener *events.Listener
}

func New(ip string, port int) (*Api, error) {
	api := Api{
		listener: events.New(4, 30),
	}
	go api.listener.Run(fmt.Sprintf("%s:%d", ip, port), nil)
	return &api, nil
}

// StatusHandler returns the Receivers
func (A *Api) StatusHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(A.listener.Receivers)
}
