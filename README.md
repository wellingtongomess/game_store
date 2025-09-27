# Game Store API

Game store API developed in Go with in-memory storage.

## Features

- ✅ Complete CRUD for Customers
- ✅ Complete CRUD for Products
- ✅ Complete CRUD for Invoices
- 🔄 In-memory storage (data is lost on restart)
- 🌐 REST API with CORS enabled
- 📝 Data validation
- 🔒 Thread-safe with mutex
- 🔍 Advanced filters for products (category, platform, active)
- 📦 Stock control
- 💰 Invoice system with items, discounts and status
- 📊 Sales reports and totals by period
- 🛒 Shopping cart management (invoice items)
- 📚 **Interactive API Documentation with Swagger UI**

## How to run

1. Install dependencies:
```bash
go mod tidy
```

2. Run the server:
```bash
go run main.go
```

The server will be available at `http://localhost:8080`

## API Documentation

### Swagger UI
The API includes interactive documentation powered by Swagger UI:

- **Swagger UI**: `http://localhost:8080/swagger/`
- **Documentation Redirect**: `http://localhost:8080/docs`

The Swagger UI provides:
- 📖 Complete API documentation
- 🧪 Interactive API testing
- 📋 Request/response examples
- 🔍 Schema definitions
- 🎯 Try-it-out functionality

## API Endpoints

### Health Check
- **GET** `/health` - Check if the API is running

### Customers
- **POST** `/api/customers` - Create new customer
- **GET** `/api/customers` - List all customers
- **GET** `/api/customers/{id}` - Find customer by ID
- **PUT** `/api/customers/{id}` - Update customer
- **DELETE** `/api/customers/{id}` - Delete customer

### Products
- **POST** `/api/products` - Create new product
- **GET** `/api/products` - List all products
- **GET** `/api/products?active=true` - List only active products
- **GET** `/api/products/category?category=X` - Search products by category
- **GET** `/api/products/platform?platform=X` - Search products by platform
- **GET** `/api/products/{id}` - Find product by ID
- **PUT** `/api/products/{id}` - Update product
- **PATCH** `/api/products/{id}/stock` - Update product stock
- **DELETE** `/api/products/{id}` - Delete product

### Invoices
- **POST** `/api/invoices` - Create new invoice
- **GET** `/api/invoices` - List all invoices
- **GET** `/api/invoices?customer_id=X` - Search invoices by customer
- **GET** `/api/invoices?status=X` - Search invoices by status
- **GET** `/api/invoices/period?start=YYYY-MM-DD&end=YYYY-MM-DD` - Search by period
- **GET** `/api/invoices/total` - Total sales
- **GET** `/api/invoices/total?start=YYYY-MM-DD&end=YYYY-MM-DD` - Total by period
- **GET** `/api/invoices/{id}` - Find invoice by ID
- **PUT** `/api/invoices/{id}` - Update invoice
- **DELETE** `/api/invoices/{id}` - Delete invoice
- **PATCH** `/api/invoices/{id}/paid` - Mark as paid
- **PATCH** `/api/invoices/{id}/cancel` - Cancel invoice
- **POST** `/api/invoices/{id}/items` - Add item to invoice
- **PUT** `/api/invoices/{id}/items?product_id=X` - Update invoice item
- **DELETE** `/api/invoices/{id}/items?product_id=X` - Remove invoice item

## Usage Examples

### Create customer
```bash
curl -X POST http://localhost:8080/api/customers \
  -H "Content-Type: application/json" \
  -d '{
    "name": "John Silva",
    "email": "john@email.com",
    "phone": "(11) 99999-9999",
    "address": "Flower Street, 123"
  }'
```

### List customers
```bash
curl http://localhost:8080/api/customers
```

### Find customer by ID
```bash
curl http://localhost:8080/api/customers/{id}
```

### Update customer
```bash
curl -X PUT http://localhost:8080/api/customers/{id} \
  -H "Content-Type: application/json" \
  -d '{
    "name": "John Silva Santos",
    "email": "john.santos@email.com",
    "phone": "(11) 88888-8888",
    "address": "Flower Street, 456"
  }'
```

### Delete customer
```bash
curl -X DELETE http://localhost:8080/api/customers/{id}
```

