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
	"strings"
	"time"
)

var (
	//go:embed dist
	dist                  embed.FS
	listenPort            = flag.String("listen", ":80", "the address to listen on")
	winterhillIniPath     = flag.String("winterhill-ini-path", "/home/pi/winterhill/winterhill.ini", "path to winterhill.ini")
	winterhillIP          = flag.String("winterhill-ip", "127.0.0.1", "winterhill IP")
	winterhillCommandPort = flag.Int("winterhill-command-port", 9901, "winterhill command port")
	winterhillBasePort    = flag.Int64("winterhill-base-port", 9920, "winterhill base port")
	srResetDuration       = flag.Duration("symbol-rate-reset-duration", 30*time.Minute, "time after which the symbol rate is reset to default")
	webSubPath            = flag.String("web-path", "/", "base path for the web pages and api")
)

func main() {
	flag.Parse()
	if !strings.HasSuffix(*webSubPath, "/") {
		*webSubPath = *webSubPath + "/"
	}

	// keeps a log of actions of users
	logStream := log_stream.New()

	statusApi, err := statusapi.New(*winterhillIP, *winterhillCommandPort, *winterhillIniPath)
	if err != nil {
		panic(err)
	}

	commandApi, err := commandapi.New(*winterhillIP, *winterhillBasePort, *srResetDuration, logStream, *winterhillIniPath)
	if err != nil {
		panic(err)
	}

	var rootRouter = mux.NewRouter()
	rootRouter.Path("/metrics").Handler(promhttp.Handler())
	subPathRouter := rootRouter
	if *webSubPath != "/" {
		subPathRouter = rootRouter.PathPrefix(*webSubPath).Subrouter()
	}

	logApi := log_api.New(logStream)
	apiRouter := subPathRouter.PathPrefix("/api").Subrouter()
	apiRouter.Path("/log/ws").HandlerFunc(logApi.WsHandler)
	apiRouter.Path("/config").HandlerFunc(statusApi.ConfigHandler)
	apiRouter.Path("/status").HandlerFunc(statusApi.StatusHandler)
	apiRouter.Path("/summary").HandlerFunc(statusApi.SummaryHandler)
	apiRouter.
		Path("/set/srate/{receiver:[1-4]}/{srate:[0-9]+}").
		HandlerFunc(commandApi.SetSymbolRateHandler).
		Methods("POST")

	static, _ := fs.Sub(dist, "dist")
	staticServer := http.FileServer(http.FS(static))
	subPathRouter.PathPrefix("/").Handler(http.StripPrefix(*webSubPath, staticServer))

	fmt.Println("starting webserver")
	// print all routes
	rootRouter.Walk(func(route *mux.Route, router *mux.Router, ancestors []*mux.Route) error {
		t, err := route.GetPathTemplate()
		if err != nil {
			return err
		}
		fmt.Println(t)
		return nil
	})

	_ = http.ListenAndServe(*listenPort, rootRouter)
}
