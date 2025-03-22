package store

import (
	"database/sql"
	"fmt"
	"os"
	"path/filepath"

	_ "modernc.org/sqlite"
)

type Store struct {
	db *sql.DB
}

func New() (*Store, error) {
	configDir, err := os.UserConfigDir()
	if err != nil {
		return nil, fmt.Errorf("NewStore: os.UserConfigDir: %w", err)
	}

	appDir := filepath.Join(configDir, "goeland")
	if err := os.MkdirAll(appDir, 0755); err != nil {
		return nil, fmt.Errorf("NewStore: os.MkdirAll: %w", err)
	}

	dbPath := filepath.Join(appDir, "goeland.db")
	db, err := sql.Open("sqlite", dbPath)
	if err != nil {
		return nil, fmt.Errorf("NewStore: sql.Open: %w", err)
	}

	if _, err := db.Exec(`PRAGMA foreign_keys = ON;`); err != nil {
		return nil, fmt.Errorf("NewStore: db.Exec: %w", err)
	}

	if err := initSchema(db); err != nil {
		return nil, fmt.Errorf("NewStore: initSchema: %w", err)
	}

	return &Store{db: db}, nil
}

func (s *Store) Close() error {
	if s.db != nil {
		return s.db.Close()
	}
	return nil
}

func initSchema(db *sql.DB) error {
	_, err := db.Exec(`
		CREATE TABLE IF NOT EXISTS accounts (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			address TEXT NOT NULL UNIQUE,
			password TEXT NOT NULL
		);
	`)
	return err
}
