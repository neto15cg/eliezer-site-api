FROM golang:1.21-alpine AS builder

WORKDIR /app

RUN apk add --no-cache gcc musl-dev

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o main .

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/main .
COPY --from=builder /app/db/migrations ./db/migrations

# Define build arguments
ARG APP_NAME
ARG APP_ENV
ARG APP_PORT
ARG DB_HOST
ARG DB_PORT
ARG DB_USER
ARG DB_PASSWORD
ARG DB_NAME
ARG DB_SSL_MODE
ARG API_USER
ARG API_PASSWORD
ARG LOG_LEVEL
ARG OPENAI_API_KEY
ARG CHATBOT_PROMPT

# Set environment variables
ENV APP_NAME=$APP_NAME \
    APP_ENV=$APP_ENV \
    APP_PORT=$APP_PORT \
    DB_HOST=$DB_HOST \
    DB_PORT=$DB_PORT \
    DB_USER=$DB_USER \
    DB_PASSWORD=$DB_PASSWORD \
    DB_NAME=$DB_NAME \
    DB_SSL_MODE=$DB_SSL_MODE \
    API_USER=$API_USER \
    API_PASSWORD=$API_PASSWORD \
    LOG_LEVEL=$LOG_LEVEL \
    MIGRATION_PATH=$MIGRATION_PATH \
    OPENAI_API_KEY=$OPENAI_API_KEY \
    CHATBOT_PROMPT=$CHATBOT_PROMPT

RUN apk add --no-cache postgresql-client

EXPOSE 8080

CMD ["./main"]
