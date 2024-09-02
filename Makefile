run: build
	@./bin/go-in-memory-store

build:
	@go build -o bin/go-in-memory-store .
