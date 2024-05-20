.PHONY: run build
default: run

run: build
	go run ./cmd/api/main.go

build:
	go build ./...