package main

import (
	"context"

	"az3r.me.video_on_demand/modules/api"
	"az3r.me.video_on_demand/modules/db"
	"az3r.me.video_on_demand/modules/ytb"
)

func main() {
	db.InitDatabase()
	defer db.Conn.Close(context.Background())
	api.InitHttpServer()

	ytb.DownloadVideo("https://www.youtube.com/watch?v=6KqQwn4Y9UA")
	// log.Printf("Listening on port 3000")
	// http.ListenAndServe(":3000", nil)
}
