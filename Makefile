last_revision := $(shell git rev-list --tags --max-count=1)
last_tag := $(shell git describe --tags $(last_revision))

.PHONY: test
test:
	@go test ./...

.PHONY: build
build:
	@GOOS=linux GOARCH=amd64 go build -ldflags "-s -w -X 'github.com/imballinst/gelp/src/helpers.Version=$(last_tag)'" -o publish/gelp-linux-amd64 main.go
	@chmod +x publish/gelp-linux-amd64
	@upx -1 -q publish/gelp-linux-amd64
