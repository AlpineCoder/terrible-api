package main

import (
	"log"
	"net/http"

	"github.com/AlpineCoder/terrible-api/foundation/web"
)

func main() {
	w := web.NewRouter()
	log.Println("Server starting on :8080")
	log.Fatal(http.ListenAndServe(":8080", w))
}
