package main

import (
	"embed"
	"flag"
	"fmt"
	"github.com/gorilla/mux"
	log_stream "github.com/pi6atv/winterhill-lib/pkg/log"
	commandapi "github.com/pi6atv/winterhill-lib/pkg/web/command-api"
	log_api "github.com/pi6atv/winterhill-lib/pkg/web/log"
	statusapi "github.com/pi6atv/winterhill-lib/pkg/web/status-api"
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
	webSubPath            = flag.String("web-path", "/winterhill", "base path for the web pages and api")
)

func main() {
	flag.Parse()
	subdir, _ := fs.Sub(all, "dist")

	// keeps a log of actions of users
	logStream := log_stream.New()

	statusApi, err := statusapi.New(*winterhillIP, *winterhillCommandPort)
	if err != nil {
		panic(err)
	}

	commandApi, err := commandapi.New(*winterhillIP, *winterhillBasePort, *srResetDuration, logStream)
	if err != nil {
		panic(err)
	}

	var rootRouter = mux.NewRouter()
	subPathRouter := rootRouter.Path(*webSubPath)
	rootRouter.Path("/metrics").Handler(promhttp.Handler())

	logApi := log_api.New(logStream)
	apiRouter := subPathRouter.Path("/api")
	apiRouter.Path("/log/ws").HandlerFunc(logApi.WsHandler)
	apiRouter.Path("/config").HandlerFunc(statusApi.ConfigHandler)
	apiRouter.Path("/status").HandlerFunc(statusApi.StatusHandler)
	apiRouter.Path("/summary").HandlerFunc(statusApi.SummaryHandler)
	apiRouter.
		Path("/set/srate/{receiver:[1-4]}/{srate:[0-9]+}").
		HandlerFunc(commandApi.SetSymbolRateHandler).
		Methods("POST")
	subPathRouter.PathPrefix("/").Handler(http.StripPrefix(*webSubPath+"/", http.FileServer(http.FS(subdir))))

	fmt.Println("starting webserver")
	_ = http.ListenAndServe(*listenPort, rootRouter)
}
