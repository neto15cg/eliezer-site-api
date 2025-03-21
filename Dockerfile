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
COPY .env .

RUN apk add --no-cache postgresql-client

EXPOSE 8080

CMD ["./main"]
