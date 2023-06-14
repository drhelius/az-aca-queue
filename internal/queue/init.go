package queue

import (
	"log"
	"strconv"

	"github.com/Azure/azure-sdk-for-go/sdk/storage/azqueue"
)

const MaxRandomQueues = 10

var StaticQueue *azqueue.QueueClient
var RandomQueue [MaxRandomQueues]*azqueue.QueueClient

func Init() {

	accountName, accountKey, serviceURL, staticQueueName, randomQueuePrefix := accountInfo()

	cred, err := azqueue.NewSharedKeyCredential(accountName, accountKey)
	handleError(err)
	serviceClient, err := azqueue.NewServiceClientWithSharedKeyCredential(serviceURL, cred, nil)
	handleError(err)

	log.Printf("Static Queue initialization: %s/%s", serviceURL, staticQueueName)

	StaticQueue = serviceClient.NewQueueClient(staticQueueName)

	for i := 0; i < MaxRandomQueues; i++ {
		queueName := randomQueuePrefix + strconv.Itoa(i)
		log.Printf("Random Queue initialization: %s/%s", serviceURL, queueName)
		RandomQueue[i] = serviceClient.NewQueueClient(queueName)
	}
}
