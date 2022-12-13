FROM golang:3.7-alpine

RUN apk add build-base

RUN go get github.com/rubenv/sql-migrate/...

WORKDIR /go-migration

ADD . .

RUN go mod download

RUN go get github.com/githubnemo/CompileDaemon

ENTRYPOINT CompileDaemon -command="./go-migrat
ENTRYPOINT CompileDaemon -command="./go-migrationion"