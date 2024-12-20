package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/AlpineCoder/terrible-api/backend"
)

type MonsterPayload struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
	Size string `json:"size"`
}

func (h *HandlerWithStore) CreateMonsterHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	var payload backend.Monster

	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// h.Store.Set(h.Store.NextID(), payload)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(payload)
}
