VERSION:=$(shell git describe --tags --always)
DIRECTORIES=go list -f '{{.Dir}}' ./... | grep -v /vendor/
PACKAGES=go list ./... | grep -v /vendor/

.PHONY: build
build: clean
	mkdir -p bin/ && go build -ldflags "-X main.Version=${VERSION}" -o ./bin/ `$(call DIRECTORIES)`

.PHONY: checks
checks: format vet

.PHONY: clean
clean:
	go clean
	rm -rf ./bin

.PHONY: deps
deps:
	go mod tidy
	go mod verify
	go mod vendor

.PHONY: format
format:
	gofmt -w `$(call DIRECTORIES)`

.PHONY: generate
generate:
	find . -name "wire_gen.go" -type f | xargs rm
	go generate `$(call DIRECTORIES)`

.PHONY: test
test:
	go test -v -mod=vendor `$(call PACKAGES)`

.PHONY: test-race
test-race:
	go test -v -race -mod=vendor `$(call PACKAGES)`

.PHONY: test-api
test-api:
	go test -p 1 -count=1 -cpu=1 -v -mod=vendor -tags=api_test ./test/...

.PHONY: vet
vet:
	go vet -mod=vendor `$(call PACKAGES)`