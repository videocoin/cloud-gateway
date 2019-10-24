module github.com/videocoin/cloud-gateway

go 1.12

require (
	github.com/gogo/gateway v1.1.0
	github.com/grpc-ecosystem/grpc-gateway v1.11.3
	github.com/kelseyhightower/envconfig v1.4.0
	github.com/labstack/echo v3.3.10+incompatible
	github.com/prometheus/client_golang v0.9.4
	github.com/sirupsen/logrus v1.4.2
	github.com/videocoin/cloud-api v0.2.14
	github.com/videocoin/cloud-pkg v0.0.6
	google.golang.org/grpc v1.21.1
)

replace github.com/videocoin/cloud-api => ../cloud-api

replace github.com/videocoin/cloud-pkg => ../cloud-pkg
