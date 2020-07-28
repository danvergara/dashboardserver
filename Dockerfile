FROM golang:1.14 AS base

WORKDIR /app

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o app cmd/api/main.go

FROM alpine:latest as certs
RUN apk --update add ca-certificates

WORKDIR /app

COPY --from=base /app/app .
COPY --from=base /app/swagger.yaml .
ENTRYPOINT ["./app"]
EXPOSE 8000
