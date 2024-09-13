ifneq (,$(wildcard ./.env))
    include .env
    export
endif

ifneq (,$(wildcard ./.env.secret))
    include .env.secret
    export
endif

CMD=cmd
BIN=bin

build:
	@mkdir -p $(BIN)
	@go build -o ./$(BIN)/serve ./$(CMD)/serve 
	@chmod +x ./$(BIN)/serve

db-drop:
	@echo "Checking if database $(DB_NAME) exists..."
	@if psql -U $(DB_USER) -d postgres -tc "SELECT 1 FROM pg_database WHERE datname = '$(DB_NAME)'" | grep -q 1; then \
		echo "Connecting to Postgres and terminating connections to $(DB_NAME)..."; \
		psql -U $(DB_USER) -d postgres -c "SELECT pg_terminate_backend(pid) FROM pg_stat_activity WHERE datname = '$(DB_NAME)';" >> /dev/null; \
		echo "Dropping database $(DB_NAME)..."; \
		psql -U $(DB_USER) -d postgres -c "DROP DATABASE $(DB_NAME);" >> /dev/null; \
		echo "Database $(DB_NAME) dropped successfully."; \
	else \
		echo "Database $(DB_NAME) doesn't exist, no need to drop. Exiting."; \
	fi

db-create:
	@echo "Creating database $(DB_NAME)..."
	@psql -U $(DB_USER) -d postgres -c "CREATE DATABASE $(DB_NAME);" >> /dev/null
	@echo "Database $(DB_NAME) created successfully."

db-migrate:
	@echo "Migrating database $(DB_NAME) with files from $(MIGRATIONS_DIR)..."$(DB_NAME)?sslmode=$(SSL_MODE) up
	@echo "Database $(DB_NAME) migrated successfully."
	@migrate -source file://$(MIGRATIONS_DIR) -database $(DB_PROTOCOL)://$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=$(SSL_MODE) up

db-connect:
	@echo "Connecting to database $(DB_NAME)..."
	@psql -U $(DB_USER) -d $(DB_NAME)

db-reset: db-drop db-create db-migrate dump
	@echo "Database $(DB_NAME) reset successfully."

dump:
	@echo "---=== DB Creds ===---"
	@echo MIGRATIONS_DIR=$(MIGRATIONS_DIR)
	@echo DB_PROTOCOL=$(DB_PROTOCOL)
	@echo DB_USER=$(DB_USER)
	@echo DB_HOST=$(DB_HOST)
	@echo DB_PORT=$(DB_PORT)
	@echo DB_NAME=$(DB_NAME)
	@echo SSL_MODE=$(SSL_MODE)
	@echo DB_CONNECTION_STRING=$(DB_CONNECTION_STRING)
	@echo "---=== Server Settings ===---"
	@echo SCHEME=$(SCHEME)
	@echo HOST=$(HOST)
	@echo PORT=$(PORT)
	@echo "---=== Sender Settings ===---"
	@echo SENDER_SCHEME=$(SENDER_SCHEME)
	@echo SENDER_HOST=$(SENDER_HOST)
	@echo SENDER_PORT=$(SENDER_PORT)

generate:
	@go generate

test: 
	@go test ./...

serve:
	@./$(BIN)/serve

serve-pretty:
	@make serve | humanlog -i

send-dev:
	@air -c cmd/send/.air.toml

send-dev-pretty:
	@make send-dev | humanlog -i

fresh-run: db-reset generate build serve-pretty

dev: 
	@air -c .air.toml

	
dev-pretty: 
	@make dev | humanlog -i

dev-fresh: db-reset generate dev-pretty
