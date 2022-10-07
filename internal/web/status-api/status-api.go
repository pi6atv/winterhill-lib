package status_api

import (
	"encoding/json"
	"fmt"
	"github.com/libp2p/go-reuseport"
	"github.com/pi6atv/winterhill-lib/pkg/events"
	"net/http"
	"strings"
	"time"
)

type Api struct {
	addr      string
	Receivers []*events.StatusEvent
	stop      chan bool
}

func New(ip string, port int) (*Api, error) {

	api := Api{
		addr:      fmt.Sprintf("%s:%d", ip, port),
		stop:      make(chan bool),
		Receivers: make([]*events.StatusEvent, 4),
	}
	go api.listen()
	return &api, nil
}

// listen will receive the UDP messages from Winterhill and use them to update the Receivers in the struct
func (A *Api) listen() {
	const bufLen = 1024
	var buf [bufLen]byte

	sock, err := reuseport.ListenPacket("udp", A.addr)
	if err != nil {
		// fixme, handle error
		panic(err)
	}
	//_ = sock.SetReadBuffer(1048576)
	// fixme, handle error
	for {
		select {
		case <-A.stop:
			return
		default:
			sock.SetReadDeadline(time.Now().Add(time.Second))
			_, _, err := sock.ReadFrom(buf[:])
			if err != nil {
				time.Sleep(time.Millisecond * 10)
				continue // fixme, handle error
			}
			//fmt.Printf("received %d bytes, err: %+v\n", n, err)
			go func(in [bufLen]byte) {
				//fmt.Println("go func called")
				//fmt.Printf("received data: %s\n", in)
				for _, block := range strings.Split(string(in[:]), "\000") {
					//fmt.Printf("handling block: %s\n", block)
					event, err := events.ParseEvent(block)
					if event == nil || err != nil || event.Index == 0 {
						//fmt.Printf("error on parse event: %+v, event: %+v", err, event)
						continue // fixme, handle error
					}
					fmt.Printf("created event %+v\n", *event)
					A.Receivers[event.Index-1] = event
				}
			}(buf)
		}
	}
}

// StatusHandler returns the Receivers
func (A *Api) StatusHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(A.Receivers)
}
