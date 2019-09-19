SERVICE := backend
PACKAGE := github.com/igor-karpukhin/user-api-test
VERSION := ${shell git describe --tags --always}
BUILDTIME := ${shell date -u '+%Y-%m-%d_%H:%M:%S'}
LDFLAGS := -s -w -X '${PACKAGE}/pkg/version.Version=${VERSION}' \
					-X '${PACKAGE}/pkg/version.BuildTime=${BUILDTIME}'
ifdef OSX
	TARGET_OS=darin
else
	TARGET_OS=linux
endif

all: clean build test docker

dep:
	dep ensure -v --vendor-only

build:
	mkdir -p bin/
	CGO_ENABLED=0 GOOS=$(TARGET_OS) go build -ldflags "${LDFLAGS}" -o bin/$(SERVICE) $(PACKAGE)/cmd/backend 

test:
	go test ./...

clean:
	rm -rf bin/

docker:
	docker build . -t backend:$(VERSION)
	docker build scripts/pg -t backend-db:$(VERSION)
