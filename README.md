# Message API

A RESTful API service built with Go, Gin, and PostgreSQL.

## Prerequisites

- Docker and Docker Compose (v2.x)
- Go 1.21 or higher
- Make

## Quick Start

1. Clone the repository
```bash
git clone <repository-url>
cd eliezer-site
```

2. Copy environment file
```bash
cp .env.example .env
```

3. Start the application
```bash
make dev
```

## Development Commands

```bash
# Start development environment with hot reload
make dev

# Stop all services
make dev-stop

# View service logs
make dev-logs

# Check service status
make dev-status

# Database migrations
make migrate-create  # Create new migration
make migrate-up     # Apply migrations
make migrate-down   # Rollback migrations
make migrate-fix    # Fix dirty migration state
```

## API Endpoints

### Messages
- `POST /messages` - Create a new message
  ```json
  {
    "message": "Hello, World!"
  }
  ```

- `GET /messages` - List all messages
- `GET /messages/:id` - Get message by ID

### Health Check
- `GET /health` - Check API status

## Project Structure

```
.
├── db/
│   ├── Dockerfile
│   └── migrations/       # Database migrations
├── internal/
│   ├── domain/          # Interfaces and domain logic
│   ├── handler/         # HTTP handlers
│   ├── repository/      # Data access layer
│   ├── routes/          # Route definitions
│   └── service/         # Business logic
├── models/              # Data models
├── pkg/
│   ├── config/          # Configuration
│   └── database/        # Database utilities
└── docker-compose.yml   # Container orchestration
```

## Environment Variables

| Variable | Description | Default |
|----------|-------------|---------|
| APP_NAME | Application name | eliezer-site |
| APP_PORT | HTTP port | 8080 |
| DB_HOST | Database host | db |
| DB_PORT | Database port | 5432 |
| DB_USER | Database user | postgres |
| DB_PASSWORD | Database password | postgres |
| DB_NAME | Database name | postgres |

## Architecture

This project follows:
- Clean Architecture principles
- SOLID design principles
- Domain-Driven Design concepts
- Repository pattern
- Dependency Injection

## Features

- Hot reload development
- PostgreSQL database
- RESTful API
- Automated migrations
- Docker containerization
- Structured logging
- Health checks