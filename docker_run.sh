./wait_for.sh rabbit 5672

go build -o auth_service . && ./auth_service \
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
	-mode "release"