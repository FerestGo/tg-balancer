.PHONY:

build:
	go build -o ./.bin/bot cmd/bot/main.go

run: build
	./.bin/bot

start:
	go build -o ./.bin/bot cmd/bot/main.go && ./.bin/bot

build-image:
	docker build -t zhashkevych/pocketer:0.1 .

start-container:
	docker run --env-file .env -p 80:80 zhashkevych/pocketer:0.1