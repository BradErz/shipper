proto-def:
	protoc -I. --go_out=plugins=micro:$(GOPATH)/src/github.com/BradErz/shippy/consignment-service \
	proto/consignment/consignment.proto

build: proto-def
	docker build -t braderz/consignment-service .

run:
	docker run -p 50051:50051 \
	 -e MICRO_SERVER_ADDRESS=:50051 \
	 -e MICRO_REGISTRY=mdns \
	 braderz/consignment-service