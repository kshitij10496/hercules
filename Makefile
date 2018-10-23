.PHONY: authors

docker/build:
	docker build -t hercules:latest .

docker/run_attached:
	docker run --rm -p 8080:8080 hercules:latest

authors:
	git log --pretty="%an <%ae>" | sort | uniq >> AUTHORS
