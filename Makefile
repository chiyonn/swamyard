.DEFAULT_GOAL := help

ENV ?= dev
COMPOSE := docker compose -f deploy/compose.$(ENV).yaml -p swarmyard

.PHONY: build
build: ## build containers
	@$(COMPOSE) build

.PHONY: up
up: ## start up containers
	@$(COMPOSE) up

.PHONY: build-up
build-up: ## build and start up containers
	@$(COMPOSE) up --build

.PHONY: down
down: ## stop and remove containers
	@$(COMPOSE) down

.PHONY: help
help: ## show help
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | \
		awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'
