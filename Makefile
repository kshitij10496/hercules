.PHONY: authors

GOPACKAGES = $(shell go list ./...  | grep -v /vendor/)

build:
	go build -o hercules

test:
	go test -v $(GOPACKAGES)

authors:
	rm AUTHORS
	git log --pretty="%an <%ae>" | sort | uniq >> AUTHORS

prepare:
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o hercules .
	mv hercules dev-env/app/
