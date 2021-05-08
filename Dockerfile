FROM golang:1.15-alpine3.12 AS builder

RUN go version

COPY . /github.com/ferestgo/tg-balancer/
WORKDIR /github.com/ferestgo/tg-balancer/

RUN go mod download
RUN GOOS=linux go build -o ./.bin/bot ./cmd/bot/main.go

FROM alpine:latest

WORKDIR /root/

COPY --from=0 /github.com/ferestgo/tg-balancer/.bin/bot .
COPY --from=0 /github.com/ferestgo/tg-balancer/configs configs/

EXPOSE 80

CMD ["./bot"]