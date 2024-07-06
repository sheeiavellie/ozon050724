run: build
	@./bin/ozon050724
build:
	@go build -o bin/ozon050724 server.go
