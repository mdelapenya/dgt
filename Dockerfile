FROM golang:alpine

COPY . $GOPATH/src/github.com/mdelapenya/dgt/
WORKDIR $GOPATH/src/github.com/mdelapenya/dgt

RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -ldflags="-w -s" -o /go/bin/dgt

EXPOSE 8080

ENTRYPOINT ["/go/bin/dgt"]