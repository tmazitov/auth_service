CONFIG="./config.json"

# Database
DB_ADDR="localhost:5432"
DB_NAME="auth_db"
DB_USER="auth_client"
DB_PASS="auth_client"

# Redis
CACHE_ADDR="localhost:6379"
CACHE_DB=0

# Rabbit
AMQP_HOST="localhost"
AMQP_PORT=5672
AMQP_USER="guest"
AMQP_PASS="guest"

PORT=5000

# JWT
JWT_SECRET="supersecretkey"

go build -o auth_service . && ./auth_service \
	-port $PORT \
	-config $CONFIG \
    -db_addr $DB_ADDR \
    -db_name $DB_NAME \
    -db_user $DB_USER \
    -db_pass $DB_PASS \
    -cache_addr $CACHE_ADDR \
    -cache_db $CACHE_DB \
    -amqp_host $AMQP_HOST \
    -amqp_port $AMQP_PORT \
    -amqp_user $AMQP_USER \
    -amqp_pass $AMQP_PASS \
    -jwt_secret $JWT_SECRET