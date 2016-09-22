PROJECT = torus
AUTHOR = dethi

GOBINDATA = go-bindata -pkg $(PROJECT)
GOBUILD = go build -v -i -ldflags "-X main.version=`git rev-parse --short HEAD``date -u +-%Y%m%d.%H%M%S`" ./cmd/...

all: build-debug

release: build-prod docker-build docker-push clean

build-debug:
	$(GOBINDATA) -debug tmpl
	$(GOBUILD)

build-prod:
	$(GOBINDATA) -nomemcopy -nometadata tmpl
	docker run --rm -v "${PWD}":/go/src/github.com/$(AUTHOR)/$(PROJECT) \
		-w /go/src/github.com/$(AUTHOR)/$(PROJECT) \
		golang $(GOBUILD)

docker-build:
	docker build -t $(AUTHOR)/$(PROJECT) .

docker-push:
	docker push $(AUTHOR)/$(PROJECT)

clean:
	rm -f $(PROJECT) *.upx

.PHONY: all release build-dev build-prod docker-build docker-push clean
