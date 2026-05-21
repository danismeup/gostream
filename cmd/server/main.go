package main

import (
	"log"
	"net/http"

	"github.com/danismeup/gostream/internal/signaling"
)

func main() {
	http.HandleFunc("/ws", signaling.HandleWebSocket)
	http.Handle("/", http.FileServer(http.Dir("web/client")))

	log.Println("GoStream server starting on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
