package handlers

import (
	"encoding/json"
	"net/http"
)

func (h *HandlerWithStore) HelloHandler(w http.ResponseWriter, r *http.Request) {

	name := r.URL.Query().Get("name")
	if name == "" {
		name = "world"
	}

	response := map[string]string{"message": "Hello, " + name + "!"}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
