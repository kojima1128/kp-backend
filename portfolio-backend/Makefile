.PHONY: help env up down stop logs clean ps build generate_gql

help:
	@echo "Available commands:"
	@echo "  make env           - Setup environment (copy .env-sample to .env)"
	@echo "  make up            - Start Docker containers and run database migrations"
	@echo "  make down          - Stop and remove all containers (including volumes)"
	@echo "  make stop          - Stop containers without removing them"
	@echo "  make logs          - View container logs"
	@echo "  make clean         - Remove all containers and volumes"
	@echo "  make ps            - Show container status"
	@echo "  make build         - Build local binaries"
	@echo "  make generate_gql  - Run gqlgen for GraphQL code generation"

env:
	@[ -f .env ] || cp .env-sample .env
	@echo "✓ .env file created from .env-sample"

up: env
	docker-compose up -d --build

down:
	docker-compose down -v

stop:
	docker-compose stop

logs:
	docker-compose logs -f

clean: down

ps:
	docker-compose ps

build:
	cd portfolio-backend && go build ./cmd/api && go build ./cmd/worker

generate_gql:
	cd portfolio-backend && go generate ./...
