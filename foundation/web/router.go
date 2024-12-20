package web

import (
	"net/http"

	debug "github.com/AlpineCoder/terrible-api/api/debug/handlers"
	"github.com/AlpineCoder/terrible-api/api/v1/handlers"
	"github.com/AlpineCoder/terrible-api/middleware"
	"github.com/AlpineCoder/terrible-api/store"
)

// Router for live- and readiness probes

// Authenticated Router
func NewRouter() http.Handler {

	store := store.NewBackend()

	// Create handler with the datastore injected
	handlerWithStore := handlers.HandlerWithStore{Store: store}

	mainMux := http.NewServeMux()

	// Healthz probes
	mainMux.HandleFunc(HealthzLivenessPath, debug.DebugHandler)
	mainMux.HandleFunc(HealthzReadinessPath, debug.DebugHandler)

	// Authenticated API routes
	authMux := http.NewServeMux()
	authMux.HandleFunc("/api/hello", handlerWithStore.HelloHandler)
	authMux.HandleFunc("/api/monsters", handlerWithStore.MonstersHandler)
	// authMux.HandleFunc("/api/monsters/", handlerWithStore.HelloHandler)

	authenticatedRoutes := middleware.BasicAuth(authMux)

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case HealthzLivenessPath, HealthzReadinessPath:
			mainMux.ServeHTTP(w, r)
		default:
			authenticatedRoutes.ServeHTTP(w, r)
		}
	})
}
