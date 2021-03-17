
VERSION ?= 0.1.0
ARCH ?= amd64
BUILD=`date +%FT%T%z`
NAME=assignment
GO_FLAGS=GOARCH=${ARCH} GOSUMDB=off CGO_ENABLED=0
LD_FLAGS=-ldflags " -w -X main.Name=${NAME} -X main.Version=${VERSION} -X main.Build=${BUILD}"
IMAGE_NAME=assignment

build:
	$(GO_FLAGS) go build ${LD_FLAGS}

test: build
	go test ./... -v

docker: build
	docker build -t ${IMAGE_NAME}:${VERSION} .

fmt:
	go fmt ./...
	go mod tidy
	go mod vendor
