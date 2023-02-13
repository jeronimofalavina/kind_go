FROM golang:1.20.0-alpine3.17 AS builder
RUN apk update && apk add --no-cache git

WORKDIR $GOPATH/src/adjust/http/
COPY . .
RUN GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o /go/bin/app

FROM scratch
COPY --from=builder /go/bin/app /go/bin/app
EXPOSE 8181

CMD ["/go/bin/app"]