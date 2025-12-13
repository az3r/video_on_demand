package api

import (
	"net/http"

	"az3r.me.video_on_demand/handlers"
)

func InitHttpServer() {
	storages := http.FileServer(http.Dir("./storages"))
	http.Handle("/storages/", http.StripPrefix("/storages/", storages))

	http.HandleFunc("/upload", handlers.HandleUploadVideo)
	http.HandleFunc("/get_video_list", handlers.GetVideoListHandler)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		var data string = "Hello, World!"
		w.Write([]byte(data))
	})
}
