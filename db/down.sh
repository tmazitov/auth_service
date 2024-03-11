DRIVER="postgres"
DATABASE_URL="postgresql://auth_client:auth_client@127.0.0.1:5432/auth_db?sslmode=disable"

goose -dir db/migrations $DRIVER $DATABASE_URL down