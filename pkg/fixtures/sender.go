package fixtures

import (
	"fmt"
	"github.com/pkg/errors"
	"net"
)

func Send(data []string, port int) error {
	conn, err := net.DialUDP("udp", nil, &net.UDPAddr{IP: net.ParseIP("127.0.0.1"), Port: port})
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("connecting to UDP port %d", port))
	}
	defer conn.Close()
	for _, packet := range data {
		_, err := conn.Write([]byte(packet))
		if err != nil {
			return errors.Wrap(err, "failed sending packet")
		}
	}
	return nil
}
