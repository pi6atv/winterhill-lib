package main

import (
	"embed"
	"fmt"
	"github.com/gorilla/mux"
	commandapi "github.com/pi6atv/winterhill-lib/internal/web/command-api"
	statusapi "github.com/pi6atv/winterhill-lib/internal/web/status-api"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"io/fs"
	"net/http"
)

var (
	//go:embed dist
	all embed.FS
)

func main() {
	subdir, _ := fs.Sub(all, "dist")

	statusApi, err := statusapi.New("127.0.0.1", 9901)
	if err != nil {
		panic(err)
	}

	commandApi, err := commandapi.New("127.0.0.1", 9921)
	if err != nil {
		panic(err)
	}

	var router = mux.NewRouter()

	router.Path("/winterhill/api/status").HandlerFunc(statusApi.StatusHandler)
	router.Path("/winterhill/api/config").HandlerFunc(statusApi.ConfigHandler)
	router.Path("/winterhill/api/set/srate/{receiver:[1-4]}/{srate:[0-9]+}").HandlerFunc(commandApi.SetSymbolRateHandler).Methods("POST")
	router.Path("/winterhill/").Handler(http.StripPrefix("/winterhill/", http.FileServer(http.FS(subdir))))
	router.Path("/metrics").Handler(promhttp.Handler())

	fmt.Println("starting webserver")
	_ = http.ListenAndServe(":8080", router)
}
