# Check to see if we can use ash, in Alpine images, or default to BASH.
SHELL_PATH = /bin/ash
SHELL = $(if $(wildcard $(SHELL_PATH)),/bin/ash,/bin/bash)

# ==============================================================================
# CLASS NOTES
#
# Kind
# 	For full Kind v0.19 release notes: https://github.com/kubernetes-sigs/kind/releases/tag/v0.19.0

# ==============================================================================
# Define dependencies

GOLANG          := golang:1.20
ALPINE          := alpine:3.18
KIND            := kindest/node:v1.27.1
POSTGRES        := postgres:15.3
VAULT           := hashicorp/vault:1.13
ZIPKIN          := openzipkin/zipkin:2.24
TELEPRESENCE    := datawire/tel2:2.13.2

KIND_CLUSTER    := ardan-starter-cluster
NAMESPACE       := sales-system
APP             := sales
BASE_IMAGE_NAME := ardanlabs/service
SERVICE_NAME    := sales-api
VERSION         := 0.0.1
SERVICE_IMAGE   := $(BASE_IMAGE_NAME)/$(SERVICE_NAME):$(VERSION)
METRICS_IMAGE   := $(BASE_IMAGE_NAME)/$(SERVICE_NAME)-metrics:$(VERSION)

# VERSION       := "0.0.1-$(shell git rev-parse --short HEAD)"

# ==============================================================================
# Building containers

all: service

service:
	docker build \
		-f zarf/docker/dockerfile.service \
		-t $(SERVICE_IMAGE) \
		--build-arg BUILD_REF=$(VERSION) \
		--build-arg BUILD_DATE=`date -u +"%Y-%m-%dT%H:%M:%SZ"` \
		.

# ==============================================================================
# Running from within k8s/kind

dev-up-local:
	kind create cluster \
		--image $(KIND) \
		--name $(KIND_CLUSTER) \
		--config zarf/k8s/dev/kind-config.yaml

	kubectl.docker wait --timeout=120s --namespace=local-path-storage --for=condition=Available deployment/local-path-provisioner

dev-up: dev-up-local

dev-down-local:
	kind delete cluster --name $(KIND_CLUSTER)

dev-down:
	kind delete cluster --name $(KIND_CLUSTER)

dev-load:
	kind load docker-image $(SERVICE_IMAGE) --name $(KIND_CLUSTER)

dev-apply:
	kustomize build zarf/k8s/dev/sales | kubectl.docker apply -f -
	kubectl.docker wait pods --namespace=$(NAMESPACE) --selector app=$(APP) --for=condition=Ready

# ------------------------------------------------------------------------------

dev-status:
	kubectl.docker get nodes -o wide
	kubectl.docker get svc -o wide
	kubectl.docker get pods -o wide --watch --all-namespaces

dev-restart:
	kubectl.docker rollout restart deployment $(APP) --namespace=$(NAMESPACE)

dev-update: all dev-load dev-restart

dev-update-apply: all dev-load dev-apply

# ------------------------------------------------------------------------------

dev-logs:
	kubectl.docker logs --namespace=$(NAMESPACE) -l app=$(APP) --all-containers=true -f --tail=100 | go run app/tooling/logfmt/main.go -service=$(SERVICE_NAME)

dev-describe-deployment:
	kubectl.docker describe deployment --namespace=$(NAMESPACE) $(APP)

dev-describe-sales:
	kubectl.docker describe pod --namespace=$(NAMESPACE) -l app=$(APP)

# ==============================================================================

run-local:
	go run app/services/sales-api/main.go | go run app/tooling/logfmt/main.go -service=$(SERVICE_NAME)

run-local-help:
	go run app/services/sales-api/main.go --help

tidy:
	go mod tidy
	go mod vendor