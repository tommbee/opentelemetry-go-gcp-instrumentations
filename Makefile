Version := $(shell git describe --tags --dirty)
#GitCommit := $(shell git rev-parse HEAD)
#LDFLAGS := "-s -w -X main.Version=$(Version) -X main.GitCommit=$(GitCommit)"

export GO111MODULE=on

.PHONY: all
all: gofmt test

.PHONY: test
test:
	find . -name go.mod -execdir go test ./... -cover \;

.PHONY: gofmt
gofmt:
	@test -z $(shell gofmt -l -s $(SOURCE_DIRS) ./ | tee /dev/stderr) || (echo "[WARN] Fix formatting issues with 'make fmt'" && exit 1)
