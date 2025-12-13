package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

func GetVideoListHandler(w http.ResponseWriter, r *http.Request) {
	storages := filepath.Join(".", "storages")
	entries, err := os.ReadDir(storages)
	if err != nil {
		log.Printf("ReadDir command failed: %v", err)
		w.WriteHeader(500)
		w.Write([]byte(err.Error()))
		return
	}

	list := []string{}
	for _, v := range entries {
		if v.IsDir() {
			list = append(list, v.Name())
		}
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(list)
	if err != nil {
		log.Printf("Encode response failed: %v", err)
		w.WriteHeader(500)
		w.Write([]byte(err.Error()))
		return
	}
}
