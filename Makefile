PROJECT = torus
AUTHOR = dethi

GOBUILD = go build -v -i -ldflags "-X main.version=`git rev-parse --short HEAD``date -u +-%Y%m%d.%H%M%S`"

all: build-debug

release: build-prod docker-build docker-push clean

build-debug:
	go-bindata -debug tmpl
	$(GOBUILD)

build-prod:
	go-bindata -nomemcopy -nometadata tmpl
	docker run --rm -v "${PWD}":/go/src/$(PROJECT) -w /go/src/$(PROJECT) \
		golang $(GOBUILD)

docker-build:
	docker build -t $(AUTHOR)/$(PROJECT) .

docker-push:
	docker push $(AUTHOR)/$(PROJECT)

clean:
	go clean
	rm -f *.upx

.PHONY: all release build-dev build-prod docker-build docker-push clean
