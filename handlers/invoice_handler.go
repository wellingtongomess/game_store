package handlers

import (
	"encoding/json"
	"game_store/models"
	"game_store/repository"
	"net/http"
	"time"
)

// InvoiceHandler manages HTTP requests for invoices
type InvoiceHandler struct {
	repo repository.InvoiceRepository
}

// NewInvoiceHandler creates a new instance of the invoice handler
func NewInvoiceHandler(repo repository.InvoiceRepository) *InvoiceHandler {
	return &InvoiceHandler{repo: repo}
}

// CreateInvoiceRequest represents the data structure for creating an invoice
type CreateInvoiceRequest struct {
	CustomerID    string `json:"customer_id"`
	CustomerName  string `json:"customer_name"`
	PaymentMethod string `json:"payment_method"`
	Notes         string `json:"notes"`
}

// AddItemRequest represents the structure for adding an item to the invoice
type AddItemRequest struct {
	ProductID string  `json:"product_id"`
	Name      string  `json:"name"`
	Price     float64 `json:"price"`
	Quantity  int     `json:"quantity"`
}

// UpdateItemRequest represents the structure for updating an invoice item
type UpdateItemRequest struct {
	Quantity int `json:"quantity"`
}

// UpdateInvoiceRequest represents the data structure for updating an invoice
type UpdateInvoiceRequest struct {
	CustomerName  string  `json:"customer_name"`
	Discount      float64 `json:"discount"`
	Status        string  `json:"status"`
	PaymentMethod string  `json:"payment_method"`
	Notes         string  `json:"notes"`
}

// CreateInvoice creates a new invoice
func (h *InvoiceHandler) CreateInvoice(w http.ResponseWriter, r *http.Request) {
	var req CreateInvoiceRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid data", http.StatusBadRequest)
		return
	}

	// Basic validation
	if req.CustomerID == "" || req.CustomerName == "" || req.PaymentMethod == "" {
		http.Error(w, "Customer ID, customer name and payment method are required", http.StatusBadRequest)
		return
	}

	invoice := models.NewInvoice(req.CustomerID, req.CustomerName, req.PaymentMethod, req.Notes)

	if err := h.repo.Create(invoice); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(invoice)
}

// GetInvoice finds an invoice by ID
func (h *InvoiceHandler) GetInvoice(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "ID is required", http.StatusBadRequest)
		return
	}

	invoice, err := h.repo.GetByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(invoice)
}

// GetAllInvoices returns all invoices
func (h *InvoiceHandler) GetAllInvoices(w http.ResponseWriter, r *http.Request) {
	// Check optional filters
	customerID := r.URL.Query().Get("customer_id")
	status := r.URL.Query().Get("status")

	var invoices []*models.Invoice
	var err error

	if customerID != "" {
		invoices, err = h.repo.GetByCustomerID(customerID)
	} else if status != "" {
		invoices, err = h.repo.GetByStatus(status)
	} else {
		invoices, err = h.repo.GetAll()
	}

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(invoices)
}

// GetInvoicesByPeriod returns invoices by period
func (h *InvoiceHandler) GetInvoicesByPeriod(w http.ResponseWriter, r *http.Request) {
	startStr := r.URL.Query().Get("start")
	endStr := r.URL.Query().Get("end")

	if startStr == "" || endStr == "" {
		http.Error(w, "Parameters 'start' and 'end' are required (format: 2006-01-02)", http.StatusBadRequest)
		return
	}

	start, err := time.Parse("2006-01-02", startStr)
	if err != nil {
		http.Error(w, "Invalid date format for 'start'", http.StatusBadRequest)
		return
	}

	end, err := time.Parse("2006-01-02", endStr)
	if err != nil {
		http.Error(w, "Invalid date format for 'end'", http.StatusBadRequest)
		return
	}

	invoices, err := h.repo.GetByPeriod(start, end)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(invoices)
}

