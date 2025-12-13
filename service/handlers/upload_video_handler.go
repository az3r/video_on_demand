package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"az3r.me.video_on_demand/dtos"
	"az3r.me.video_on_demand/modules/ytb"
)

func HandleUploadVideo() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			w.WriteHeader(http.StatusMethodNotAllowed)
			w.Write([]byte("Method is not allowed"))
			return
		}

		var dto dtos.UploadDto
		err := json.NewDecoder(r.Body).Decode(&dto)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}

		metadata, err := ytb.DownloadVideo(dto.Url)
		if err != nil {
			log.Printf("Failed to download video: %v", err)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}

		result, err := ytb.ConvertToHls(metadata)
		if err != nil {
			os.Remove(metadata.FilePath)
			log.Printf("Failed to convert video to HLS format: %v", err)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}

		w.Header().Set("Content-Type", "application/json")
		err = json.NewEncoder(w).Encode(result)
		if err != nil {
			log.Printf("Failed to send response: %v", err)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}
	})
}
