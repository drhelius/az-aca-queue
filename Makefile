APP_NAME=azqueue
BINARY_PATH=bin

CMD_MAIN_PATH=cmd/$(APP_NAME)/main.go
CMD_BINARY_PATH=$(BINARY_PATH)/$(APP_NAME)
CMD_SERVER_PORT=8080


clean:
	go clean $(CMD_MAIN_PATH)
	rm -f $(BINARY_PATH)/*

build:
	go build -o $(CMD_BINARY_PATH)

run:
	$(CMD_BINARY_PATH)

docker-build:
	CGO_ENABLED=0 GOOS=linux go build -a -o $(CMD_BINARY_PATH)
	docker build -t $(APP_NAME) .

docker-run:
	. ./test.env
	docker run --env-file test.env -p $(CMD_SERVER_PORT):$(CMD_SERVER_PORT) $(APP_NAME)
