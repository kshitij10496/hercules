FROM golang:1.11

ENV PORT=8080
ENV DATABASE_URL="user=kshitij10496 dbname=herculesdb sslmode=disable"

RUN mkdir -p /go/src/github.com/kshitij10496/hercules
WORKDIR /go/src/github.com/kshitij10496/hercules
ADD . /go/src/github.com/kshitij10496/hercules

RUN go build .

CMD ["./hercules"]

EXPOSE 8080
