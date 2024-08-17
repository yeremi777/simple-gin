package configs

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/pressly/goose/v3"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func LoadMigrations() {
	dsn := os.Getenv("GOOSE_DBSTRING")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

	sqlDB, err := db.DB()

	if err != nil {
		log.Fatal(err)
	}

	defer sqlDB.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	goose.SetBaseFS(os.DirFS("."))
	goose.SetDialect("postgres")

	command := os.Args[1]
	if err := goose.RunContext(ctx, command, sqlDB, os.Getenv("GOOSE_MIGRATION_DIR")); err != nil {
		log.Fatalf("failed to run goose command: %v", err)
	}
}
