package repository

import (
	"errors"
	"game_store/models"
	"sync"
)

// CustomerRepository interface defines methods for customer operations
type CustomerRepository interface {
	Create(customer *models.Customer) error
	GetByID(id string) (*models.Customer, error)
	GetAll() ([]*models.Customer, error)
	Update(customer *models.Customer) error
	Delete(id string) error
}

// InMemoryCustomerRepository implements CustomerRepository using in-memory storage
type InMemoryCustomerRepository struct {
	customers map[string]*models.Customer
	mutex     sync.RWMutex
}

// NewInMemoryCustomerRepository creates a new instance of the in-memory repository
func NewInMemoryCustomerRepository() *InMemoryCustomerRepository {
	return &InMemoryCustomerRepository{
		customers: make(map[string]*models.Customer),
	}
}

// Create adds a new customer to the repository
func (r *InMemoryCustomerRepository) Create(customer *models.Customer) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	// Check if a customer with the same email already exists
	for _, c := range r.customers {
		if c.Email == customer.Email {
			return errors.New("customer with this email already exists")
		}
	}

	r.customers[customer.ID] = customer
	return nil
}

// GetByID finds a customer by ID
func (r *InMemoryCustomerRepository) GetByID(id string) (*models.Customer, error) {
	r.mutex.RLock()
	defer r.mutex.RUnlock()

	customer, exists := r.customers[id]
	if !exists {
		return nil, errors.New("customer not found")
	}

	return customer, nil
}

// GetAll returns all customers
func (r *InMemoryCustomerRepository) GetAll() ([]*models.Customer, error) {
	r.mutex.RLock()
	defer r.mutex.RUnlock()

	customers := make([]*models.Customer, 0, len(r.customers))
	for _, customer := range r.customers {
		customers = append(customers, customer)
	}

	return customers, nil
}

// Update updates an existing customer
func (r *InMemoryCustomerRepository) Update(customer *models.Customer) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	// Check if the customer exists
	_, exists := r.customers[customer.ID]
	if !exists {
		return errors.New("customer not found")
	}

	// Check if the email is not being used by another customer
	for id, c := range r.customers {
		if c.Email == customer.Email && id != customer.ID {
			return errors.New("customer with this email already exists")
		}
	}

	r.customers[customer.ID] = customer
	return nil
}

// Delete removes a customer from the repository
func (r *InMemoryCustomerRepository) Delete(id string) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	_, exists := r.customers[id]
	if !exists {
		return errors.New("customer not found")
	}

	delete(r.customers, id)
	return nil
}
