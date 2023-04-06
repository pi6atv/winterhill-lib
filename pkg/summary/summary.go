package summary

import (
	"fmt"
	"github.com/libp2p/go-reuseport"
	"github.com/pkg/errors"
	"strings"
	"time"
)

type Listener struct {
	LastEvent string
}

func New() *Listener {
	return &Listener{}
}

func (L *Listener) Run(addr string, stop chan bool) error {
	const bufLen = 1500
	var buf [bufLen]byte

	sock, err := reuseport.ListenPacket("udp", addr)
	if err != nil {
		return errors.Wrapf(err, "listen on '%s'", addr)
	}
	for {
		select {
		case <-stop:
			return nil
		default:
			_ = sock.SetReadDeadline(time.Now().Add(time.Second))
			size, _, err := sock.ReadFrom(buf[:])
			if err != nil || size == 0 {
				time.Sleep(time.Millisecond * 10) // 'rate limit'
				continue                          // fixme, handle error
			}

			if strings.HasPrefix(string(buf[0:4]), "$0,0") {
				//fmt.Println("skipping winterhill global config")
				continue
			}
			L.LastEvent = string(buf[0:size]) + fmt.Sprintf("\nupdated: %s", time.Now().String())
		}
	}
}
