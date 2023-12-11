.DEFAULT_GOAL := all

clean:
	go clean
	rm -fr bin/

build:
	go build ./...

test:
	go test ./...

format:
	go fmt ./...

lint:
	golangci-lint run

update:
	go get -u -t ./...
	go mod tidy

verify: lint test

bin/day01:
	go build -o bin/day01 day01/cmd/cmd.go

bin/day02:
	go build -o bin/day02 day02/cmd/cmd.go

all: clean verify build bin/day01 bin/day02
