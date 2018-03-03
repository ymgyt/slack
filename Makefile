PACKAGES = $(shell \
	find . '(' -path '*/.*' -o -path './vendor' ')' -prune \
	-o -name '??*' -type d -print)

.PHONY: test
test:
	go test  ${PACKAGES}

.PHONY: dependencies
dependencies:
	dep version || go get -u github.com/golang/dep/cmd/dep
	dep ensure
