FROM golang:1.12 as builder

COPY . /go/bin/src

WORKDIR /go/bin/src

ENV GO111MODULE=on

RUN go build -o app -mod=vendor



FROM debian:stretch-slim

WORKDIR /go/bin/src

COPY --from=builder /go/bin/src/app /go/bin/app

EXPOSE 8080

ENTRYPOINT ["/go/bin/app"]
