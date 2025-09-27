package models

import (
	"time"

	"github.com/google/uuid"
)

// Customer represents a customer of the game store
type Customer struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Phone     string    `json:"phone"`
	Address   string    `json:"address"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// NewCustomer creates a new Customer instance
func NewCustomer(name, email, phone, address string) *Customer {
	now := time.Now()
	return &Customer{
		ID:        uuid.New().String(),
		Name:      name,
		Email:     email,
		Phone:     phone,
		Address:   address,
		CreatedAt: now,
		UpdatedAt: now,
	}
}
