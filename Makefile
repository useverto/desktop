web:
	cd verto && yarn build && yarn export

run:
	go run ./fs/embed.go -src=./verto/__sapper__/export
	go run desktop.go

build:
	go run ./fs/embed.go -src=./verto/__sapper__/export
	go build desktop.go
