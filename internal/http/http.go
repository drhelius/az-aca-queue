package http

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/Pallinder/go-randomdata"
	"github.com/drhelius/az-aca-queue/internal/queue"
)

func Serve(http_port string) {
	mux := http.NewServeMux()
	mux.HandleFunc("/health", health)
	mux.HandleFunc("/queue-single", insertIntoQueueSingle)
	mux.HandleFunc("/queue-random", insertIntoQueueRandom)
	mux.HandleFunc("/event-grid", insertIntoEventGrid)

	log.Printf("HTTP server listening on port %s...", http_port)

	err := http.ListenAndServe(fmt.Sprintf(":%s", http_port), mux)
	if err != nil {
		log.Fatal(err)
	}
}

func health(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "OK")
}

func insertIntoQueueSingle(w http.ResponseWriter, r *http.Request) {

	insertions := 1
	hasCount := r.URL.Query().Has("count")

	if hasCount {
		countStr := r.URL.Query().Get("count")
		count, err := strconv.Atoi(countStr)

		if err != nil {
			log.Printf("Bad count parameter")
		} else {
			insertions = count
		}
	}

	for i := 0; i < insertions; i++ {
		queue.InsertSingleMessage(randomdata.Alphanumeric(128))
	}
}

func insertIntoQueueRandom(w http.ResponseWriter, r *http.Request) {

	insertions := 1
	hasCount := r.URL.Query().Has("count")

	if hasCount {
		countStr := r.URL.Query().Get("count")
		count, err := strconv.Atoi(countStr)

		if err != nil {
			log.Printf("Bad count parameter")
		} else {
			insertions = count
		}
	}

	queueCount := 10
	hasQueues := r.URL.Query().Has("queues")

	if hasQueues {
		queuesStr := r.URL.Query().Get("queues")
		queues, err := strconv.Atoi(queuesStr)

		if err != nil {
			log.Printf("Bad queues parameter")
		} else {
			queueCount = queues
		}
	}

	for i := 0; i < insertions; i++ {
		queue.InsertRandomMessage(randomdata.Alphanumeric(128), queueCount)
	}
}

func insertIntoEventGrid(_ http.ResponseWriter, _ *http.Request) {
	// TODO
}
