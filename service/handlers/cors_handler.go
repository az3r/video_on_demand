package handlers

import "net/http"

func HandleCORS(w http.ResponseWriter, r *http.Request) {
	// Allow all origins
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

	if r.Method != http.MethodOptions {
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func WithCORS(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		HandleCORS(w, r)
		handler.ServeHTTP(w, r)
	})
}
