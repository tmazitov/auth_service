#!/bin/bash
GOOSE_DBSTRING="postgresql://$DB_USER:$DB_PASS@$DB_ADDR/$DB_NAME?sslmode=disable"
GOOSE_DRIVER="postgres"
goose "$GOOSE_DBSTRING" up