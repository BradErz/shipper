build:
	GOOS=linux GOARCH=amd64 go build .
	docker build -t braderz/consignment-cli .

run:
	docker run -e MICRO_REGISTRY=mdns braderz/consignment-cli