FROM golang:1.21.0-alpine AS builder
WORKDIR /go/src/github.com/AH-dark/epay-cli

COPY . .

RUN go build -o epay-cli .

# Path: Dockerfile
FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /app/

COPY --from=builder /go/src/github.com/AH-dark/epay-cli/epay-cli /app/

ENTRYPOINT ["/app/epay-cli"]
