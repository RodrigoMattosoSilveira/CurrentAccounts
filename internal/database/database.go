package database

import (
	"log"

	"github.com/RodrigoMattosoSilveira/CurrentAccounts/internal/entities/people"

	"fmt"
	"log/slog"
	"os"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitDatabase(dbPath string) (*gorm.DB, error) {
	dbName := os.Getenv("DB_NAME")
	if dbName == "" {
		panic(fmt.Sprintf("DBInit: Invalid DB_NAME environment %s variable", dbName))
	}

	db, err := gorm.Open(sqlite.Open(dbName), &gorm.Config{})
	if err != nil {
		panic(fmt.Sprintf("failed to connect to SQLite %s database", dbName))
	}
	slog.Info(fmt.Sprintf("DBInit: connected successfully to %s", dbName))

	// Miigrate here instead of waiting!
	if err = RunMigrations(db); err != nil {
		panic(fmt.Sprintf("failed to migrate database %s", dbName))
	}
	return db, nil
}

func RunMigrations(db *gorm.DB) error {
	log.Println("Running database migrations...")

	// Migrate the Peple table
	if err := db.AutoMigrate(&people.Person{}); err != nil {
		slog.Info(fmt.Sprintf("DBMigrations: Failed to Migrated %s table", "people"))
		return err
	}
	return nil
}
