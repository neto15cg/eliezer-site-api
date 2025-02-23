FROM golang:1.21-alpine

WORKDIR /app

# Add required packages and tools
RUN apk add --no-cache \
    netcat-openbsd \
    postgresql-client \
    gcc \
    musl-dev \
    curl

# Install air for hot reload
RUN go install github.com/cosmtrek/air@v1.44.0

# Install golang-migrate
RUN curl -L https://github.com/golang-migrate/migrate/releases/download/v4.17.0/migrate.linux-amd64.tar.gz | tar xvz \
    && mv migrate /usr/local/bin/migrate \
    && chmod +x /usr/local/bin/migrate

# Initialize Go module
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the application
COPY . .

# Ensure dependencies are available
RUN go mod tidy

# Development with hot reload
EXPOSE 8080
CMD ["air"]
