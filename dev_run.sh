CONFIG="./config.json"
DB_ADDR="localhost:5432"
DB_NAME="auth_db"
DB_USER="auth_client"
DB_PASS="auth_client"
CACHE_ADDR="localhost:6379"
CACHE_DB=0
PORT=5000

go build -o auth_service . && ./auth_service \
	-port $PORT \
	-config $CONFIG \
    -db_addr $DB_ADDR \
    -db_name $DB_NAME \
    -db_user $DB_USER \
    -db_pass $DB_PASS \
    -cache_addr $CACHE_ADDR \
    -cache_db $CACHE_DB