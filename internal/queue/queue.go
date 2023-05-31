package queue

import (
	"context"
)

func InsertMessage(message string) {
	// log.Printf("Inserting '%s' message into queue", message)

	// opts := &azqueue.EnqueueMessageOptions{TimeToLive: to.Ptr(int32(10))}
	_, err := Queue.EnqueueMessage(context.Background(), message, nil)
	handleError(err)
	// log.Printf("Enqueue: %+v", resp1)

	// resp2, err := Queue.DequeueMessage(context.Background(), nil)
	// handleError(err)
	// log.Printf("Dequeue: %+v", resp2)
	// log.Printf("Dequeue: %s", *resp2.Messages[0].MessageText)
}
