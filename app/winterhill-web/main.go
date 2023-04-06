package main

import (
	"embed"
	"flag"
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
	all                   embed.FS
	listenPort            = flag.String("listen", ":8080", "the address to listen on")
	winterhillIP          = flag.String("winterhill-ip", "127.0.0.1", "winterhill IP")
	winterhillCommandPort = flag.Int("winterhill-command-port", 9901, "winterhill command port")
	winterhillBasePort    = flag.Int64("winterhill-base-port", 9920, "winterhill base port")
	srResetDuration       = flag.Duration("symbol-rate-reset-duration", 30*time.Minute, "time after which the symbol rate is reset to default")
	webPath               = flag.String("web-path", "/winterhill", "base path for the web pages and api")
	//logEnabled            = flag.Bool()
)

func main() {
	flag.Parse()
	subdir, _ := fs.Sub(all, "dist")

	logStream := log_stream.New()

	statusApi, err := statusapi.New(*winterhillIP, *winterhillCommandPort)
	if err != nil {
		panic(err)
	}

	commandApi, err := commandapi.New(*winterhillIP, *winterhillBasePort, *srResetDuration, logStream)
	if err != nil {
		panic(err)
	}

	var router = mux.NewRouter()
	router.Use(middlewares.ExtractAuthMiddleware) // adds username from token to context

	logApi := log_api.New(logStream)
	apiRouter := router.Path(*webPath)
	apiRouter.Path("/api/log/ws").HandlerFunc(logApi.WsHandler)
	apiRouter.Path("/api/status").HandlerFunc(statusApi.StatusHandler)
	apiRouter.Path("/api/summary").HandlerFunc(statusApi.SummaryHandler)
	apiRouter.Path("/api/config").HandlerFunc(statusApi.ConfigHandler)
	apiRouter.Path("/api/set/srate/{receiver:[1-4]}/{srate:[0-9]+}").HandlerFunc(commandApi.SetSymbolRateHandler).Methods("POST")
	apiRouter.PathPrefix("/").Handler(http.StripPrefix(*webPath+"/", http.FileServer(http.FS(subdir))))

	router.Path("/metrics").Handler(promhttp.Handler())
	fmt.Println("starting webserver")
	_ = http.ListenAndServe(*listenPort, router)
}
