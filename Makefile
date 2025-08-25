include .env
export

BINARY=control-center
MIGRATIONS_DIR=./migrations

.PHONY: all build migrate psql clean create-db help

all: build create-db migrate

build:
	go build -o $(BINARY) ./main.go

create-db:
	@echo "Creating database '$(POSTGRES_DB)' if it doesn't exist..."
	@psql postgres://$(POSTGRES_USER):$(POSTGRES_PASSWORD)@localhost:5432/postgres?sslmode=disable -c "CREATE DATABASE $(POSTGRES_DB);" 2>/dev/null || true
	@echo "Database ready."

migrate: create-db
	@if ! command -v goose > /dev/null 2>&1; then \
		echo "Installing goose..."; \
		go install github.com/pressly/goose/v3/cmd/goose@latest; \
	fi
	@echo "Running migrations..."
	@goose -dir $(MIGRATIONS_DIR) postgres "$(DB_STRING)" up
	@echo "Migrations completed."

psql:
	psql "$(DB_STRING)"

clean:
	rm -f $(BINARY)

help:
	@echo "Available commands:"
	@echo "  make all      - Build, create DB, migrate, and run"
	@echo "  make build    - Build the application"
	@echo "  make migrate  - Run database migrations"
	@echo "  make run      - Run the application"
	@echo "  make psql     - Connect to the database"
	@echo "  make clean    - Remove the binary"