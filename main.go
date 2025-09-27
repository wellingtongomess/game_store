package main

import (
	"fmt"
	"game_store/routes"
	"log"
	"net/http"
)

func main() {
	// Configure routes
	router := routes.SetupRoutes()

	// Configure server
	port := ":8080"
	fmt.Printf("🚀 Server started on port %s\n", port)
	fmt.Println("📋 Available endpoints:")
	fmt.Println("  GET    /health - Health check")
	fmt.Println("  GET    /docs - API Documentation (Swagger UI)")
	fmt.Println("  GET    /swagger/ - Swagger UI Interface")
	fmt.Println()
	fmt.Println("👥 CUSTOMERS:")
	fmt.Println("  POST   /api/customers - Create customer")
	fmt.Println("  GET    /api/customers - List all customers")
	fmt.Println("  GET    /api/customers/{id} - Find customer by ID")
	fmt.Println("  PUT    /api/customers/{id} - Update customer")
	fmt.Println("  DELETE /api/customers/{id} - Delete customer")
	fmt.Println()
	fmt.Println("🎮 PRODUCTS:")
	fmt.Println("  POST   /api/products - Create product")
	fmt.Println("  GET    /api/products - List all products")
	fmt.Println("  GET    /api/products?active=true - List only active products")
	fmt.Println("  GET    /api/products/category?category=X - Search by category")
	fmt.Println("  GET    /api/products/platform?platform=X - Search by platform")
	fmt.Println("  GET    /api/products/{id} - Find product by ID")
	fmt.Println("  PUT    /api/products/{id} - Update product")
	fmt.Println("  PATCH  /api/products/{id}/stock - Update stock")
	fmt.Println("  DELETE /api/products/{id} - Delete product")
	fmt.Println()
	fmt.Println("💰 INVOICES:")
	fmt.Println("  POST   /api/invoices - Create invoice")
	fmt.Println("  GET    /api/invoices - List all invoices")
	fmt.Println("  GET    /api/invoices?customer_id=X - Search by customer")
	fmt.Println("  GET    /api/invoices?status=X - Search by status")
	fmt.Println("  GET    /api/invoices/period?start=YYYY-MM-DD&end=YYYY-MM-DD - Search by period")
	fmt.Println("  GET    /api/invoices/total - Total sales")
	fmt.Println("  GET    /api/invoices/total?start=YYYY-MM-DD&end=YYYY-MM-DD - Total by period")
	fmt.Println("  GET    /api/invoices/{id} - Find invoice by ID")
	fmt.Println("  PUT    /api/invoices/{id} - Update invoice")
	fmt.Println("  DELETE /api/invoices/{id} - Delete invoice")
	fmt.Println("  PATCH  /api/invoices/{id}/paid - Mark as paid")
	fmt.Println("  PATCH  /api/invoices/{id}/cancel - Cancel invoice")
	fmt.Println("  POST   /api/invoices/{id}/items - Add item")
	fmt.Println("  PUT    /api/invoices/{id}/items?product_id=X - Update item")
	fmt.Println("  DELETE /api/invoices/{id}/items?product_id=X - Remove item")
	fmt.Println()

	// Start server
	log.Fatal(http.ListenAndServe(port, router))
}
