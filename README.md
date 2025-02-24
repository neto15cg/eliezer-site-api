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
cd eliezer-site-api
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
- `POST /conversation` - Create a new message
  ```json
  {
    "message": "Hello, World!",
    "conversation_id": "eace393d-15f4-495c-8abe-2f50832147c3"
  }
  ```

- `GET /conversation/:id` - List all messages from this conversation

### Health Check
- `GET /health` - Check API status

## Environment Variables

| Variable | Description | Default |
|----------|-------------|---------|
| APP_NAME | Application name | eliezer-site |
| APP_ENV | Application environment | development |
| APP_PORT | HTTP port | 8080 |
| DB_HOST | Database host | db |
| DB_PORT | Database port | 5432 |
| DB_USER | Database user | postgres |
| DB_PASSWORD | Database password | postgres |
| DB_NAME | Database name | postgres |
| DB_SSL_MODE | Database SSL mode | disable |
| MIGRATION_PATH | Database migrations path | db/migrations |
| OPENAI_API_KEY | OpenAI API key | - |
| CHATBOT_PROMPT | Chatbot initial prompt configuration | - |

