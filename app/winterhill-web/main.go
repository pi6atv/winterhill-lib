package main

import (
	"embed"
	"fmt"
	status_api "github.com/pi6atv/winterhill-lib/internal/web/status-api"
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

	webApi, err := status_api.New("127.0.0.1", 9901)
	if err != nil {
		panic(err)
	}
	http.HandleFunc("/api/status", webApi.StatusHandler)
	http.Handle("/", http.FileServer(http.FS(subdir)))
	http.Handle("/metrics", promhttp.Handler())
	fmt.Println("starting webserver")
	http.ListenAndServe(":8080", nil)
}
