FROM golang:1.20-alpine AS builder

RUN apk update && \
    apk upgrade && \
    apk add bash && \
    rm -rf /var/cache/apk/*

ADD https://github.com/pressly/goose/releases/download/v3.19.2/goose_linux_x86_64 /bin/goose
RUN chmod +x /bin/goose

WORKDIR /migrations

ENV GOOSE_DRIVER="postgres"
ENV GOOSE_URL="postgresql://auth_client:auth_client@127.0.0.1:5432/auth_db?sslmode=disable"
ENV DB_NAME=auth_db
ENV DB_USER=auth_client
ENV DB_PASS=auth_client
ENV DB_ADDR=localhost:5432
ENV DB_SSL=false

ADD *.sql /migrations/
ADD entrypoint.sh /migrations/

RUN chmod +x /migrations/entrypoint.sh

ENTRYPOINT ["/migrations/entrypoint.sh"]