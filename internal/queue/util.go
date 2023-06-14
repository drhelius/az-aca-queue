package queue

import (
	"log"
	"os"
)

func accountInfo() (string, string, string, string, string) {
	return os.Getenv("ACCOUNT_NAME"), os.Getenv("ACCOUNT_KEY"), os.Getenv("SERVICE_URL"),
		os.Getenv("STATIC_QUEUE_NAME"), os.Getenv("RANDOM_QUEUE_PREFIX")
}

func handleError(err error) {
	if err != nil {
		log.Fatal(err.Error())
	}
}
