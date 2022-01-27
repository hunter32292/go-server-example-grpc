GIT_HASH = $(shell git describe --tags --dirty --always)

.PHONY: all run lint fmt test build clean help

help: ## Show help messages
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

all: clean test build run ## clean dev env, run all linting and tests, build binary and run server

run: ## Run the project
	./bin/project
	
test: lint fmt compile-proto unit ## Execute all linting and tests 

unit: ## Execute unit tests 
	go test ./... -coverprofile cover.out
	
lint: ## Lint go codebase
	go vet ./...

fmt: ## Format go codebase
	go fmt ./...

compile-proto: ## Compile GRPC files into models
	protoc --go_out=plugins=grpc,paths=source_relative:. protos/**/*.proto 

build: test ## Build project binary
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 GO111MODULE=on GOSUMDB=off go build -o bin/project main.go

docker: ## Build project dockerfile
	docker build . -f image/Dockerfile

docker-dev: ## Build project development dockerfile $ docker run --entrypoint=/bin/sh TAG
	docker build . -f image/debug/Dockerfile	
	
clean: ## Clean the development environment
	rm -rf bin cover.out *.log *.db
