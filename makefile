gen_doc:
	swag init -d cmd -o docs/api

.PHONY:gen_doc

ROOT_DIR := ./build

# Compose files relative to the root directory
COMPOSE_FILES := \
	-f $(ROOT_DIR)/project/docker-compose.yaml

# Default target to bring up services
up:
	docker compose $(COMPOSE_FILES) up

# Target to bring down services
down:
	docker compose $(COMPOSE_FILES) down

# Target to view logs
logs:
	docker compose $(COMPOSE_FILES) logs -f

# Target to rebuild and restart services
rebuild:
	docker compose $(COMPOSE_FILES) up  --build

.PHONY: up down logs rebuild
