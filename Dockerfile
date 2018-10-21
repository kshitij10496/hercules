FROM golang:1.11 as builder
WORKDIR /go/src/github.com/kshitij10496/hercules/
ADD . /go/src/github.com/kshitij10496/hercules
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o hercules .

FROM alpine:latest  
RUN apk --no-cache add ca-certificates
ENV PORT=8080
ENV HERCULES_DATABASE="user=kshitij10496 dbname=herculesdb sslmode=disable"
WORKDIR /root/
COPY --from=builder /go/src/github.com/kshitij10496/hercules/hercules .
CMD ["./hercules"]
EXPOSE 8080