all: dev build

release: prod build push

prod:
	GOOS=linux GOARCH=amd64 go build -v -a -ldflags '-s -w'
	upx torrent_service

dev:
	GOOS=linux GOARCH=amd64 go build

build:
	docker build -t dethi/torrent_service .

push:
	docker push dethi/torrent_service

.PHONY: all release prod dev build push
