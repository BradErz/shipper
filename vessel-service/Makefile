goimports:
	goimports -w .

deps:
	dep ensure -v

proto:
	protoc -I. --go_out=plugins=micro:$(GOPATH)/src/github.com/BradErz/shipper/vessel-service \
	proto/vessel/vessel.proto

build: goimports deps proto
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo .
	docker build -t vessel-service .

run: build
	docker run -p 50052:50051 \
        -e MICRO_SERVER_ADDRESS=:50051 \
        -e MICRO_REGISTRY=mdns vessel-service