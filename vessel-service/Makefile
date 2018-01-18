proto-def:
	protoc -I. --go_out=plugins=micro:$(GOPATH)/src/github.com/BradErz/shippy/vessel-service \
	proto/vessel/vessel.proto

build: proto-def
	docker build -t braderz/vessel-service .

run:
	docker run -p 50052:50051 \
	-e MICRO_SERVER_ADDRESS=:50051 \
	-e MICRO_REGISTRY=mdns \
	braderz/vessel-service