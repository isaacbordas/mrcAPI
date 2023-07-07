.ONESHELL:

SHELL := /bin/bash
.SHELLFLAGS := -ec

export LOCAL_PKG_PREFIX=mrcAPI
export GOFLAGS=-mod=vendor

BIN=$(CURDIR)/bin
VCS_REF=`git rev-parse --short HEAD`
BUILD_DATE=`date -u +"%Y-%m-%dT%H:%M:%SZ"`

go-wire:
	@echo "generating wire files ..."
	@find $(CURDIR) -name "wire_gen.go" -not -path "$(CURDIR)/vendor/*" -delete
	@go run github.com/google/wire/cmd/wire ./...
	@echo "... done"

go-install-vendor: ## Install dependencies using vendoring
	go mod vendor

go-update-vendor: ## Updates dependencies
	go mod tidy && go mod vendor

go-compile:
	go build \
		-ldflags "-s -w -X main.buildVersion=$(VCS_REF) -X main.buildDate=$(BUILD_DATE)" \
		-o $(BIN)/app_arm \
		./cmd/app

	env GOOS=linux GOARCH=amd64 go build \
		-ldflags "-s -w -X main.buildVersion=$(VCS_REF) -X main.buildDate=$(BUILD_DATE)" \
		-o $(BIN)/app \
		./cmd/app

up: pull-latest-dev ## Start containers
	docker-compose up -d

down: ## Destroy all the containers. This will remove the file system of the containers persisting only the shared volumenes
	docker-compose down

restart: down up ## Restart all the containers. Its invoke down and up

mysql-bash: ## Open an interactive shell console in Go container
	docker-compose exec mysql sh -c "bash"

pull-latest-dev:
	docker-compose pull