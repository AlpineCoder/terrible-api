package main

import (
	"log"
	"net/http"

	"github.com/AlpineCoder/terrible-api/router"
)

func main() {
	r := router.NewRouter()
	log.Println("Server starting on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
