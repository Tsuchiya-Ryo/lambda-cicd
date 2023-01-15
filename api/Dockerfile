FROM golang:1.19-alpine as builder

COPY . /api
WORKDIR /api

RUN go mod tidy && \
    GOOS=linux GOARCH=amd64 go build -o engine app/main.go

FROM alpine:latest

ARG ENV
ENV ENV=$ENV

EXPOSE 80

COPY ./env /api/env
WORKDIR /api

COPY --from=builder /api/engine .

ENTRYPOINT ["./engine"]
