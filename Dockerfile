FROM alpine:3.7

COPY bin/gateway /opt/videocoin/bin/gateway

CMD ["/opt/videocoin/bin/gateway"]
