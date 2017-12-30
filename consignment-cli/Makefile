goimports:
	goimports -w .

build: goimports
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo .
	docker build -t consignment-cli .

run: build
	docker run -e MICRO_REGISTRY=mdns consignment-cli