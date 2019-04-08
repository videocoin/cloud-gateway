GOOS?=linux
GOARCH?=amd64

NAME=vc-gateway
VERSION=$$(git describe --abbrev=0)-$$(git rev-parse --short HEAD)

CHARTS_BUCKET=videocoin-cloud-charts
GCP_PROJECT=videocoin

version:
	@echo ${VERSION}

build:
	GOOS=${GOOS} GOARCH=${GOARCH} \
		go build \
			-ldflags="-w -s -X main.Version=${VERSION}" \
			-o bin/${NAME} \
			./cmd/main.go

docker-build:
	docker build -t gcr.io/${GCP_PROJECT}/${NAME}:${VERSION} -f Dockerfile .

docker-push:
	gcloud docker -- push gcr.io/${GCP_PROJECT}/${NAME}:${VERSION}

helm-package:
	@echo "Packaging ${NAME}..."
	@helm package --save=false -d helm/ helm/charts/${NAME}

helm-repo-index:
	@echo "Indexing charts repository..."
	@helm repo index helm/repo --url https://${CHARTS_BUCKET}.storage.googleapis.com

helm-repo-sync:
	@echo "Syncing repo..."
	@gsutil -m -h "Cache-Control:public,max-age=0" cp -a public-read helm/repo/* gs://${CHARTS_BUCKET}

helm-repo-update: helm-package helm-repo-index helm-repo-sync

release: build docker-build docker-push
