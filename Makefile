.PHONY: authors

GOPACKAGES = $(shell go list ./...  | grep -v /vendor/)

docker/build:
	docker build -t hercules:latest .

docker/run_attached:
	docker run --rm -p 8080:8080 hercules:latest

build:
	go build -o hercules

test:
	go test -v $(GOPACKAGES)

authors:
	rm AUTHORS
	git log --pretty="%an <%ae>" | sort | uniq >> AUTHORS
