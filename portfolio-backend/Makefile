.PHONY: help up down stop logs ps build generate

help:
	@echo "Available commands:"
	@echo "  make up        - Start Docker containers and run database migrations"
	@echo "  make down      - Stop and remove all containers (including volumes)"
	@echo "  make stop      - Stop containers without removing them"
	@echo "  make logs      - View container logs"
	@echo "  make ps        - Show container status"
	@echo "  make build     - Build the API binary locally"
	@echo "  make generate  - Regenerate GraphQL code from schema"

up:
	@if [ ! -f .env ]; then cp .env.example .env; echo "Created .env from .env.example"; fi
	@echo "Starting Docker containers..."
	docker-compose up -d
	@echo "Waiting for database to be ready..."
	sleep 5
	@echo "Running database migrations..."
	docker-compose exec -T db mysql -h localhost -u root -proot common_db < migrations/001_create_users_table.sql
	@echo "✓ Setup complete! App is running on http://localhost:8080"

down:
	@echo "Stopping containers and removing volumes..."
	docker-compose down -v
	@echo "✓ All containers and volumes removed"

stop:
	@echo "Stopping containers..."
	docker-compose stop
	@echo "✓ Containers stopped"

logs:
	docker-compose logs -f

ps:
	docker-compose ps

build:
	go build -o api ./cmd/api

generate:
	go run github.com/99designs/gqlgen generate
