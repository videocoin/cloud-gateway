FROM golang:1.14 as builder

WORKDIR /go/src/github.com/videocoin/cloud-gateway
COPY . .

RUN make build


FROM bitnami/minideb:jessie

COPY --from=builder /go/src/github.com/videocoin/cloud-gateway/bin/gateway /gateway

CMD ["/gateway"]
