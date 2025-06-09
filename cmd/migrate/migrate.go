package main

import (
	"log"
	"os"

	"github.com/username/go-app/config"

	"github.com/pressly/goose/v3"

	_ "github.com/lib/pq"
)

func main() {
	config.LoadEnv()

	dbString := "postgres://" + os.Getenv("DB_USER") + ":" + os.Getenv("DB_PASSWORD") + "@" +
		os.Getenv("DB_HOST") + ":" + os.Getenv("DB_PORT") + "/" + os.Getenv("DB_NAME") + "?sslmode=disable"

	db, err := goose.OpenDBWithDriver("postgres", dbString)
	if err != nil {
		log.Fatalf("goose: failed to open DB: %v\n", err)
	}

	if err := goose.Up(db, "database/migration"); err != nil {
		log.Fatalf("goose up failed: %v\n", err)
	}
}
