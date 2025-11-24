package api

import (
	"database/sql"
	"net/http"
	"strconv"

	"saas-go-app/internal/db"
	"saas-go-app/internal/models"

	"github.com/gin-gonic/gin"
)

// GetAccounts retrieves all accounts
func GetAccounts(c *gin.Context) {
	rows, err := db.PrimaryDB.Query(
		"SELECT id, customer_id, name, status, created_at, updated_at FROM accounts ORDER BY created_at DESC",
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch accounts"})
		return
	}
	defer rows.Close()

	var accounts []models.Account
	for rows.Next() {
		var account models.Account
		if err := rows.Scan(&account.ID, &account.CustomerID, &account.Name, &account.Status, &account.CreatedAt, &account.UpdatedAt); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to scan account"})
			return
		}
		accounts = append(accounts, account)
	}

	c.JSON(http.StatusOK, accounts)
}

// GetAccount retrieves a single account by ID
func GetAccount(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid account ID"})
		return
	}

	var account models.Account
	err = db.PrimaryDB.QueryRow(
		"SELECT id, customer_id, name, status, created_at, updated_at FROM accounts WHERE id = $1",
		id,
	).Scan(&account.ID, &account.CustomerID, &account.Name, &account.Status, &account.CreatedAt, &account.UpdatedAt)

	if err == sql.ErrNoRows {
		c.JSON(http.StatusNotFound, gin.H{"error": "Account not found"})
		return
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch account"})
		return
	}

	c.JSON(http.StatusOK, account)
}

// CreateAccount creates a new account
func CreateAccount(c *gin.Context) {
	var req models.CreateAccountRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var account models.Account
	err := db.PrimaryDB.QueryRow(
		"INSERT INTO accounts (customer_id, name, status) VALUES ($1, $2, $3) RETURNING id, customer_id, name, status, created_at, updated_at",
		req.CustomerID, req.Name, req.Status,
	).Scan(&account.ID, &account.CustomerID, &account.Name, &account.Status, &account.CreatedAt, &account.UpdatedAt)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create account"})
		return
	}

	c.JSON(http.StatusCreated, account)
}

// UpdateAccount updates an existing account
func UpdateAccount(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid account ID"})
		return
	}

	var req models.UpdateAccountRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var account models.Account
	err = db.PrimaryDB.QueryRow(
		"UPDATE accounts SET name = $1, status = $2, updated_at = CURRENT_TIMESTAMP WHERE id = $3 RETURNING id, customer_id, name, status, created_at, updated_at",
		req.Name, req.Status, id,
	).Scan(&account.ID, &account.CustomerID, &account.Name, &account.Status, &account.CreatedAt, &account.UpdatedAt)

	if err == sql.ErrNoRows {
		c.JSON(http.StatusNotFound, gin.H{"error": "Account not found"})
		return
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update account"})
		return
	}

	c.JSON(http.StatusOK, account)
}

// DeleteAccount deletes an account
func DeleteAccount(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid account ID"})
		return
	}

	result, err := db.PrimaryDB.Exec("DELETE FROM accounts WHERE id = $1", id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete account"})
		return
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Account not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Account deleted successfully"})
}

