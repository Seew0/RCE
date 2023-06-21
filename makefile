build:
	@go build -o ./bin/rce
run: build
	@./bin/rce
