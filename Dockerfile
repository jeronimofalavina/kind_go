FROM golang:1.20.0-alpine3.17 AS builder

WORKDIR $GOPATH/src/http/
COPY go/* ./
RUN GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o /go/bin/hostname-app

FROM scratch
COPY --from=builder /go/bin/hostname-app /go/bin/hostname-app
EXPOSE 8181

CMD ["/go/bin/hostname-app"]
