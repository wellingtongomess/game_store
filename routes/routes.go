package routes

import (
	"game_store/handlers"
	"game_store/repository"
	"net/http"

	"github.com/gorilla/mux"
)

// SetupRoutes configures all API routes
func SetupRoutes() *mux.Router {
	router := mux.NewRouter()

	// Initialize repositories
	customerRepo := repository.NewInMemoryCustomerRepository()
	productRepo := repository.NewInMemoryProductRepository()
	invoiceRepo := repository.NewInMemoryInvoiceRepository()

	// Initialize handlers
	customerHandler := handlers.NewCustomerHandler(customerRepo)
	productHandler := handlers.NewProductHandler(productRepo)
	invoiceHandler := handlers.NewInvoiceHandler(invoiceRepo)

	// Customer routes
	customerRoutes := router.PathPrefix("/api/customers").Subrouter()
	customerRoutes.HandleFunc("", customerHandler.CreateCustomer).Methods("POST")
	customerRoutes.HandleFunc("", customerHandler.GetAllCustomers).Methods("GET")
	customerRoutes.HandleFunc("/{id}", customerHandler.GetCustomer).Methods("GET")
	customerRoutes.HandleFunc("/{id}", customerHandler.UpdateCustomer).Methods("PUT")
	customerRoutes.HandleFunc("/{id}", customerHandler.DeleteCustomer).Methods("DELETE")

	// Product routes
	productRoutes := router.PathPrefix("/api/products").Subrouter()
	productRoutes.HandleFunc("", productHandler.CreateProduct).Methods("POST")
	productRoutes.HandleFunc("", productHandler.GetAllProducts).Methods("GET")
	productRoutes.HandleFunc("/category", productHandler.GetProductsByCategory).Methods("GET")
	productRoutes.HandleFunc("/platform", productHandler.GetProductsByPlatform).Methods("GET")
	productRoutes.HandleFunc("/{id}", productHandler.GetProduct).Methods("GET")
	productRoutes.HandleFunc("/{id}", productHandler.UpdateProduct).Methods("PUT")
	productRoutes.HandleFunc("/{id}/stock", productHandler.UpdateStock).Methods("PATCH")
	productRoutes.HandleFunc("/{id}", productHandler.DeleteProduct).Methods("DELETE")

	// Invoice routes
	invoiceRoutes := router.PathPrefix("/api/invoices").Subrouter()
	invoiceRoutes.HandleFunc("", invoiceHandler.CreateInvoice).Methods("POST")
	invoiceRoutes.HandleFunc("", invoiceHandler.GetAllInvoices).Methods("GET")
	invoiceRoutes.HandleFunc("/period", invoiceHandler.GetInvoicesByPeriod).Methods("GET")
	invoiceRoutes.HandleFunc("/total", invoiceHandler.GetTotalSales).Methods("GET")
	invoiceRoutes.HandleFunc("/{id}", invoiceHandler.GetInvoice).Methods("GET")
	invoiceRoutes.HandleFunc("/{id}", invoiceHandler.UpdateInvoice).Methods("PUT")
	invoiceRoutes.HandleFunc("/{id}", invoiceHandler.DeleteInvoice).Methods("DELETE")
	invoiceRoutes.HandleFunc("/{id}/paid", invoiceHandler.MarkAsPaid).Methods("PATCH")
	invoiceRoutes.HandleFunc("/{id}/cancel", invoiceHandler.CancelInvoice).Methods("PATCH")
	invoiceRoutes.HandleFunc("/{id}/items", invoiceHandler.AddItem).Methods("POST")
	invoiceRoutes.HandleFunc("/{id}/items", invoiceHandler.UpdateItem).Methods("PUT")
	invoiceRoutes.HandleFunc("/{id}/items", invoiceHandler.RemoveItem).Methods("DELETE")

	// Health check route
	router.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status": "ok", "message": "API is running"}`))
	}).Methods("GET")

	// Swagger documentation routes
	router.PathPrefix("/swagger/").Handler(http.StripPrefix("/swagger/", http.FileServer(http.Dir("./swagger/"))))
	router.HandleFunc("/docs", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/swagger/", http.StatusMovedPermanently)
	})

	// CORS middleware
	router.Use(func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
			w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

			if r.Method == "OPTIONS" {
				w.WriteHeader(http.StatusOK)
				return
			}

			next.ServeHTTP(w, r)
		})
	})

	return router
}
