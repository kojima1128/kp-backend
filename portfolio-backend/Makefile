.PHONY: help up down stop

help:
	@echo "Available commands:"
	@echo "  make up       - Start Docker containers and run database migrations"
	@echo "  make down     - Stop and remove all containers (including volumes)"
	@echo "  make stop     - Stop containers without removing them"
	@echo "  make logs     - View container logs"
	@echo "  make clean    - Remove all containers and volumes"

up:
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

clean: down
	@echo "Cleanup complete"

ps:
	docker-compose ps
