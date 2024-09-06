PORT= 8069
PROTOCOL=postgresql
DB_HOST=localhost
DB_PORT=5432
DB_USER=jesse
DB_NAME=tiver
SSL_MODE=disable
DB_CONNECTION_STRING=$(PROTOCOL)://$(DB_USER)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=$(SSL_MODE)
MIGRATIONS_DIR=migrations


serve:
	@echo "serving $(APP_NAME) on port $(PORT)"
	go run main.go

db-drop:
	@echo "Checking if database $(DB_NAME) exists..."
	@if psql -U $(DB_USER) -d postgres -tc "SELECT 1 FROM pg_database WHERE datname = '$(DB_NAME)'" | grep -q 1; then \
		echo "Connecting to Postgres and terminating connections to $(DB_NAME)..."; \
		psql -U $(DB_USER) -d postgres -c "SELECT pg_terminate_backend(pid) FROM pg_stat_activity WHERE datname = '$(DB_NAME)';"; \
		echo "Dropping database $(DB_NAME)..."; \
		psql -U $(DB_USER) -d postgres -c "DROP DATABASE $(DB_NAME);"; \
		echo "Database $(DB_NAME) dropped successfully."; \
	else \
		echo "Database $(DB_NAME) doesn't exist, no need to drop. Exiting."; \
	fi

db-create:
	@echo "Creating database $(DB_NAME)..."
	psql -U $(DB_USER) -d postgres -c "CREATE DATABASE $(DB_NAME);"
	@echo "Database $(DB_NAME) created successfully."

db-migrate:
	@echo "Migrating database $(DB_NAME) with files from $(MIGRATIONS_DIR)..."$(DB_NAME)?sslmode=$(SSL_MODE) up 2
	@echo "Database $(DB_NAME) migrated successfully."
	migrate -source file://$(MIGRATIONS_DIR) -database $(PROTOCOL)://$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=$(SSL_MODE) up 2

db-connect:
	@echo "Connecting to database $(DB_NAME)..."
	psql -U $(DB_USER) -d $(DB_NAME)

db-reset: db-drop db-create db-migrate dump
	@echo "Database $(DB_NAME) reset successfully."

db-tables:
	@psql -U $(DB_USER) -d $(DB_NAME) -c "\dt" 

dump:
	@echo DB_CONNECTION_STRING=$(DB_CONNECTION_STRING)
	@echo "TABLES\n"
	@make db-tables

serve-jq:
	@go run main.go | jq -c

generate:
	@go generate

fresh-run: db-reset generate serve-jq
