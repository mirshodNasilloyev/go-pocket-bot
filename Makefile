.PHOHY:
.SILENT:


build:
	go build -o ./.bin/bot cmd/bot/main.go
run: build
	./.bin/bot

build-image:
	docker build -t tg-bot-youtube-go:v0.1 .

start-container:
	docker run --name telegram-bot -p 80:80 --env-file .env tg-bot-youtube-go:v0.1
