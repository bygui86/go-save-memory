
# VARIABLES
# -


# CONFIG
.PHONY: help print-variables
.DEFAULT_GOAL := help


# ACTIONS

## http-server

run-http-server :		## Start HTTP server
	export GO111MODULE=on && \
	godotenv -f local.env go run ./http-server/main.go

build-http-server :		## Build HTTP server
	export GO111MODULE=on && \
	go build ./http-server

## cli-dashboard

run-cli-dashboard :		## Start CLI dashboard
	export GO111MODULE=on && \
	go run ./cli-dashboard/main.go

build-cli-dashboard :		## Build CLI dashboard
	export GO111MODULE=on && \
	go build ./cli-dashboard

## helpers

help :		## Help
	@echo ""
	@echo "*** \033[33mMakefile help\033[0m ***"
	@echo ""
	@echo "Targets list:"
	@grep -E '^[a-zA-Z_-]+ :.*?## .*$$' $(MAKEFILE_LIST) | sort -k 1,1 | awk 'BEGIN {FS = ":.*?## "}; {printf "\t\033[36m%-30s\033[0m %s\n", $$1, $$2}'
	@echo ""

print-variables :		## Print variables values
	@echo ""
	@echo "*** \033[33mMakefile variables\033[0m ***"
	@echo ""
	@echo "- - - makefile - - -"
	@echo "MAKE: $(MAKE)"
	@echo "MAKEFILES: $(MAKEFILES)"
	@echo "MAKEFILE_LIST: $(MAKEFILE_LIST)"
	@echo "- - -"
	@echo "NAME: $(NAME)"
	@echo "CONTAINER_NAME: $(CONTAINER_NAME)"
	@echo "CONTAINER_VERSION: $(CONTAINER_VERSION)"
	@echo "CONTAINER_PORTS: $(CONTAINER_PORTS)"
	@echo "LOCAL_CONTAINER_TAG: $(LOCAL_CONTAINER_TAG)"
	@echo "REMOTE_REGISTRY_PREFIX: $(REMOTE_REGISTRY_PREFIX)"
	@echo "REMOTE_CONTAINER_TAG: $(REMOTE_CONTAINER_TAG)"
	@echo ""