// AddItem adds an item to the invoice
func (h *InvoiceHandler) AddItem(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "Invoice ID is required", http.StatusBadRequest)
		return
	}

	var req AddItemRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid data", http.StatusBadRequest)
		return
	}

	// Basic validation
	if req.ProductID == "" || req.Name == "" || req.Price <= 0 || req.Quantity <= 0 {
		http.Error(w, "Product ID, name, price and quantity are required (price and quantity must be greater than zero)", http.StatusBadRequest)
		return
	}

	invoice, err := h.repo.GetByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	invoice.AddItem(req.ProductID, req.Name, req.Price, req.Quantity)

	if err := h.repo.Update(invoice); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(invoice)
}

// UpdateItem updates the quantity of an item in the invoice
func (h *InvoiceHandler) UpdateItem(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	productID := r.URL.Query().Get("product_id")

	if id == "" || productID == "" {
		http.Error(w, "Invoice ID and product ID are required", http.StatusBadRequest)
		return
	}

	var req UpdateItemRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid data", http.StatusBadRequest)
		return
	}

	if req.Quantity <= 0 {
		http.Error(w, "Quantity must be greater than zero", http.StatusBadRequest)
		return
	}

	invoice, err := h.repo.GetByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	if !invoice.UpdateItemQuantity(productID, req.Quantity) {
		http.Error(w, "Item not found in invoice", http.StatusNotFound)
		return
	}

	if err := h.repo.Update(invoice); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(invoice)
}

// RemoveItem removes an item from the invoice
func (h *InvoiceHandler) RemoveItem(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	productID := r.URL.Query().Get("product_id")

	if id == "" || productID == "" {
		http.Error(w, "Invoice ID and product ID are required", http.StatusBadRequest)
		return
	}

	invoice, err := h.repo.GetByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	if !invoice.RemoveItem(productID) {
		http.Error(w, "Item not found in invoice", http.StatusNotFound)
		return
	}

	if err := h.repo.Update(invoice); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(invoice)
}

// UpdateInvoice updates an existing invoice
func (h *InvoiceHandler) UpdateInvoice(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "ID is required", http.StatusBadRequest)
		return
	}

	var req UpdateInvoiceRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid data", http.StatusBadRequest)
		return
	}

	invoice, err := h.repo.GetByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	// Update the fields
	invoice.CustomerName = req.CustomerName
	invoice.SetDiscount(req.Discount)
	invoice.Status = req.Status
	invoice.PaymentMethod = req.PaymentMethod
	invoice.Notes = req.Notes

	if err := h.repo.Update(invoice); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(invoice)
}

// MarkAsPaid marks an invoice as paid
func (h *InvoiceHandler) MarkAsPaid(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "ID is required", http.StatusBadRequest)
		return
	}

	invoice, err := h.repo.GetByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	invoice.MarkAsPaid()

	if err := h.repo.Update(invoice); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(invoice)
}

// CancelInvoice cancels an invoice
func (h *InvoiceHandler) CancelInvoice(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "ID is required", http.StatusBadRequest)
		return
	}

	invoice, err := h.repo.GetByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	invoice.Cancel()

	if err := h.repo.Update(invoice); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(invoice)
}

// GetTotalSales returns the total sales
func (h *InvoiceHandler) GetTotalSales(w http.ResponseWriter, r *http.Request) {
	startStr := r.URL.Query().Get("start")
	endStr := r.URL.Query().Get("end")

	var total float64
	var err error

	if startStr != "" && endStr != "" {
		start, err := time.Parse("2006-01-02", startStr)
		if err != nil {
			http.Error(w, "Invalid date format for 'start'", http.StatusBadRequest)
			return
		}

		end, err := time.Parse("2006-01-02", endStr)
		if err != nil {
			http.Error(w, "Invalid date format for 'end'", http.StatusBadRequest)
			return
		}

		total, err = h.repo.GetTotalSalesByPeriod(start, end)
	} else {
		total, err = h.repo.GetTotalSales()
	}

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := map[string]interface{}{
		"total": total,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// DeleteInvoice removes an invoice
func (h *InvoiceHandler) DeleteInvoice(w http.ResponseWriter, r *http.Request) {
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
