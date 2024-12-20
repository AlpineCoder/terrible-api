package debug

import (
	"net/http"
)

// We're always OK
func DebugHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}
