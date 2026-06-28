.PHONY: api cli build test docker-up docker-down fmt lint

api:
	go run ./cmd/api

cli:
	go run ./cmd/cli

build:
	go build -o bin/devplatform-api ./cmd/api
	go build -o bin/devplatform ./cmd/cli

test:
	go test ./...

fmt:
	go fmt ./...

lint:
	golangci-lint run

docker-up:
	docker compose up -d

docker-down:
	docker compose down