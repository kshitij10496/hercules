FROM golang:1.11-alpine

RUN apk --no-cache add ca-certificates

# git is a prereq of go get
RUN apk --update add git

# get "fresh" - rebuilds the server whenever files change
RUN go get github.com/pilu/fresh

WORKDIR /go/src/github.com/kshitij10496/hercules

CMD ["fresh"]
