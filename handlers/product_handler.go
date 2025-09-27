package handlers

import (
	"encoding/json"
	"game_store/models"
	"game_store/repository"
	"net/http"
	"time"
)

// ProductHandler manages HTTP requests for products
type ProductHandler struct {
	repo repository.ProductRepository
}

// NewProductHandler creates a new instance of the product handler
func NewProductHandler(repo repository.ProductRepository) *ProductHandler {
	return &ProductHandler{repo: repo}
}

// CreateProductRequest represents the data structure for creating a product
type CreateProductRequest struct {
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Category    string  `json:"category"`
	Platform    string  `json:"platform"`
	Stock       int     `json:"stock"`
}

// UpdateProductRequest represents the data structure for updating a product
type UpdateProductRequest struct {
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Category    string  `json:"category"`
	Platform    string  `json:"platform"`
	Stock       int     `json:"stock"`
	Active      bool    `json:"active"`
}

// UpdateStockRequest represents the structure for updating stock
type UpdateStockRequest struct {
	Quantity int `json:"quantity"`
}

// CreateProduct creates a new product
func (h *ProductHandler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	var req CreateProductRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid data", http.StatusBadRequest)
		return
	}

	// Basic validation
	if req.Name == "" || req.Description == "" || req.Price <= 0 {
		http.Error(w, "Name, description and price are required (price must be greater than zero)", http.StatusBadRequest)
		return
	}

	product := models.NewProduct(req.Name, req.Description, req.Category, req.Platform, req.Price, req.Stock)

	if err := h.repo.Create(product); err != nil {
		http.Error(w, err.Error(), http.StatusConflict)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(product)
}

// GetProduct finds a product by ID
func (h *ProductHandler) GetProduct(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "ID is required", http.StatusBadRequest)
		return
	}

	product, err := h.repo.GetByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(product)
}

// GetAllProducts returns all products
func (h *ProductHandler) GetAllProducts(w http.ResponseWriter, r *http.Request) {
	// Check if should return only active products
	activeOnly := r.URL.Query().Get("active") == "true"

	var products []*models.Product
	var err error

	if activeOnly {
		products, err = h.repo.GetActive()
	} else {
		products, err = h.repo.GetAll()
	}

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(products)
}

// GetProductsByCategory returns products by category
func (h *ProductHandler) GetProductsByCategory(w http.ResponseWriter, r *http.Request) {
	category := r.URL.Query().Get("category")
	if category == "" {
		http.Error(w, "Category is required", http.StatusBadRequest)
		return
	}

	products, err := h.repo.GetByCategory(category)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(products)
}

// GetProductsByPlatform returns products by platform
func (h *ProductHandler) GetProductsByPlatform(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	if platform == "" {
		http.Error(w, "Platform is required", http.StatusBadRequest)
		return
	}

	products, err := h.repo.GetByPlatform(platform)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(products)
}

// UpdateProduct updates an existing product
func (h *ProductHandler) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "ID is required", http.StatusBadRequest)
		return
	}

	var req UpdateProductRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid data", http.StatusBadRequest)
		return
	}

	// Find the existing product
	product, err := h.repo.GetByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	// Update the fields
	product.Name = req.Name
	product.Description = req.Description
	product.Price = req.Price
	product.Category = req.Category
	product.Platform = req.Platform
	product.Stock = req.Stock
	product.Active = req.Active
	product.UpdatedAt = time.Now()

	if err := h.repo.Update(product); err != nil {
		http.Error(w, err.Error(), http.StatusConflict)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(product)
}

// UpdateStock updates the stock of a product
func (h *ProductHandler) UpdateStock(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "ID is required", http.StatusBadRequest)
		return
	}

	var req UpdateStockRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid data", http.StatusBadRequest)
		return
	}

	if err := h.repo.UpdateStock(id, req.Quantity); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Return the updated product
	product, err := h.repo.GetByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(product)
}

// DeleteProduct removes a product
func (h *ProductHandler) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "ID is required", http.StatusBadRequest)
		return
	}

	if err := h.repo.Delete(id); err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
