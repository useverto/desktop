web:
	cd verto && yarn build && yarn export

run:
	go run verto.go

build:
	statik -src=./verto/__sapper__/export
	go build desktop.go