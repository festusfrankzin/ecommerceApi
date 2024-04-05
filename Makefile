build:
    @go build -o bin/ecommerceapi cmd/main.go

test:
    @go test -v ./...

run: build
    @./bin/ecommerceapi

