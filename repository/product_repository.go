package repository

import (
	"errors"
	"game_store/models"
	"sync"
	"time"
)

// ProductRepository interface defines methods for product operations
type ProductRepository interface {
	Create(product *models.Product) error
	GetByID(id string) (*models.Product, error)
	GetAll() ([]*models.Product, error)
	GetByCategory(category string) ([]*models.Product, error)
	GetByPlatform(platform string) ([]*models.Product, error)
	GetActive() ([]*models.Product, error)
	Update(product *models.Product) error
	Delete(id string) error
	UpdateStock(id string, quantity int) error
}

// InMemoryProductRepository implements ProductRepository using in-memory storage
type InMemoryProductRepository struct {
	products map[string]*models.Product
	mutex    sync.RWMutex
}

// NewInMemoryProductRepository creates a new instance of the in-memory repository
func NewInMemoryProductRepository() *InMemoryProductRepository {
	return &InMemoryProductRepository{
		products: make(map[string]*models.Product),
	}
}

// Create adds a new product to the repository
func (r *InMemoryProductRepository) Create(product *models.Product) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	// Check if a product with the same name already exists
	for _, p := range r.products {
		if p.Name == product.Name {
			return errors.New("product with this name already exists")
		}
	}

	r.products[product.ID] = product
	return nil
}

// GetByID finds a product by ID
func (r *InMemoryProductRepository) GetByID(id string) (*models.Product, error) {
	r.mutex.RLock()
	defer r.mutex.RUnlock()

	product, exists := r.products[id]
	if !exists {
		return nil, errors.New("product not found")
	}

	return product, nil
}

// GetAll returns all products
func (r *InMemoryProductRepository) GetAll() ([]*models.Product, error) {
	r.mutex.RLock()
	defer r.mutex.RUnlock()

	products := make([]*models.Product, 0, len(r.products))
	for _, product := range r.products {
		products = append(products, product)
	}

	return products, nil
}

// GetByCategory returns products by category
func (r *InMemoryProductRepository) GetByCategory(category string) ([]*models.Product, error) {
	r.mutex.RLock()
	defer r.mutex.RUnlock()

	var products []*models.Product
	for _, product := range r.products {
		if product.Category == category {
			products = append(products, product)
		}
	}

	return products, nil
}

// GetByPlatform returns products by platform
func (r *InMemoryProductRepository) GetByPlatform(platform string) ([]*models.Product, error) {
	r.mutex.RLock()
	defer r.mutex.RUnlock()

	var products []*models.Product
	for _, product := range r.products {
		if product.Platform == platform {
			products = append(products, product)
		}
	}

	return products, nil
}

// GetActive returns only active products
func (r *InMemoryProductRepository) GetActive() ([]*models.Product, error) {
	r.mutex.RLock()
	defer r.mutex.RUnlock()

	var products []*models.Product
	for _, product := range r.products {
		if product.Active {
			products = append(products, product)
		}
	}

	return products, nil
}

// Update updates an existing product
func (r *InMemoryProductRepository) Update(product *models.Product) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	// Check if the product exists
	_, exists := r.products[product.ID]
	if !exists {
		return errors.New("product not found")
	}

	// Check if the name is not being used by another product
	for id, p := range r.products {
		if p.Name == product.Name && id != product.ID {
			return errors.New("product with this name already exists")
		}
	}

	r.products[product.ID] = product
	return nil
}

// Delete removes a product from the repository
func (r *InMemoryProductRepository) Delete(id string) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	_, exists := r.products[id]
	if !exists {
		return errors.New("product not found")
	}

	delete(r.products, id)
	return nil
}

// UpdateStock updates the stock quantity of a product
func (r *InMemoryProductRepository) UpdateStock(id string, quantity int) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	product, exists := r.products[id]
	if !exists {
		return errors.New("product not found")
	}

	if product.Stock+quantity < 0 {
		return errors.New("insufficient stock quantity")
	}

	product.Stock += quantity
	product.UpdatedAt = time.Now()
	return nil
}
