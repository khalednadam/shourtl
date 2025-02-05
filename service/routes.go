package service

import (
	"khalednadam/shourtl/cmd/api"
	"net/http"
)

// ✅ Corrected function
func TestRoutes(router *api.APIServer) {
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// ✅ Ensure only GET requests are allowed
		if r.Method != http.MethodGet {
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
			return
		}

		w.Write([]byte("Hello, world"))
	})
}
