FROM golang:1.19-alpine as builder

RUN mkdir -p /lambda-cicd/api
COPY ["./go.mod", "./go.sum", "/lambda-cicd/"]
COPY "./api" "/lambda-cicd/api/"

WORKDIR /lambda-cicd

RUN go mod tidy && \
    go build api/app/main.go

FROM alpine:latest

ARG ENV
ENV ENV=$ENV
RUN mkdir -p /lambda-cicd/api/env

COPY --from=builder /lambda-cicd/main /lambda-cicd
COPY --from=builder /lambda-cicd/api/env/ /lambda-cicd/api/env/

WORKDIR /lambda-cicd

CMD ["./main"]
