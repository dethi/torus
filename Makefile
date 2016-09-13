release: prod build push

prod:
	go-bindata -nomemcopy -nometadata tmpl
	#GOOS=linux GOARCH=amd64 go build -i
	docker run --rm -v "$PWD":/go/src/torus -w /go/src/torus golang go build -v
	upx torus

build:
	docker build -t dethi/torus .

push:
	docker push dethi/torus

clean:
	go clean
	rm -f *.upx

.PHONY: release prod build push clean
