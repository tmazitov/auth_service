./wait_for.sh "$AMQP_HOST" "$AMQP_PORT"

IFS=':' read -r GRPC_USER_SERVICE_ADDR GRPC_USER_SERVICE_PORT <<< "$GRPC_USER_SERVICE"

./wait_for.sh "$GRPC_USER_SERVICE_ADDR" "$GRPC_USER_SERVICE_PORT"

go build -o auth_service . && exec ./auth_service \
	-port $PORT \
	-config $CONFIG_PATH \
    -db_addr $DB_ADDR \
    -db_name $DB_NAME \
    -db_user $DB_USER \
    -db_pass $DB_PASS \
    -cache_addr $CACHE_ADDR \
    -cache_db $CACHE_DB \
    -jwt_secret $JWT_SECRET \
    -amqp_host $AMQP_HOST \
    -amqp_port $AMQP_PORT \
    -amqp_user $AMQP_USER \
    -amqp_pass $AMQP_PASS \
    -grpc_user_service $GRPC_USER_SERVICE \
    -frontend $FRONTEND
	-mode "release"