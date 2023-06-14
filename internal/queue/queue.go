package queue

import (
	"context"
	"log"

	"github.com/Pallinder/go-randomdata"
)

func InsertSingleMessage(message string) {
	// log.Printf("Inserting '%s' message into queue", message)

	// opts := &azqueue.EnqueueMessageOptions{TimeToLive: to.Ptr(int32(10))}
	_, err := StaticQueue.EnqueueMessage(context.Background(), message, nil)
	handleError(err)
	// log.Printf("Enqueue: %+v", resp1)

	// resp2, err := Queue.DequeueMessage(context.Background(), nil)
	// handleError(err)
	// log.Printf("Dequeue: %+v", resp2)
	// log.Printf("Dequeue: %s", *resp2.Messages[0].MessageText)
}

func InsertRandomMessage(message string, queueCount int) {

	if queueCount <= 0 {
		log.Printf("Invalid queueCount: %d", queueCount)
		return
	}

	i := randomdata.Number(queueCount)
	_, err := RandomQueue[i].EnqueueMessage(context.Background(), message, nil)
	handleError(err)
}
