package queue

import (
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/storage/azqueue"
)

var Queue *azqueue.QueueClient

func Init() {

	accountName, accountKey, serviceURL, queueName := accountInfo()

	log.Printf("Queue initialization: %s/%s", serviceURL, queueName)

	cred, err := azqueue.NewSharedKeyCredential(accountName, accountKey)
	handleError(err)
	serviceClient, err := azqueue.NewServiceClientWithSharedKeyCredential(serviceURL, cred, nil)
	handleError(err)

	Queue = serviceClient.NewQueueClient(queueName)
}
