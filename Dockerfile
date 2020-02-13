FROM golang:1.13 AS base

WORKDIR /app

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o app .

FROM alpine:latest as certs
RUN apk --update add ca-certificates

WORKDIR /app

COPY --from=base /app/app .
ENTRYPOINT ["./app"]
