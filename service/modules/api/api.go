package api

import "net/http"

func InitHttpServer() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		var data string = "Hello, World!"
		w.Write([]byte(data))
	})
}
