package log

import (
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/pi6atv/winterhill-lib/pkg/log"
	"net/http"
)

type Api struct {
	hub *hub
}

func New(log *log.Stream) *Api {
	hub := NewHub(log)
	go hub.Start()
	return &Api{
		hub: hub,
	}
}

func (A *Api) WsHandler(w http.ResponseWriter, r *http.Request) {
	upgrader := &websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool { return true },
	}
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Printf("ws upgrade failed: %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	client := NewClient(conn, A.hub)
	go client.handle()

	A.hub.register <- client
}
