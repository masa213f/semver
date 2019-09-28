VERSION = 0.1.0
TARGET = semver
SOURCE = $(shell find . -type f -name "*.go" -not -name "*_test.go")

all: build

setup:
	go get -u golang.org/x/tools/cmd/goimports
	go get -u golang.org/x/lint/golint

mod:
	go mod tidy
	go mod vendor

build: mod $(TARGET)

$(TARGET): $(SOURCE)
	go build -o $(TARGET) -ldflags "-X main.version=$(VERSION)" ./cmd/$(TARGET)/...

run: $(TARGET)
	# Run the example commands in README.md.
	@echo
	./$(TARGET) v1.2.3-rc.0+build.20190925        ; echo "# => exit status: $$?"
	@echo
	./$(TARGET) v1.2.3-rc.0+build.20190925 --json ; echo "# => exit status: $$?"
	@echo
	./semver v1.12                                ; echo "# => exit status: $$?"
	@echo
	./semver v1.01.0                              ; echo "# => exit status: $$?"
	@echo
	./semver -p v1.1.2-rc.0                       ; echo "# => exit status: $$?"
	@echo
	./semver -p v1.1.2                            ; echo "# => exit status: $$?"

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

.PHONY: all setup mod build run clean distclean fmt test
