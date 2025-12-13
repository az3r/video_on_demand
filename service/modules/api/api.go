package api

import (
	"net/http"

	"az3r.me.video_on_demand/handlers"
)

func InitHttpServer() {
	storages := http.FileServer(http.Dir("./storages"))
	http.Handle("/storages/", handlers.WithCORS(http.StripPrefix("/storages/", storages)))

	http.Handle("/upload", handlers.WithCORS(handlers.HandleUploadVideo()))
	http.Handle("/get_video_list", handlers.WithCORS(handlers.GetVideoListHandler()))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		var data string = "Hello, World!"
		w.Write([]byte(data))
	})
}
