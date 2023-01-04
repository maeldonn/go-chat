SRCS = $(shell git ls-files '*.go')

# Build the tcp server
.PHONY: build
build:
	go build -o bin/tcpchat main.go

# Run the tcp server
.PHONY: run
run: build
	./bin/tcpchat

# Run the tcp client
.PHONY: dev
dev:
	go run main.go

## Format the Code
.PHONY: fmt
fmt:
	gofmt -s -l -w $(SRCS)
