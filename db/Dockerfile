FROM postgres:15-alpine

# Set default environment variables
ENV POSTGRES_USER=postgres
ENV POSTGRES_PASSWORD=postgres
ENV POSTGRES_DB=postgres

# Copy migrations to docker-entrypoint-initdb.d
COPY migrations/*.sql /docker-entrypoint-initdb.d/

# Expose the PostgreSQL port
EXPOSE 5432
