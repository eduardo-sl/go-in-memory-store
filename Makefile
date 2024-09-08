run: build
	@./bin/go-in-memory-store --listenAddr :5001

build:
	@go build -o bin/go-in-memory-store .
