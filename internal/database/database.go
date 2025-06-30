package database

import (
	"database/sql"
	"embed"
	"fmt"
	"io/fs"
	"path/filepath"
	"strings"
)

//go:embed migrations/*.sql
var migrationsFS embed.FS

// Database represents a database connection
type Database struct {
	*sql.DB
}

// NewDatabase creates a new database connection
func NewDatabase(dbURL string) (*Database, error) {
	db, err := sql.Open("sqlite3", dbURL)
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %w", err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	return &Database{DB: db}, nil
}

// Migrate runs database migrations
func (d *Database) Migrate() error {
	// Create the migrations table if it doesn't exist
	_, err := d.Exec(`
		CREATE TABLE IF NOT EXISTS migrations (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			name TEXT NOT NULL UNIQUE,
			applied_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
		)
	`)
	if err != nil {
		return fmt.Errorf("failed to create migrations table: %w", err)
	}

	// Load and apply all migration files
	migrations, err := fs.ReadDir(migrationsFS, "migrations")
	if err != nil {
		return fmt.Errorf("failed to read migrations directory: %w", err)
	}

	for _, migration := range migrations {
		if migration.IsDir() || !strings.HasSuffix(migration.Name(), ".sql") {
			continue
		}

		// Check if migration has already been applied
		var exists bool
		err := d.QueryRow("SELECT EXISTS(SELECT 1 FROM migrations WHERE name = ?)", migration.Name()).Scan(&exists)
		if err != nil {
			return fmt.Errorf("failed to check if migration exists: %w", err)
		}

		if exists {
			continue
		}

		// Read migration file
		migrationContent, err := fs.ReadFile(migrationsFS, filepath.Join("migrations", migration.Name()))
		if err != nil {
			return fmt.Errorf("failed to read migration file %s: %w", migration.Name(), err)
		}

		// Execute migration within a transaction
		tx, err := d.Begin()
		if err != nil {
			return fmt.Errorf("failed to begin transaction: %w", err)
		}

		_, err = tx.Exec(string(migrationContent))
		if err != nil {
			tx.Rollback()
			return fmt.Errorf("failed to execute migration %s: %w", migration.Name(), err)
		}

		// Record the migration
		_, err = tx.Exec("INSERT INTO migrations (name) VALUES (?)", migration.Name())
		if err != nil {
			tx.Rollback()
			return fmt.Errorf("failed to record migration %s: %w", migration.Name(), err)
		}

		if err := tx.Commit(); err != nil {
			return fmt.Errorf("failed to commit transaction: %w", err)
		}

		fmt.Printf("Applied migration: %s\n", migration.Name())
	}

	return nil
}