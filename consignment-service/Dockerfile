# consignment-service/Dockerfile
FROM golang:1.9 as builder

WORKDIR /go/src/github.com/BradErz/shippy/consignment-service
COPY . .

RUN go get -u -v github.com/golang/dep/cmd/dep \
    && dep init \
    && dep ensure

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo .

# runner
FROM alpine:latest

RUN apk --no-cache add ca-certificates \
    && mkdir /app

WORKDIR /app

COPY --from=builder /go/src/github.com/BradErz/shippy/consignment-service/consignment-service .
CMD ["./consignment-service"]