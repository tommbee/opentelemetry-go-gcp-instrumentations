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

.PHONY: lint
lint:
	@hash golint > /dev/null 2>&1; if [ $$? -ne 0 ]; then \
		$(GO) get -u golang.org/x/lint/golint; \
	fi
	for PKG in $(PACKAGES); do golint -set_exit_status $$PKG || exit 1; done;
