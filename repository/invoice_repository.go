package repository

import (
	"errors"
	"game_store/models"
	"sync"
	"time"
)

// InvoiceRepository interface defines methods for invoice operations
type InvoiceRepository interface {
	Create(invoice *models.Invoice) error
	GetByID(id string) (*models.Invoice, error)
	GetAll() ([]*models.Invoice, error)
	GetByCustomerID(customerID string) ([]*models.Invoice, error)
	GetByStatus(status string) ([]*models.Invoice, error)
	GetByPeriod(start, end time.Time) ([]*models.Invoice, error)
	Update(invoice *models.Invoice) error
	Delete(id string) error
	GetTotalSales() (float64, error)
	GetTotalSalesByPeriod(start, end time.Time) (float64, error)
}

// InMemoryInvoiceRepository implements InvoiceRepository using in-memory storage
type InMemoryInvoiceRepository struct {
	invoices map[string]*models.Invoice
	mutex    sync.RWMutex
}

// NewInMemoryInvoiceRepository creates a new instance of the in-memory repository
func NewInMemoryInvoiceRepository() *InMemoryInvoiceRepository {
	return &InMemoryInvoiceRepository{
		invoices: make(map[string]*models.Invoice),
	}
}

// Create adds a new invoice to the repository
func (r *InMemoryInvoiceRepository) Create(invoice *models.Invoice) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	r.invoices[invoice.ID] = invoice
	return nil
}

// GetByID finds an invoice by ID
func (r *InMemoryInvoiceRepository) GetByID(id string) (*models.Invoice, error) {
	r.mutex.RLock()
	defer r.mutex.RUnlock()

	invoice, exists := r.invoices[id]
	if !exists {
		return nil, errors.New("invoice not found")
	}

	return invoice, nil
}

// GetAll returns all invoices
func (r *InMemoryInvoiceRepository) GetAll() ([]*models.Invoice, error) {
	r.mutex.RLock()
	defer r.mutex.RUnlock()

	invoices := make([]*models.Invoice, 0, len(r.invoices))
	for _, invoice := range r.invoices {
		invoices = append(invoices, invoice)
	}

	return invoices, nil
}

// GetByCustomerID returns invoices by customer
func (r *InMemoryInvoiceRepository) GetByCustomerID(customerID string) ([]*models.Invoice, error) {
	r.mutex.RLock()
	defer r.mutex.RUnlock()

	var invoices []*models.Invoice
	for _, invoice := range r.invoices {
		if invoice.CustomerID == customerID {
			invoices = append(invoices, invoice)
		}
	}

	return invoices, nil
}

// GetByStatus returns invoices by status
func (r *InMemoryInvoiceRepository) GetByStatus(status string) ([]*models.Invoice, error) {
	r.mutex.RLock()
	defer r.mutex.RUnlock()

	var invoices []*models.Invoice
	for _, invoice := range r.invoices {
		if invoice.Status == status {
			invoices = append(invoices, invoice)
		}
	}

	return invoices, nil
}

// GetByPeriod returns invoices by period
func (r *InMemoryInvoiceRepository) GetByPeriod(start, end time.Time) ([]*models.Invoice, error) {
	r.mutex.RLock()
	defer r.mutex.RUnlock()

	var invoices []*models.Invoice
	for _, invoice := range r.invoices {
		if invoice.CreatedAt.After(start) && invoice.CreatedAt.Before(end) {
			invoices = append(invoices, invoice)
		}
	}

	return invoices, nil
}

// Update updates an existing invoice
func (r *InMemoryInvoiceRepository) Update(invoice *models.Invoice) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	// Check if the invoice exists
	_, exists := r.invoices[invoice.ID]
	if !exists {
		return errors.New("invoice not found")
	}

	invoice.UpdatedAt = time.Now()
	r.invoices[invoice.ID] = invoice
	return nil
}

// Delete removes an invoice from the repository
func (r *InMemoryInvoiceRepository) Delete(id string) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	_, exists := r.invoices[id]
	if !exists {
		return errors.New("invoice not found")
	}

	delete(r.invoices, id)
	return nil
}

// GetTotalSales returns the total sales of all paid invoices
func (r *InMemoryInvoiceRepository) GetTotalSales() (float64, error) {
	r.mutex.RLock()
	defer r.mutex.RUnlock()

	var total float64
	for _, invoice := range r.invoices {
		if invoice.Status == "paid" {
			total += invoice.Total
		}
	}

	return total, nil
}

// GetTotalSalesByPeriod returns the total sales in a specific period
func (r *InMemoryInvoiceRepository) GetTotalSalesByPeriod(start, end time.Time) (float64, error) {
	r.mutex.RLock()
	defer r.mutex.RUnlock()

	var total float64
	for _, invoice := range r.invoices {
		if invoice.Status == "paid" &&
			invoice.CreatedAt.After(start) &&
			invoice.CreatedAt.Before(end) {
			total += invoice.Total
		}
	}

	return total, nil
}
