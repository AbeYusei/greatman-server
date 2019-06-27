# dev stage

ARG flavor=alpine

FROM golang:1.12.6-$flavor AS dev

ENV GO111MODULE on

WORKDIR /go/src/app
COPY go.mod .
COPY Makefile .
COPY src .

RUN apk update \
  && apk add --no-cache git curl gcc make
RUN go mod tidy \
  && go get -u github.com/codegangsta/gin \
  && make all

CMD ["gin", "-t", "api", "-p", "8080"]

# build stage

FROM $flavor AS build

WORKDIR /go/src/app

COPY --from=dev /go/src/app/app ./app

CMD ["./app"]
