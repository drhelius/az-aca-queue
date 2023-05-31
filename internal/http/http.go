package http

import (
	"fmt"
	"log"
	"net/http"

	"github.com/drhelius/az-function-queue/internal/queue"
)

func Serve(http_port string) {
	mux := http.NewServeMux()
	mux.HandleFunc("/insert", insertIntoQueue)

	log.Printf("HTTP server listening on port %s...", http_port)

	err := http.ListenAndServe(fmt.Sprintf(":%s", http_port), mux)
	if err != nil {
		log.Fatal(err)
	}
}

func insertIntoQueue(w http.ResponseWriter, r *http.Request) {
	// log.Println("Insertion HTTP handler")
	queue.InsertMessage("test content")
}
