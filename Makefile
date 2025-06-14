# --------- CONFIGURATION ---------
DB_URL=postgres://postgres:root@localhost:5432/go-auth?sslmode=disable
MIGRATION_DIR=database/migration
SEED_DIR=database/seed
BINARY_NAME=bin/go-auth

# --------- MIGRATION COMMANDS ---------
Migration:
	goose -dir $(MIGRATION_DIR) create $(name) sql

Seeder:
	goose -dir $(SEED_DIR) create $(name) sql

migrate:
	go run cmd/migrate/migrate.go

reset:
	goose -dir $(SEED_DIR) postgres "$(DB_URL)" reset

rollback:
	goose -dir $(MIGRATION_DIR) postgres "$(DB_URL)" down

rollback-seed:
	goose -dir $(SEED_DIR) postgres "$(DB_URL)" down

migrate--seed:
	make migrate
	make seed

# --------- SEEDING ---------
seed:
	go run cmd/seed/seed.go

# --------- APP RUN / DEV ---------
serve:
	go run cmd/main.go

dev:
	air

# --------- BUILD COMMANDS ---------
build:
	go build -o $(BINARY_NAME) ./cmd/main.go

build-linux:
	GOOS=linux GOARCH=amd64 go build -o bin/go-auth-linux ./cmd/main.go

build-windows:
	GOOS=windows GOARCH=amd64 go build -o bin/go-auth.exe ./cmd/main.go

build-mac:
	GOOS=darwin GOARCH=arm64 go build -o bin/go-auth-mac ./cmd/main.go

# --------- CLEAN ---------
clean:
	rm -f $(BINARY_NAME) bin/go-auth-linux bin/go-auth.exe bin/go-auth-mac

# --------- PHONY ---------
.PHONY: Migration Seeder Migrate rollback rollback-seed seed serve dev build build-linux build-windows build-mac clean
