package main

import (
	"log"
	"net/http"

	"az3r.me.video_on_demand/modules"
)

func main() {
	modules.InitDatabase()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		var data string = "Hello, World!"
		w.Write([]byte(data))

	})

	log.Printf("Listening on port 3000")
	http.ListenAndServe(":3000", nil)
}
