FROM golang:1.12 as builder

WORKDIR /go/src/be
COPY . /go/src/be

ENV GO111MODULE=on

RUN go build -mod=vendor

FROM debian:stretch-slim
COPY --from=builder /go/src/be /usr/bin/be
ENTRYPOINT ["/usr/bin/be"]
EXPOSE 8080
