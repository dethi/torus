AUTHOR = dethi
PROJECT = torus
NAME := ${AUTHOR}/${PROJECT}

VERSION = 0.1.0
BUILD := $(shell git rev-parse --short HEAD)
BUILD_TIME := $(shell date -u +%Y/%m/%d-%H:%M:%S)

LDFLAGS := -ldflags "-X github.com/${NAME}.Version=${VERSION} \
	-X github.com/${NAME}.Build=${BUILD} \
	-X github.com/${NAME}.BuildTime=${BUILD_TIME}"

GOBUILD := go build -v -i ${LDFLAGS} ./cmd/torus
GOPACKAGE := $(shell glide nv)

all: build-debug

release: build-prod docker-build docker-push clean

web/template/files.go: web/template/files/*
	staticfiles -o $@ web/template/files

build-static: web/template/files.go

build-debug: build-static
	@${GOBUILD}

build-prod:
	docker run --rm -v "${PWD}":/go/src/github.com/${NAME} \
		-w /go/src/github.com/${NAME} \
		golang ${GOBUILD}

docker-build:
	docker build -t ${NAME} .

docker-push:
	docker push ${NAME}

vet:
	@go vet ${GOPACKAGE}

test:
	@go test $(GOPACKAGE)

clean:
	rm -f ${PROJECT} *.upx

.PHONY: all release build-static build-dev build-prod \
	docker-build docker-push vet test clean
