FROM golang:1.12.4 as builder
WORKDIR /go/src/github.com/videocoin/cloud-gateway
COPY . .
RUN make build

FROM bitnami/minideb:jessie
COPY --from=builder /go/src/github.com/videocoin/cloud-gateway/bin/gateway /opt/videocoin/bin/gateway
CMD ["/opt/videocoin/bin/gateway"]
