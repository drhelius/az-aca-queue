package main

import (
	"github.com/drhelius/az-function-queue/internal/http"
	"github.com/drhelius/az-function-queue/internal/queue"
)

func main() {
	queue.Init()
	http.Serve("8080")
}
