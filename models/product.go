package models

import (
	"time"

	"github.com/google/uuid"
)

// Product represents a product of the game store
type Product struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Price       float64   `json:"price"`
	Category    string    `json:"category"`
	Platform    string    `json:"platform"`
	Stock       int       `json:"stock"`
	Active      bool      `json:"active"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// NewProduct creates a new Product instance
func NewProduct(name, description, category, platform string, price float64, stock int) *Product {
	now := time.Now()
	return &Product{
		ID:          uuid.New().String(),
		Name:        name,
		Description: description,
		Price:       price,
		Category:    category,
		Platform:    platform,
		Stock:       stock,
		Active:      true,
		CreatedAt:   now,
		UpdatedAt:   now,
	}
}
