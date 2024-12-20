package handlers

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/AlpineCoder/terrible-api/business/models"
)

func (h *HandlerWithStore) MonstersHandler(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case http.MethodGet:
		// Return a list of users
		h.ListMonsters(w, r)
	case http.MethodPost:
		// Create a new user
		h.CreateMonster(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func (h *HandlerWithStore) MonsterHandler(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/monsters/")
	if id == "" {
		http.Error(w, "User ID is required", http.StatusBadRequest)
		return
	}

	switch r.Method {
	// case http.MethodGet:
	// 	h.GetMonster(w, r, id)
	// case http.MethodPut:
	// 	h.UpdateMonster(w, r, id)
	// case http.MethodDelete:
	// 	h.DeleteMonster(w, r, id)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func (h *HandlerWithStore) ListMonsters(w http.ResponseWriter, r *http.Request) {
	monsters := make(map[int]models.Monster)

	// Simulate fetching users from the datastore
	h.Store.Mu.Lock()
	for id, monster := range h.Store.Monsters {
		monsters[id] = monster
	}
	h.Store.Mu.Unlock()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(monsters)
}

func (h *HandlerWithStore) CreateMonster(w http.ResponseWriter, r *http.Request) {
	var monster models.Monster
	if err := json.NewDecoder(r.Body).Decode(&monster); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	id := h.Store.NextID()
	h.Store.Set(id, monster)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "Monster created"})
}
