package command_api

import (
	"github.com/gorilla/mux"
	"github.com/pi6atv/winterhill-lib/internal/web/metrics"
	"github.com/pi6atv/winterhill-lib/pkg/commands"
	"github.com/pi6atv/winterhill-lib/pkg/config"
	"github.com/pi6atv/winterhill-lib/pkg/log"
	"github.com/pkg/errors"
	"net/http"
	"strconv"
	"time"
)

var validSymbolRates = []int64{25, 35, 66, 125, 250, 333, 360, 500, 1000, 1200, 1500, 2000, 3000, 4000, 4167, 22000, 27500}

type Api struct {
	remoteHost     string
	remoteBasePort int64
	config         *config.WinterhillConfig
	resetInterval  time.Duration
	log            *log.Stream
}

func New(ip string, basePort int64, resetInterval time.Duration, logStream *log.Stream) (*Api, error) {
	iniConfig, err := config.New("")
	if err != nil {
		return nil, errors.Wrap(err, "reading winterhill.init")
	}

	api := Api{
		remoteHost:     ip,
		remoteBasePort: basePort,
		config:         iniConfig,
		resetInterval:  resetInterval,
		log:            logStream,
	}
	return &api, nil
}

// SetSymbolRateHandler will set the symbol rate for the specified receiver
func (A *Api) SetSymbolRateHandler(w http.ResponseWriter, r *http.Request) {
	metrics.RequestMetrics.WithLabelValues("set/srate").Inc()
	vars := mux.Vars(r)
	receiver, err := strconv.ParseInt(vars["receiver"], 10, 64)
	if err != nil {
		metrics.RequestErrorMetrics.WithLabelValues("set/srate").Inc()
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	srate, err := strconv.ParseInt(vars["srate"], 10, 64)
	if err != nil || !IsvalidSymbolRate(srate) {
		metrics.RequestErrorMetrics.WithLabelValues("set/srate").Inc()
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	A.log.Log(r, receiver, "srate", vars["srate"])
	command := A.getPresetCommand(receiver)
	command.SymbolRate = srate
	err = command.Send(A.remoteHost, A.remoteBasePort)
	if err != nil {
		metrics.RequestErrorMetrics.WithLabelValues("set/srate").Inc()
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// send reset
	go func() {
		time.Sleep(A.resetInterval)
		metrics.RequestMetrics.WithLabelValues("set/srate reset").Inc()
		command := A.getPresetCommand(receiver)
		err := command.Send(A.remoteHost, A.remoteBasePort)
		if err != nil {
			metrics.RequestErrorMetrics.WithLabelValues("set/srate reset").Inc()
		}
	}()
}

// IsvalidSymbolRate will return if the given symbol rate exists in the list of valid ones
func IsvalidSymbolRate(srate int64) bool {
	for _, valid := range validSymbolRates {
		if srate == valid {
			return true
		}
	}
	return false
}

// getPresetCommand will prefill a commands.WhShort struct with the values found in the config file
func (A *Api) getPresetCommand(receiver int64) commands.WhShort {
	return commands.WhShort{
		Index:      receiver,
		Frequency:  A.config.Receivers[receiver-1].Frequency,
		Offset:     A.config.Receivers[receiver-1].Offset,
		SymbolRate: A.config.Receivers[receiver-1].SymbolRate,
		Antenna:    A.config.Receivers[receiver-1].Antenna,
		Voltage:    A.config.Receivers[receiver-1].Voltage,
	}
}
