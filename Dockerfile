FROM golang:1.15-alpine3.12 AS builder

RUN go version

RUN go env -w GOPRIVATE=github.com/ereshzealous

COPY . /github.com/ferestgo/tg-balancer/
WORKDIR /github.com/ferestgo/tg-balancer/

RUN go mod download
export token=${token}
RUN GOOS=linux go build -o ./.bin/bot ./cmd/bot/main.go
# RUN git config --global url."https://golang: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@github.com".insteadOf "https://github.com"


FROM alpine:latest

WORKDIR /root/

COPY --from=0 /github.com/ferestgo/tg-balancer/.bin/bot .
COPY --from=0 /github.com/ferestgo/tg-balancer/configs configs/

EXPOSE 80

CMD ["./bot"]