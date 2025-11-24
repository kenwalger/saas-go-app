package auth

import (
	"testing"
	"time"
)

func TestInitJWT(t *testing.T) {
	err := InitJWT()
	if err != nil {
		t.Fatalf("Failed to initialize JWT: %v", err)
	}

	if len(jwtSecret) == 0 {
		t.Fatal("JWT secret is empty")
	}
}

func TestGenerateAndValidateToken(t *testing.T) {
	_ = InitJWT()

	username := "testuser"
	token, err := GenerateToken(username)
	if err != nil {
		t.Fatalf("Failed to generate token: %v", err)
	}

	if token == "" {
		t.Fatal("Generated token is empty")
	}

	claims, err := ValidateToken(token)
	if err != nil {
		t.Fatalf("Failed to validate token: %v", err)
	}

	if claims.Username != username {
		t.Errorf("Expected username %s, got %s", username, claims.Username)
	}
}

func TestHashPassword(t *testing.T) {
	password := "testpassword123"
	hash, err := HashPassword(password)
	if err != nil {
		t.Fatalf("Failed to hash password: %v", err)
	}

	if hash == "" {
		t.Fatal("Hash is empty")
	}

	if hash == password {
		t.Fatal("Hash should not equal the original password")
	}
}

func TestCheckPasswordHash(t *testing.T) {
	password := "testpassword123"
	hash, err := HashPassword(password)
	if err != nil {
		t.Fatalf("Failed to hash password: %v", err)
	}

	if !CheckPasswordHash(password, hash) {
		t.Fatal("Password hash check failed")
	}

	if CheckPasswordHash("wrongpassword", hash) {
		t.Fatal("Password hash check should fail for wrong password")
	}
}

func TestTokenExpiration(t *testing.T) {
	_ = InitJWT()

	// This test would require mocking time or using a very short expiration
	// For now, we'll just verify the token structure
	username := "testuser"
	token, err := GenerateToken(username)
	if err != nil {
		t.Fatalf("Failed to generate token: %v", err)
	}

	claims, err := ValidateToken(token)
	if err != nil {
		t.Fatalf("Failed to validate token: %v", err)
	}

	if claims.ExpiresAt == nil {
		t.Fatal("Token expiration is not set")
	}

	expirationTime := claims.ExpiresAt.Time
	if expirationTime.Before(time.Now()) {
		t.Fatal("Token expiration is in the past")
	}
}

