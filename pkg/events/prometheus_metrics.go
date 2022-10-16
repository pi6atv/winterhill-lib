package events

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	promReceiverUpdate = promauto.NewCounterVec(
		prometheus.CounterOpts{
			Name: "winterhill_receiver_updates",
			Help: "increases with every update received from Winterhill, for the given receiver",
		},
		[]string{"receiver"},
	)
	promReceiverMer = promauto.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "winterhill_receiver_mer",
			Help: "MER, for the given receiver",
		},
		[]string{"receiver"},
	)
	promReceiverDnumber = promauto.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "winterhill_receiver_dnumber",
			Help: "D number, for the given receiver",
		},
		[]string{"receiver"},
	)
	promReceiverSR = promauto.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "winterhill_receiver_symbol_rate",
			Help: "Symbol rate, for the given receiver",
		},
		[]string{"receiver"},
	)
	promReceiverState = promauto.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "winterhill_receiver_status",
			Help: "Status, for the given receiver",
		},
		[]string{"receiver"},
	)
	promReceiverFreq = promauto.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "winterhill_receiver_freq",
			Help: "Frequency, for the given receiver",
		},
		[]string{"receiver"},
	)
)
