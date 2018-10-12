FROM golang:1.11

ENV PORT=8080
ENV DATABASE_URL="user=kshitij10496 dbname=herculesdb sslmode=disable"

RUN mkdir -p /go/src/github.com/kshitij10496/hercules
WORKDIR /go/src/github.com/kshitij10496/hercules
ADD . /go/src/github.com/kshitij10496/hercules

RUN go build .

CMD ["./hercules"]

EXPOSE 8080
D55C222D8C84172B36936948A64867DE.worker1