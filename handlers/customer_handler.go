package handlers

import (
	"encoding/json"
	"game_store/models"
	"game_store/repository"
	"net/http"
	"time"
)

// CustomerHandler manages HTTP requests for customers
type CustomerHandler struct {
	repo repository.CustomerRepository
}

// NewCustomerHandler creates a new instance of the customer handler
func NewCustomerHandler(repo repository.CustomerRepository) *CustomerHandler {
	return &CustomerHandler{repo: repo}
}

// CreateCustomerRequest represents the data structure for creating a customer
type CreateCustomerRequest struct {
	Name    string `json:"name"`
	Email   string `json:"email"`
	Phone   string `json:"phone"`
	Address string `json:"address"`
}

// UpdateCustomerRequest represents the data structure for updating a customer
type UpdateCustomerRequest struct {
	Name    string `json:"name"`
	Email   string `json:"email"`
	Phone   string `json:"phone"`
	Address string `json:"address"`
}

// CreateCustomer creates a new customer
func (h *CustomerHandler) CreateCustomer(w http.ResponseWriter, r *http.Request) {
	var req CreateCustomerRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid data", http.StatusBadRequest)
		return
	}

	// Basic validation
	if req.Name == "" || req.Email == "" {
		http.Error(w, "Name and email are required", http.StatusBadRequest)
		return
	}

	customer := models.NewCustomer(req.Name, req.Email, req.Phone, req.Address)

	if err := h.repo.Create(customer); err != nil {
		http.Error(w, err.Error(), http.StatusConflict)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(customer)
}

// GetCustomer finds a customer by ID
func (h *CustomerHandler) GetCustomer(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "ID is required", http.StatusBadRequest)
		return
	}

	customer, err := h.repo.GetByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(customer)
}

// GetAllCustomers returns all customers
func (h *CustomerHandler) GetAllCustomers(w http.ResponseWriter, r *http.Request) {
	customers, err := h.repo.GetAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(customers)
}

// UpdateCustomer updates an existing customer
func (h *CustomerHandler) UpdateCustomer(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "ID is required", http.StatusBadRequest)
		return
	}

	var req UpdateCustomerRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid data", http.StatusBadRequest)
		return
	}

	// Find the existing customer
	customer, err := h.repo.GetByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	// Update the fields
	customer.Name = req.Name
	customer.Email = req.Email
	customer.Phone = req.Phone
	customer.Address = req.Address
	customer.UpdatedAt = time.Now()

	if err := h.repo.Update(customer); err != nil {
		http.Error(w, err.Error(), http.StatusConflict)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(customer)
}

// DeleteCustomer removes a customer
func (h *CustomerHandler) DeleteCustomer(w http.ResponseWriter, r *http.Request) {
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
