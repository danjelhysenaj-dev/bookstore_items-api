
build:
	@go build -o bin/bookstore_items-api
run: build
	@./bin/bookstore_items-api

test:
	@go test -v ./..