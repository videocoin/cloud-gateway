GOOS?=linux
GOARCH?=amd64

NAME=gateway
GCP_PROJECT=videocoin-network
VERSION=$$(git describe --abbrev=0)-$$(git rev-parse --short HEAD)

.PHONY: deploy

default: build

version:
	@echo ${VERSION}

build:
	GOOS=${GOOS} GOARCH=${GOARCH} \
		go build \
			-ldflags="-w -s -X main.Version=${VERSION}" \
			-o bin/${NAME} \
			./cmd/main.go

deps:
	env GO111MODULE=on go mod vendor

docker-build:
	docker build -t gcr.io/${GCP_PROJECT}/${NAME}:${VERSION} -f Dockerfile .

docker-push:
	gcloud docker -- push gcr.io/${GCP_PROJECT}/${NAME}:${VERSION}

release: docker-build docker-push

deploy:
	ENV=${ENV} deploy/deploy.sh