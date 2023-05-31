package main

import (
	"github.com/drhelius/az-aca-queue/internal/http"
	"github.com/drhelius/az-aca-queue/internal/queue"
)

func main() {
	queue.Init()
	http.Serve("8080")
}
