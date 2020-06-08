include .env

GOBASE=$(shell pwd)
MAKEFLAGS += --silent
TAG="runlevl4/selfdemo:$(BUILD_TAG)"

build:
	echo "***** building mac binary *****"
	GOOS=darwin GOARCH=amd64 go build -o bin/mac/selfdemo api/cmd/main.go
	echo "***** build complete *****"

build-linux:
	echo "***** building linux binary *****"
	GOOS=linux GOARCH=amd64 go build -o bin/linux/selfdemo api/cmd/main.go
	echo "***** build complete *****"

build-all: build-mac build-linux

docker-build:	
	docker build -t $(TAG) .

docker-push:
	docker push $(TAG)


