package database

import (
	"fmt"
	"log"

	"app/config/config"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

// RunMigrations handles database migrations
func RunMigrations(cfg *config.Config) error {
	dbURL := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s",
		cfg.DBUser,
		cfg.DBPassword,
		cfg.DBHost,
		cfg.DBPort,
		cfg.DBName,
		cfg.DBSSLMode,
	)

	m, err := migrate.New(
		"file://db/migrations",
		dbURL,
	)
	if err != nil {
		return fmt.Errorf("error creating migrate instance: %v", err)
	}
	defer m.Close()

	// Check if database is dirty
	version, dirty, err := m.Version()
	if err != nil && err != migrate.ErrNilVersion {
		return fmt.Errorf("error checking migration version: %v", err)
	}

	if dirty {
		log.Printf("Found dirty database state at version %d, attempting to force to clean state...", version)
		if err := m.Force(int(version)); err != nil {
			return fmt.Errorf("error forcing migration version: %v", err)
		}
	}

	// Run migrations
	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		return fmt.Errorf("error running migrations: %v", err)
	}

	log.Println("Migrations completed successfully")
	return nil
}
