package api

import (
	"net/http"
)

// HelloHandler responds with a JSON message
func HelloHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{"message":"Hello from backend!"}`))
}

