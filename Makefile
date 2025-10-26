.PHONY: build run docker-build docker-run docker-compose clean dev dev-db dev-db-status dev-clean dev-status

build:
	go build -o main .

run:
	go run main.go

clean:
	rm -f main
	docker compose down -v

# Clean only application files, preserve database
dev-clean:
	docker compose stop app
	docker compose rm -f app
	rm -rf tmp/

# Development command with proper initialization and database preservation
dev: mod-tidy dev-clean
	@echo "Starting development environment..."
	docker compose -f docker-compose.dev.yml up --build

# Production local test
prod-local: mod-tidy
	docker compose -f docker-compose.yml up --build

# Production deployment commands
prod-build:
	@read -p "Enter version: " version; \
	read -p "Enter image name: " image_name; \
	docker build -t $$image_name:$$version . && \
	docker tag eliezer-site-api:$$version $(AWS_ACCOUNT_ID).dkr.ecr.$(AWS_REGION).amazonaws.com/$$image_name:$$version

prod-push:
	@read -p "Enter version: " version; \
	read -p "Enter image name: " image_name; \
	aws ecr get-login-password --region $(AWS_REGION) | docker login --username AWS --password-stdin $(AWS_ACCOUNT_ID).dkr.ecr.$(AWS_REGION).amazonaws.com && \
	docker push $(AWS_ACCOUNT_ID).dkr.ecr.$(AWS_REGION).amazonaws.com/$$image_name:$$version

# Production deployment with auto-versioning
deploy:
	$(eval VERSION=v$(shell LC_ALL=C tr -dc 'A-Za-z0-9' < /dev/urandom | head -c 4))
	$(eval IMAGE_NAME=$(shell grep IMAGE_NAME .env 2>/dev/null | cut -d '=' -f2 | tr -d ' ' || echo "eliezer-api"))
	$(eval AWS_ECR_URL=$(AWS_ACCOUNT_ID).dkr.ecr.$(AWS_REGION).amazonaws.com)
	@if [ -z "$(IMAGE_NAME)" ] || [ "$(IMAGE_NAME)" = "" ]; then \
		echo "Error: IMAGE_NAME is not set in .env file"; \
		exit 1; \
	fi
	@echo "Deploying version: $(VERSION)"
	@echo "Image name from .env: $(IMAGE_NAME)"
	@echo "AWS ECR URL: $(AWS_ECR_URL)"
	docker build -t $(IMAGE_NAME):$(VERSION) . && \
	docker tag $(IMAGE_NAME):$(VERSION) $(AWS_ECR_URL)/$(IMAGE_NAME):$(VERSION) && \
	aws ecr get-login-password --region $(AWS_REGION) | docker login --username AWS --password-stdin $(AWS_ECR_URL) && \
	docker push $(AWS_ECR_URL)/$(IMAGE_NAME):$(VERSION)
	@echo "Successfully deployed $(IMAGE_NAME):$(VERSION)"

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