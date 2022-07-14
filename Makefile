export GO111MODULE=on

.PHONY: all
all: gofmt lint test

.PHONY: test
test:
	go test ./... -cover

.PHONY: gofmt
gofmt:
	@diff=$$(gofmt -s -d ./); \
	if [ -n "$$diff" ]; then \
		echo "Please run 'make fmt' and commit the result:"; \
		echo "$${diff}"; \
		exit 1; \
	fi;
