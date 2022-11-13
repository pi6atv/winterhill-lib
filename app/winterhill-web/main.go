package main

import (
	"embed"
	"fmt"
	"github.com/gorilla/mux"
	commandapi "github.com/pi6atv/winterhill-lib/internal/web/command-api"
	log_api "github.com/pi6atv/winterhill-lib/internal/web/log"
	"github.com/pi6atv/winterhill-lib/internal/web/middlewares"
	statusapi "github.com/pi6atv/winterhill-lib/internal/web/status-api"
	log_stream "github.com/pi6atv/winterhill-lib/pkg/log"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"io/fs"
	"net/http"
	"time"
)

var (
	//go:embed dist
	all embed.FS
)

func main() {
	subdir, _ := fs.Sub(all, "dist")

	logStream := log_stream.New()

	statusApi, err := statusapi.New("127.0.0.1", 9901)
	if err != nil {
		panic(err)
	}

	commandApi, err := commandapi.New("127.0.0.1", 9920, 10*time.Minute, logStream)
	if err != nil {
		panic(err)
	}

	var router = mux.NewRouter()
	router.Use(middlewares.ExtractAuthMiddleware) // adds username from token to context

	logApi := log_api.New(logStream)
	router.Path("/winterhill/api/log/ws").HandlerFunc(logApi.WsHandler)
	router.Path("/winterhill/api/status").HandlerFunc(statusApi.StatusHandler)
	router.Path("/winterhill/api/config").HandlerFunc(statusApi.ConfigHandler)
	router.Path("/winterhill/api/set/srate/{receiver:[1-4]}/{srate:[0-9]+}").HandlerFunc(commandApi.SetSymbolRateHandler).Methods("POST")
	router.PathPrefix("/winterhill/").Handler(http.StripPrefix("/winterhill/", http.FileServer(http.FS(subdir))))

	router.Path("/metrics").Handler(promhttp.Handler())
	fmt.Println("starting webserver")
	_ = http.ListenAndServe(":8080", router)
}
