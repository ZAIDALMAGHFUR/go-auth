package main

import (
	"log"
	"os"

	_ "github.com/lib/pq"
	"github.com/pressly/goose/v3"
	"github.com/zaidalmaghfur/go-app/config"
)

func main() {
	config.LoadEnv()

	dbString := "postgres://" + os.Getenv("DB_USER") + ":" + os.Getenv("DB_PASSWORD") + "@" +
		os.Getenv("DB_HOST") + ":" + os.Getenv("DB_PORT") + "/" + os.Getenv("DB_NAME") + "?sslmode=disable"

	db, err := goose.OpenDBWithDriver("postgres", dbString)
	if err != nil {
		log.Fatalf("goose: failed to open DB: %v\n", err)
	}

	if err := goose.Up(db, "database/seed"); err != nil {
		log.Fatalf("goose seed up failed: %v\n", err)
	}

	log.Println("Seed success!")
}
