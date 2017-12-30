goimports:
	goimports -w .

deps:
	dep ensure -v

proto:
	protoc -I. --go_out=plugins=micro:$(GOPATH)/src/github.com/BradErz/shipper/consignment-service \
    proto/consignment/consignment.proto

build: goimports proto deps
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo .
	docker build -t consignment-service .

run: build
	docker run -p 50051:50051 \
        -e MICRO_SERVER_ADDRESS=:50051 \
        -e MICRO_REGISTRY=mdns consignment-service