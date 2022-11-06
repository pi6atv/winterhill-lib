package command_api

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/pi6atv/winterhill-lib/pkg/commands"
	"github.com/pi6atv/winterhill-lib/pkg/config"
	"github.com/pkg/errors"
	"net/http"
	"strconv"
	"time"
)

var validSymbolRates = []int64{25, 35, 66, 125, 250, 333, 360, 500, 1000, 1200, 1500, 2000, 3000, 4000, 4167, 22000, 27500}

type Api struct {
	Remote string
	config *config.WinterhillConfig
}

func New(ip string, port int) (*Api, error) {
	iniConfig, err := config.New("")
	if err != nil {
		return nil, errors.Wrap(err, "reading winterhill.init")
	}

	api := Api{
		Remote: fmt.Sprintf("%s:%d", ip, port),
		config: iniConfig,
	}
	return &api, nil
}

// SetSymbolRateHandler will set the symbol rate for the specified receiver
func (A *Api) SetSymbolRateHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	receiver, err := strconv.ParseInt(vars["receiver"], 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	srate, err := strconv.ParseInt(vars["srate"], 10, 64)
	if err != nil || !IsvalidSymbolRate(srate) {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	command := A.getPresetCommand(receiver)
	command.SymbolRate = srate
	err = command.Send(A.Remote)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// send reset
	go func() {
		time.Sleep(10 * time.Minute)
		command := A.getPresetCommand(receiver)
		command.Send(A.Remote)
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
		Frequency:  A.config.Receivers[receiver].Frequency,
		Offset:     A.config.Receivers[receiver].Offset,
		SymbolRate: A.config.Receivers[receiver].SymbolRate,
		Antenna:    A.config.Receivers[receiver].Antenna,
		Voltage:    A.config.Receivers[receiver].Voltage,
	}
}
