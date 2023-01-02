SRCS = $(shell git ls-files '*.go')

# Build the tcp server
.PHONY: build-server
build-server:
	go build -o bin/server cmd/server/main.go

# Build the tcp client
.PHONY: build-client
build-client:
	go build -o bin/client cmd/client/main.go

# Run the tcp server
.PHONY: run-server
run-server:
	go run ./cmd/server

# Run the tcp client
.PHONY: run-client
run-client:
	go run ./cmd/client

## Format the Code
.PHONY: fmt
fmt:
	gofmt -s -l -w $(SRCS)
