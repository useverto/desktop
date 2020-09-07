web:
	cd verto && yarn build && yarn export

run:
	statik -src=./verto/__sapper__/export
	go run desktop.go

build:
	statik -src=./verto/__sapper__/export
	go build desktop.go
