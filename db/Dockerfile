# Use the official PostgreSQL image as the base image
FROM postgres:latest

# Set environment variables
ENV POSTGRES_DB=auth_db
ENV POSTGRES_PORT=5432
ENV POSTGRES_USER=auth_client
ENV POSTGRES_PASSWORD=auth_client

# COPY ./migrations /app/migrations

# Expose the PostgreSQL port
EXPOSE ${DB_PORT}