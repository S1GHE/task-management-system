include ../.env

migrate-init:
	goose -dir ./migrations/init postgres "host=localhost user=${DB_USERNAME} password=${DB_PASSWORD} dbname=${DB_NAME} port=${DB_PORT} sslmode=disable" up

migrate-user:
	goose -dir ./migrations/user postgres "host=localhost user=${DB_USERNAME} password=${DB_PASSWORD} dbname=${DB_NAME} port=${DB_PORT} sslmode=disable" up

delete-migrate-user:
	goose -dir ./migrations/user postgres "host=localhost user=${DB_USERNAME} password=${DB_PASSWORD} dbname=${DB_NAME} port=${DB_PORT} sslmode=disable" down