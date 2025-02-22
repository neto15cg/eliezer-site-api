.PHONY: build run docker-build docker-run docker-compose clean dev dev-db dev-db-status dev-clean dev-status

build:
	go build -o main .

run:
	go run main.go

docker-build:
	docker build -t go-app .

docker-run:
	docker run -p 8080:8080 go-app

docker-compose:
	docker compose up --build

clean:
	rm -f main
	docker compose down -v

# Clean only application files, preserve database
dev-clean:
	docker compose stop app
	docker compose rm -f app
	rm -rf tmp/

# Start PostgreSQL container and wait for it to be ready
dev-db:
	docker compose up -d db
	@echo "Waiting for database to be ready..."
	@for i in $$(seq 1 30); do \
		if docker compose exec -T db pg_isready -U postgres >/dev/null 2>&1; then \
			echo "Database is ready!"; \
			break; \
		fi; \
		if [ $$i -eq 30 ]; then \
			echo "Error: Database did not become ready in time"; \
			exit 1; \
		fi; \
		echo "Waiting for database... ($$i/30)"; \
		sleep 2; \
	done

# Development command with proper initialization and database preservation
dev: mod-tidy dev-clean
	@echo "Starting services..."
	docker compose up -d db
	@echo "Waiting for database..."
	@for i in $$(seq 1 30); do \
		if docker compose exec -T db pg_isready -U postgres >/dev/null 2>&1; then \
			echo "Database is ready!"; \
			docker compose up --build app; \
			break; \
		fi; \
		if [ $$i -eq 30 ]; then \
			echo "Error: Database did not become ready in time"; \
			exit 1; \
		fi; \
		echo "Waiting for database... ($$i/30)"; \
		sleep 2; \
	done

# Add clean-all command for when you want to remove everything including database
clean-all:
	docker compose down -v
	rm -rf tmp/

# Show all logs
dev-logs:
	docker compose logs -f

# Stop everything
dev-stop:
	docker compose down
	rm -rf tmp/

# Status commands
dev-status:
	@echo "=== Container Status ==="
	@docker compose ps
	@echo "\n=== Database Status ==="
	@docker compose exec db pg_isready -U postgres || echo "Database not ready"

# Migration commands
migrate-create:
	@read -p "Enter migration name: " name; \
	docker compose exec app migrate create -ext sql -dir db/migrations -seq $$name

migrate-up:
	docker compose exec app migrate -path db/migrations -database "postgresql://postgres:postgres@db:5432/postgres?sslmode=disable" up

migrate-down:
	docker compose exec app migrate -path db/migrations -database "postgresql://postgres:postgres@db:5432/postgres?sslmode=disable" down

migrate-force:
	@read -p "Enter version: " version; \
	docker compose exec app migrate -path db/migrations -database "postgresql://postgres:postgres@db:5432/postgres?sslmode=disable" force $$version

migrate-status:
	docker compose exec app migrate -path db/migrations -database "postgresql://postgres:postgres@db:5432/postgres?sslmode=disable" version

migrate-fix:
	@echo "Forcing clean state and reapplying migrations..."
	docker compose exec app migrate -path db/migrations -database "postgresql://postgres:postgres@db:5432/postgres?sslmode=disable" force 0
	docker compose exec app migrate -path db/migrations -database "postgresql://postgres:postgres@db:5432/postgres?sslmode=disable" up

# Add dependency management commands
mod-tidy:
	go mod tidy