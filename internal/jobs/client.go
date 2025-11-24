package jobs

import (
	"time"

	"github.com/hibiken/asynq"
)

// NewClient creates a new Asynq client for enqueueing jobs
func NewClient(redisURL string) (*asynq.Client, error) {
	if redisURL == "" {
		return nil, nil
	}
	return asynq.NewClient(asynq.RedisClientOpt{Addr: redisURL}), nil
}

// EnqueueAggregationTask enqueues an aggregation task
func EnqueueAggregationTask(client *asynq.Client, date time.Time) error {
	if client == nil {
		return nil // Skip if Redis is not configured
	}

	task, err := NewAggregationTask(date)
	if err != nil {
		return err
	}

	_, err = client.Enqueue(task, asynq.Queue("default"))
	return err
}

