# Use the official Golang image as a base
FROM golang:latest

RUN apt-get update && apt-get install -y netcat-openbsd

ENV PORT=5000

ENV CONFIG_PATH=./config.json

ENV DB_NAME=auth_db
ENV DB_USER=auth_client
ENV DB_PASS=auth_client
ENV DB_ADDR=localhost:5432
ENV DB_SSL=false

ENV CACHE_ADDR=localhost:6379
ENV CACHE_DB=1
ENV GRPC_USER_SERVICE=localhost:50011

ENV AMQP_HOST=localhost
ENV AMQP_PORT=5672
ENV AMQP_USER=guest
ENV AMQP_PASS=guest
ENV FRONTEND=http://localhost:5173
ENV JWT_SECRET=secret

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download


COPY ./db/migrations /app/db/migrations

# Copy the source code from the current directory to the Working Directory inside the container
COPY . .
COPY ./docker_run.sh /app/docker_run.sh
COPY ./wait_for.sh /app/wait_for.sh

RUN chmod +x /app/docker_run.sh
RUN chmod +x /app/wait_for.sh

# Expose port 5000 to the outside world
EXPOSE ${PORT}

# Command to run the executable
CMD ["bash", "/app/docker_run.sh" ]