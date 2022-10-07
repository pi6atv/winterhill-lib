package main

import (
	"embed"
	status_api "github.com/pi6atv/winterhill-lib/internal/web/status-api"
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
	http.ListenAndServe(":8080", nil)
}
