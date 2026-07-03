-include .env
-include .env.local
export

DRIVER = postgres
MIGRATIONS_DIR=db/migrations
DB_URL=user=$(DB_USER) dbname=$(DB_NAME) password=${DB_PASSWORD} port=${DB_PORT} host=${DB_HOST}

migrate-up:
	goose $(DRIVER) "$(DB_URL)" -dir $(MIGRATIONS_DIR) up
migrate-down:
	goose $(DRIVER) "$(DB_URL)" -dir $(MIGRATIONS_DIR) down
migrate-status:
	goose $(DRIVER) "$(DB_URL)" -dir $(MIGRATIONS_DIR) status
migrate-create:
	goose $(DRIVER) "$(DB_URL)" -dir $(MIGRATIONS_DIR) create $(name) sql