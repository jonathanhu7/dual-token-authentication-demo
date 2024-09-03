build:
	@go build -o bin/server cmd/server/main.go
	@chmod +x bin/server

run: build
	@./bin/server

.PHONY: build run