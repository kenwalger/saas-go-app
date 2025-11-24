package models

import "time"

// Account represents an account in the system
type Account struct {
	ID         int       `json:"id" db:"id"`
	CustomerID int       `json:"customer_id" db:"customer_id"`
	Name       string    `json:"name" db:"name"`
	Status     string    `json:"status" db:"status"`
	CreatedAt  time.Time `json:"created_at" db:"created_at"`
	UpdatedAt  time.Time `json:"updated_at" db:"updated_at"`
}

// CreateAccountRequest represents the request payload for creating an account
type CreateAccountRequest struct {
	CustomerID int    `json:"customer_id" binding:"required"`
	Name       string `json:"name" binding:"required"`
	Status     string `json:"status" binding:"required"`
}

// UpdateAccountRequest represents the request payload for updating an account
type UpdateAccountRequest struct {
	Name   string `json:"name" binding:"required"`
	Status string `json:"status" binding:"required"`
}

