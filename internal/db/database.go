package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

var (
	// PrimaryDB is the primary database connection
	PrimaryDB *sql.DB
	
	// AnalyticsDB is the follower pool connection for analytics
	AnalyticsDB *sql.DB
)

// InitPrimaryDB initializes the primary database connection
func InitPrimaryDB() error {
	databaseURL := os.Getenv("DATABASE_URL")
	if databaseURL == "" {
		return fmt.Errorf("DATABASE_URL environment variable is not set")
	}

	var err error
	PrimaryDB, err = sql.Open("postgres", databaseURL)
	if err != nil {
		return fmt.Errorf("failed to open primary database: %w", err)
	}

	if err := PrimaryDB.Ping(); err != nil {
		return fmt.Errorf("failed to ping primary database: %w", err)
	}

	log.Println("Primary database connection established")
	return nil
}

// InitAnalyticsDB initializes the analytics database connection (follower pool)
func InitAnalyticsDB() error {
	analyticsURL := os.Getenv("ANALYTICS_DB_URL")
	if analyticsURL == "" {
		log.Println("ANALYTICS_DB_URL not set, analytics endpoints will use primary DB")
		AnalyticsDB = PrimaryDB
		return nil
	}

	var err error
	AnalyticsDB, err = sql.Open("postgres", analyticsURL)
	if err != nil {
		return fmt.Errorf("failed to open analytics database: %w", err)
	}

	if err := AnalyticsDB.Ping(); err != nil {
		return fmt.Errorf("failed to ping analytics database: %w", err)
	}

	log.Println("Analytics database connection established")
	return nil
}

// CloseDB closes all database connections
func CloseDB() {
	if PrimaryDB != nil {
		PrimaryDB.Close()
	}
	if AnalyticsDB != nil && AnalyticsDB != PrimaryDB {
		AnalyticsDB.Close()
	}
}

// CreateTables creates the necessary database tables
func CreateTables() error {
	customersTable := `
	CREATE TABLE IF NOT EXISTS customers (
		id SERIAL PRIMARY KEY,
		name VARCHAR(255) NOT NULL,
		email VARCHAR(255) NOT NULL UNIQUE,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	);`

	accountsTable := `
	CREATE TABLE IF NOT EXISTS accounts (
		id SERIAL PRIMARY KEY,
		customer_id INTEGER NOT NULL REFERENCES customers(id) ON DELETE CASCADE,
		name VARCHAR(255) NOT NULL,
		status VARCHAR(50) NOT NULL,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	);`

	usersTable := `
	CREATE TABLE IF NOT EXISTS users (
		id SERIAL PRIMARY KEY,
		username VARCHAR(255) NOT NULL UNIQUE,
		password_hash VARCHAR(255) NOT NULL,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	);`

	if _, err := PrimaryDB.Exec(customersTable); err != nil {
		return fmt.Errorf("failed to create customers table: %w", err)
	}

	if _, err := PrimaryDB.Exec(accountsTable); err != nil {
		return fmt.Errorf("failed to create accounts table: %w", err)
	}

	if _, err := PrimaryDB.Exec(usersTable); err != nil {
		return fmt.Errorf("failed to create users table: %w", err)
	}

	log.Println("Database tables created successfully")
	return nil
}

