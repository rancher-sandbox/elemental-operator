GIT_COMMIT?=$(shell git rev-parse HEAD)
GIT_COMMIT_SHORT?=$(shell git rev-parse --short HEAD)
GIT_TAG?=$(shell git describe --abbrev=0 --tags 2>/dev/null || echo "v0.0.0" )
TAG?=${GIT_TAG}
REPO?=rancher/elemental-operator
export ROOT_DIR:=$(shell dirname $(realpath $(lastword $(MAKEFILE_LIST))))
CHART?=$(shell find $(ROOT_DIR) -type f  -name "elemental-operator*.tgz" -print)
CHART_VERSION?=$(subst v,,$(GIT_TAG))

.PHONY: build
build: operator installer

.PHONY: operator
operator:
	CGO_ENABLED=0 go build -ldflags "-extldflags -static -s -X 'github.com/rancher/elemental-operator/version.Version=$TAG'" -o build/elemental-operator $(ROOT_DIR)/cmd/operator

.PHONY: installer
installer:
	CGO_ENABLED=0 go build -ldflags "-extldflags -static -s -X 'github.com/rancher/elemental-operator/version.Version=$TAG'" -o build/elemental-installer $(ROOT_DIR)/cmd/installer

.PHONY: build-docker
build-docker:
	DOCKER_BUILDKIT=1 docker build \
		-f package/Dockerfile \
		--target elemental-operator \
		-t ${REPO}:${TAG} .

.PHONY: build-docker-push
build-docker-push: build-docker
	docker push ${REPO}:${TAG}

.PHONY: chart
chart:
	mkdir -p  $(ROOT_DIR)/build
	helm package --version ${CHART_VERSION} --app-version ${GIT_TAG} -d $(ROOT_DIR)/build/ $(ROOT_DIR)/chart

validate:
	scripts/validate

unit-tests-deps:
	go install -mod=mod github.com/onsi/ginkgo/v2/ginkgo@latest
	go get github.com/onsi/gomega/...

unit-tests:
	ginkgo -r -v --covermode=atomic --coverprofile=coverage.out ./pkg/...
