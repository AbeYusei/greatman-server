# 適当に設定を書く
# NAME     := hello
# VERSION  := v1
# ...

all: setup build

setup:
	go get -u -d -v ./...
	go install -v ./...

build:
	go build -o app -i api/main.go

# docker build scripts

docker-dev:
	docker build --target dev -f docker/api/Dockerfile -t greatmen-server:dev .

docker-build:
	docker build --target build -f docker/api/Dockerfile -t greatmen-server:build .
