NAME=cobertool
COMMIT=$(shell git rev-parse --short=7 HEAD)
TIMESTAMP:=$(shell date -u '+%Y-%m-%d %I:%M:%S%z')

LDFLAGS += -X "main.BuildTime=${TIMESTAMP}"
LDFLAGS += -X "main.BuildSHA=${COMMIT}"

all: quality test build

quality:
	gofmt -w *.go
	go tool vet *.go

test:
	go test -coverprofile=coverage

build: darwin linux

darwin:
	GOOS=darwin GOARCH=amd64 go build -ldflags '${LDFLAGS}' -o ${NAME}-$@

linux:
	GOOS=linux GOARCH=amd64 go build -ldflags '${LDFLAGS}' -o ${NAME}-$@
