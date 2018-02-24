PACKAGES = $(shell \
	find . '(' -path '*/.*' -o -path './vendor' ')' -prune \
	-o -name '??*' -type d -print)


test:
	go test  ${PACKAGES}

.PHONY: test
