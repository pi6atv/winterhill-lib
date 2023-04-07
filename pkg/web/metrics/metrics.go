package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	RequestMetrics      = promauto.NewGaugeVec(prometheus.GaugeOpts{Name: "winterhill_web_requests_total", Help: "web requests to the exporter"}, []string{"path"})
	RequestErrorMetrics = promauto.NewGaugeVec(prometheus.GaugeOpts{Name: "winterhill_web_errors_total", Help: "web request errors from the exporter"}, []string{"path"})
)
