gen_doc:
	swag init -d cmd -o docs/api

.PHONY: gen_doc

ROOT_DIR := ./build

# Compose files relative to the root directory
COMPOSE_FILES := \
	-f $(ROOT_DIR)/postgres/docker-compose.yaml \
	-f $(ROOT_DIR)/project/docker-compose.yaml

NETWORK_NAME=golipors-network

ensure-network:
	@echo "Ensuring project network is exists..."
	@if [ -z "$$(docker network ls --filter name=$(NETWORK_NAME) --format '{{.Name}}')" ]; then \
		echo "Network $(NETWORK_NAME) does not exist. Creating..."; \
		docker network create $(NETWORK_NAME); \
	fi

go-mod-vendor:
	@echo "Running 'go mod vendor' to sync dependencies..."
	@go mod vendor

# Default target to bring up services
up: ensure-network go-mod-vendor
	docker compose --project-directory $(ROOT_DIR) $(COMPOSE_FILES) up

# Target to bring down services
down: ensure-network go-mod-vendor
	docker compose --project-directory $(ROOT_DIR) $(COMPOSE_FILES) down

CONTAINER ?= all
FLAGS ?=
build: ensure-network go-mod-vendor
	docker compose --project-directory $(ROOT_DIR) $(COMPOSE_FILES) build $(FLAGS) $(CONTAINER)

# Target to view logs
logs: ensure-network go-mod-vendor
	docker compose --project-directory $(ROOT_DIR) $(COMPOSE_FILES) logs -f

# Target to rebuild and restart services
rebuild: ensure-network go-mod-vendor
	docker compose --project-directory $(ROOT_DIR) $(COMPOSE_FILES) up  --build

.PHONY: up down logs rebuild
