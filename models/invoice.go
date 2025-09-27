package models

import (
	"time"

	"github.com/google/uuid"
)

// InvoiceItem represents an individual item in the invoice
type InvoiceItem struct {
	ProductID string  `json:"product_id"`
	Name      string  `json:"name"`
	Price     float64 `json:"price"`
	Quantity  int     `json:"quantity"`
	Subtotal  float64 `json:"subtotal"`
}

// Invoice represents a sale/invoice of the store
type Invoice struct {
	ID            string        `json:"id"`
	CustomerID    string        `json:"customer_id"`
	CustomerName  string        `json:"customer_name"`
	Items         []InvoiceItem `json:"items"`
	Subtotal      float64       `json:"subtotal"`
	Discount      float64       `json:"discount"`
	Total         float64       `json:"total"`
	Status        string        `json:"status"`         // "pending", "paid", "cancelled"
	PaymentMethod string        `json:"payment_method"` // "cash", "card", "pix", "boleto"
	Notes         string        `json:"notes"`
	CreatedAt     time.Time     `json:"created_at"`
	UpdatedAt     time.Time     `json:"updated_at"`
}

// NewInvoice creates a new Invoice instance
func NewInvoice(customerID, customerName, paymentMethod, notes string) *Invoice {
	now := time.Now()
	return &Invoice{
		ID:            uuid.New().String(),
		CustomerID:    customerID,
		CustomerName:  customerName,
		Items:         []InvoiceItem{},
		Subtotal:      0.0,
		Discount:      0.0,
		Total:         0.0,
		Status:        "pending",
		PaymentMethod: paymentMethod,
		Notes:         notes,
		CreatedAt:     now,
		UpdatedAt:     now,
	}
}

// AddItem adds an item to the invoice
func (i *Invoice) AddItem(productID, name string, price float64, quantity int) {
	subtotal := price * float64(quantity)
	item := InvoiceItem{
		ProductID: productID,
		Name:      name,
		Price:     price,
		Quantity:  quantity,
		Subtotal:  subtotal,
	}

	i.Items = append(i.Items, item)
	i.recalculateTotal()
}

// RemoveItem removes an item from the invoice by product ID
func (i *Invoice) RemoveItem(productID string) bool {
	for idx, item := range i.Items {
		if item.ProductID == productID {
			i.Items = append(i.Items[:idx], i.Items[idx+1:]...)
			i.recalculateTotal()
			return true
		}
	}
	return false
}

// UpdateItemQuantity updates the quantity of an item
func (i *Invoice) UpdateItemQuantity(productID string, quantity int) bool {
	for idx, item := range i.Items {
		if item.ProductID == productID {
			i.Items[idx].Quantity = quantity
			i.Items[idx].Subtotal = i.Items[idx].Price * float64(quantity)
			i.recalculateTotal()
			return true
		}
	}
	return false
}

// SetDiscount sets the discount for the invoice
func (i *Invoice) SetDiscount(discount float64) {
	i.Discount = discount
	i.recalculateTotal()
}

// recalculateTotal recalculates the subtotal and total of the invoice
func (i *Invoice) recalculateTotal() {
	i.Subtotal = 0.0
	for _, item := range i.Items {
		i.Subtotal += item.Subtotal
	}
	i.Total = i.Subtotal - i.Discount
	if i.Total < 0 {
		i.Total = 0
	}
}

// MarkAsPaid marks the invoice as paid
func (i *Invoice) MarkAsPaid() {
	i.Status = "paid"
	i.UpdatedAt = time.Now()
}

// Cancel cancels the invoice
func (i *Invoice) Cancel() {
	i.Status = "cancelled"
	i.UpdatedAt = time.Now()
}
