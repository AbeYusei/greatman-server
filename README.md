# A Greatmen Web Server

require Go Version 1.12 or higher.

## A Greatmen tells you what is the great

useage

`gin -t src/api -p 8080`

## build

`make all`

## docker

`make docker-dev && docker container run -it --rm -p 8080:8080 -p 3001:3001 greatman-server:dev`
