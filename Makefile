.DEFAULT_GOAL := help

.PHONY: build
build: ## build
	@docker compose -f deploy/compose.yaml -p swarmyard build

.PHONY: up
up: ## start up containers
	@docker compose -f deploy/compose.yaml -p swarmyard up

.PHONY: build-up
build-up: ## build and start up containers
	@docker compose -f deploy/compose.yaml -p swarmyard up --build

.PHONY: down
down: ## stop and remove containers
	@docker compose -f deploy/compose.yaml -p swarmyard down

.PHONY: help
help: ## show help
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | \
		awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'
