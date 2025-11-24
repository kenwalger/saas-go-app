package jobs

import (
	"context"
	"database/sql"
	"encoding/json"
	"log"
	"time"

	"saas-go-app/internal/db"

	"github.com/hibiken/asynq"
)

const (
	TypeAggregateData = "aggregate:data"
)

// AggregationPayload represents the payload for aggregation jobs
type AggregationPayload struct {
	Date time.Time `json:"date"`
}

// NewAggregationTask creates a new aggregation task
func NewAggregationTask(date time.Time) (*asynq.Task, error) {
	payload, err := json.Marshal(AggregationPayload{Date: date})
	if err != nil {
		return nil, err
	}
	return asynq.NewTask(TypeAggregateData, payload), nil
}

// HandleAggregationTask processes aggregation tasks
func HandleAggregationTask(ctx context.Context, t *asynq.Task) error {
	var payload AggregationPayload
	if err := json.Unmarshal(t.Payload(), &payload); err != nil {
		return err
	}

	log.Printf("Processing aggregation task for date: %s", payload.Date.Format("2006-01-02"))

	// Perform data aggregation
	// This is a demo, so we'll just log some aggregated statistics
	analyticsDB := db.AnalyticsDB
	if analyticsDB == nil {
		analyticsDB = db.PrimaryDB
	}

	var totalCustomers int
	var totalAccounts int
	var activeAccounts int

	err := analyticsDB.QueryRow("SELECT COUNT(*) FROM customers").Scan(&totalCustomers)
	if err != nil && err != sql.ErrNoRows {
		log.Printf("Error aggregating customers: %v", err)
	}

	err = analyticsDB.QueryRow("SELECT COUNT(*) FROM accounts").Scan(&totalAccounts)
	if err != nil && err != sql.ErrNoRows {
		log.Printf("Error aggregating accounts: %v", err)
	}

	err = analyticsDB.QueryRow("SELECT COUNT(*) FROM accounts WHERE status = 'active'").Scan(&activeAccounts)
	if err != nil && err != sql.ErrNoRows {
		log.Printf("Error aggregating active accounts: %v", err)
	}

	log.Printf("Aggregation results - Customers: %d, Accounts: %d, Active: %d", 
		totalCustomers, totalAccounts, activeAccounts)

	// In a real application, you might store these aggregated results in a separate table
	// For this demo, we're just logging them

	return nil
}

