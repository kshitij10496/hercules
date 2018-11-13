.PHONY: authors

GOPACKAGES = $(shell go list ./...  | grep -v /vendor/)

build:
	go build -o hercules

test:
	go test -v $(GOPACKAGES)

authors:
	rm AUTHORS
	git log --pretty="%an <%ae>" | sort | uniq >> AUTHORS
