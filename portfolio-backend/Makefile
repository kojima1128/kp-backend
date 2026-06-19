.PHONY: help up down stop logs clean ps build generate_gql

help:
	@echo "Available commands:"
	@echo "  make up            - Start Docker containers and run database migrations"
	@echo "  make down          - Stop and remove all containers (including volumes)"
	@echo "  make stop          - Stop containers without removing them"
	@echo "  make logs          - View container logs"
	@echo "  make clean         - Remove all containers and volumes"
	@echo "  make ps            - Show container status"
	@echo "  make build         - Build local binaries"
	@echo "  make generate_gql  - Run gqlgen for GraphQL code generation"

up:
	@[ -f .env ] || cp .env.example .env
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
