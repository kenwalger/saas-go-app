package db

import (
	"os"
	"testing"
)

func TestInitPrimaryDB(t *testing.T) {
	// Skip if DATABASE_URL is not set
	if os.Getenv("DATABASE_URL") == "" {
		t.Skip("DATABASE_URL not set, skipping database test")
	}

	err := InitPrimaryDB()
	if err != nil {
		t.Fatalf("Failed to initialize primary database: %v", err)
	}

	if PrimaryDB == nil {
		t.Fatal("PrimaryDB is nil")
	}

	err = PrimaryDB.Ping()
	if err != nil {
		t.Fatalf("Failed to ping primary database: %v", err)
	}

	CloseDB()
}

func TestCreateTables(t *testing.T) {
	// Skip if DATABASE_URL is not set
	if os.Getenv("DATABASE_URL") == "" {
		t.Skip("DATABASE_URL not set, skipping database test")
	}

	err := InitPrimaryDB()
	if err != nil {
		t.Fatalf("Failed to initialize primary database: %v", err)
	}
	defer CloseDB()

	err = CreateTables()
	if err != nil {
		t.Fatalf("Failed to create tables: %v", err)
	}
}

