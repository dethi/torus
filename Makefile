release: prod build push

prod:
	go-bindata -nomemcopy -nometadata tmpl
	GOOS=linux GOARCH=amd64 go build -v -a -ldflags '-s -w'
	upx torrent_service

build:
	docker build -t dethi/torrent_service .

push:
	docker push dethi/torrent_service

.PHONY: release prod build push
