FROM golang:1.19-alpine AS builder

RUN apk update && apk add alpine-sdk git && rm -rf /var/cache/apk/*

RUN mkdir -p /api
WORKDIR /api

COPY go.mod .
COPY go.sum .
RUN go mod download

RUN go mod tidy

COPY . .

RUN go build -o ./app ./main.go
ENV PATH=$PATH:/api

ENTRYPOINT ["app"]
