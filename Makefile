DB_URL=postgres://postgres:root@localhost:5432/go-auth?sslmode=disable
MIGRATION_DIR=database/migration
SEED_DIR=database/seed

Migration:
	goose -dir $(MIGRATION_DIR) create $(name) sql

Seeder:
	goose -dir $(SEED_DIR) create $(name) sql

# Migrate DB
Migrate:
	go run cmd/migrate/migrate.go

# Rollback
rollback:
	goose -dir database/migration postgres "$(DB_URL)" down

rollback-seed:
	goose -dir database/seed postgres "$(DB_URL)" down

# Seed DB
seed:
	go run cmd/seed/seed.go

serve:
	go run cmd/main.go

dev:
	air