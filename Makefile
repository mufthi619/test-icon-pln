GO_VERSION := 1.20
GOBIN := $(shell go env GOPATH)/bin

run:
	go mod tidy
	go mod vendor
	go run ./cmd/api/main.go
test:
	go test -v -coverprofile=coverage.out ./...
clean:
	rm -rf coverage.out

docker-build:
	docker build -t mufthiryanda/go-icon-pln:latest .
docker-run:
	docker run -p 8001:8001 mufthiryanda/go-icon-pln:latest

.PHONY: run test clean