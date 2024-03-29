go build -o auth_service . && ./auth_service \
	-config $CONFIG_PATH \
    -db_addr $DB_ADDR \
    -db_name $DB_NAME \
    -db_user $DB_USER \
    -db_pass $DB_PASS \
    -cache_addr $CACHE_ADDR \
    -cache_db $CACHE_DB