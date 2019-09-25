VERSION = 0.1.0
TARGET = semver
SOURCE = $(shell find . -type f -name "*.go" -not -name "*_test.go")

all: build

mod:
	go mod tidy
	go mod vendor

build: mod $(TARGET)

$(TARGET): $(SOURCE)
	go build -o $(TARGET) -ldflags "-X main.version=$(VERSION)" ./cmd/$(TARGET)/...

run: build
	-./$(TARGET) "v0.0.0-rc.0+20190925"

clean:
	-rm $(TARGET)

distclean: clean
	-rm go.sum
	-rm -rf vendor

fmt:
	goimports -w $$(find . -type d -name 'vendor' -prune -o -type f -name '*.go' -print)

test:
	test -z "$$(goimports -l $$(find . -type d -name 'vendor' -prune -o -type f -name '*.go' -print) | tee /dev/stderr)"
	test -z "$$(golint $$(go list ./... | grep -v '/vendor/') | tee /dev/stderr)"
	go test -v ./...

.PHONY: all mod build run clean distclean fmt test
