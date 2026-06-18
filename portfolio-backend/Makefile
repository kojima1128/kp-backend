.PHONY: help env up down stop logs clean ps build generate

help:
	@echo "Available commands:"
	@echo "  make env      - Run WSL environment setup script"
	@echo "  make up       - Start Docker containers and run database migrations"
	@echo "  make down     - Stop and remove all containers (including volumes)"
	@echo "  make stop     - Stop containers without removing them"
	@echo "  make logs     - View container logs"
	@echo "  make clean    - Remove all containers and volumes"
	@echo "  make ps       - Show container status"
	@echo "  make build    - Build local binaries"
	@echo "  make generate - Run gqlgen"

env:
	bash setup-wsl.sh

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
	go build ./cmd/api && go build ./cmd/worker

generate:
	go generate ./...
