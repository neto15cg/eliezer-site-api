FROM golang:1.21-alpine

WORKDIR /app

# Add required packages with correct package names
RUN apk add --no-cache \
    netcat-openbsd \
    postgresql-client \
    gcc \
    musl-dev \
    && go install github.com/cosmtrek/air@v1.44.0

# Copy dependency files first
COPY go.mod go.sum ./
RUN go mod download

# Copy all application files
COPY . .

# Development with hot reload
EXPOSE 8080
CMD ["air"]
