last_revision := $(shell git rev-list --tags --max-count=1)
last_tag := $(shell git describe --tags $(last_revision))

.PHONY: test
test:
	@go test ./...

.PHONY: build
build:
	@go run -ldflags "-X main.version=$(last_tag)" -o publish/gelp main.go

