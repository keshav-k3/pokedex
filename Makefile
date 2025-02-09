.PHONY: up down build rebuild logs clean frontend-logs backend-logs

# Start all services
up:
	docker-compose up -d

# Start services and rebuild
build:
	docker-compose up -d --build

# Stop all services
down:
	docker-compose down

# Rebuild and restart all services
rebuild:
	docker-compose down
	docker-compose up -d --build

# View logs from all services
logs:
	docker-compose logs -f

# View frontend logs
frontend-logs:
	docker-compose logs -f frontend

# View backend logs
backend-logs:
	docker-compose logs -f backend

# Clean up docker resources
clean:
	docker-compose down --rmi all --volumes --remove-orphans

# Development shortcuts
dev: up

# Stop everything
stop: down

# Show running containers
ps:
	docker-compose ps

# Show container status and ports
status:
	docker ps

# Follow logs without detached mode
dev-logs:
	docker-compose up