### Create product
```bash
curl -X POST http://localhost:8080/api/products \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Cyberpunk 2077",
    "description": "Futuristic action RPG",
    "price": 199.90,
    "category": "RPG",
    "platform": "PC",
    "stock": 50
  }'
```

### List products
```bash
curl http://localhost:8080/api/products
```

### List only active products
```bash
curl http://localhost:8080/api/products?active=true
```

### Search products by category
```bash
curl http://localhost:8080/api/products/category?category=RPG
```

### Search products by platform
```bash
curl http://localhost:8080/api/products/platform?platform=PC
```

### Find product by ID
```bash
curl http://localhost:8080/api/products/{id}
```

### Update product
```bash
curl -X PUT http://localhost:8080/api/products/{id} \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Cyberpunk 2077 - Special Edition",
    "description": "Futuristic action RPG with included DLCs",
    "price": 249.90,
    "category": "RPG",
    "platform": "PC",
    "stock": 30,
    "active": true
  }'
```

### Update stock
```bash
curl -X PATCH http://localhost:8080/api/products/{id}/stock \
  -H "Content-Type: application/json" \
  -d '{
    "quantity": 10
  }'
```

### Delete product
```bash
curl -X DELETE http://localhost:8080/api/products/{id}
```

### Create invoice
```bash
curl -X POST http://localhost:8080/api/invoices \
  -H "Content-Type: application/json" \
  -d '{
    "customer_id": "customer-uuid-here",
    "customer_name": "John Silva",
    "payment_method": "card",
    "notes": "Payment in 3x without interest"
  }'
```

### List invoices
```bash
curl http://localhost:8080/api/invoices
```

### Search invoices by customer
```bash
curl http://localhost:8080/api/invoices?customer_id=customer-uuid-here
```

### Search invoices by status
```bash
curl http://localhost:8080/api/invoices?status=paid
```

### Search invoices by period
```bash
curl http://localhost:8080/api/invoices/period?start=2024-01-01&end=2024-12-31
```

### Add item to invoice
```bash
curl -X POST http://localhost:8080/api/invoices/{id}/items \
  -H "Content-Type: application/json" \
  -d '{
    "product_id": "product-uuid-here",
    "name": "Cyberpunk 2077",
    "price": 199.90,
    "quantity": 2
  }'
```

### Update item quantity
```bash
curl -X PUT http://localhost:8080/api/invoices/{id}/items?product_id=product-uuid-here \
  -H "Content-Type: application/json" \
  -d '{
    "quantity": 3
  }'
```

### Remove item from invoice
```bash
curl -X DELETE http://localhost:8080/api/invoices/{id}/items?product_id=product-uuid-here
```

### Update invoice (apply discount)
```bash
curl -X PUT http://localhost:8080/api/invoices/{id} \
  -H "Content-Type: application/json" \
  -d '{
    "customer_name": "John Silva Santos",
    "discount": 50.00,
    "status": "pending",
    "payment_method": "pix",
    "notes": "R$ 50.00 discount applied"
  }'
```

### Mark invoice as paid
```bash
curl -X PATCH http://localhost:8080/api/invoices/{id}/paid
```

### Cancel invoice
```bash
curl -X PATCH http://localhost:8080/api/invoices/{id}/cancel
```

### Get total sales
```bash
curl http://localhost:8080/api/invoices/total
```

### Get total sales by period
```bash
curl http://localhost:8080/api/invoices/total?start=2024-01-01&end=2024-12-31
```

### Delete invoice
```bash
curl -X DELETE http://localhost:8080/api/invoices/{id}
```

## Project Structure

```
game_store/
├── main.go                    # Application entry point
├── go.mod                     # Go dependencies
├── models/
│   ├── customer.go            # Customer data model
│   ├── product.go             # Product data model
│   └── invoice.go             # Invoice data model
├── repository/
│   ├── customer_repository.go # In-memory repository for Customers
│   ├── product_repository.go  # In-memory repository for Products
│   └── invoice_repository.go  # In-memory repository for Invoices
├── handlers/
│   ├── customer_handler.go    # HTTP handlers for Customers
│   ├── product_handler.go     # HTTP handlers for Products
│   └── invoice_handler.go     # HTTP handlers for Invoices
├── routes/
│   └── routes.go              # Route configuration
└── swagger/
    ├── swagger.json           # OpenAPI 3.0 specification
    └── index.html             # Swagger UI interface
```