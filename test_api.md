# Game Store API - Test Examples

## Quick Start

1. Start the server:
```bash
go run main.go
```

2. Access Swagger UI:
- Open your browser and go to: `http://localhost:8080/swagger/`
- Or use the redirect: `http://localhost:8080/docs`

## Test the API

### 1. Health Check
```bash
curl http://localhost:8080/health
```

### 2. Create a Customer
```bash
curl -X POST http://localhost:8080/api/customers \
  -H "Content-Type: application/json" \
  -d '{
    "name": "John Doe",
    "email": "john@example.com",
    "phone": "(555) 123-4567",
    "address": "123 Main St, City, State"
  }'
```

### 3. Create a Product
```bash
curl -X POST http://localhost:8080/api/products \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Cyberpunk 2077",
    "description": "Futuristic action RPG",
    "price": 59.99,
    "category": "RPG",
    "platform": "PC",
    "stock": 100
  }'
```

### 4. Create an Invoice
```bash
curl -X POST http://localhost:8080/api/invoices \
  -H "Content-Type: application/json" \
  -d '{
    "customer_id": "CUSTOMER_ID_FROM_STEP_2",
    "customer_name": "John Doe",
    "payment_method": "card",
    "notes": "First purchase"
  }'
```

### 5. Add Item to Invoice
```bash
curl -X POST "http://localhost:8080/api/invoices/INVOICE_ID_FROM_STEP_4/items" \
  -H "Content-Type: application/json" \
  -d '{
    "product_id": "PRODUCT_ID_FROM_STEP_3",
    "name": "Cyberpunk 2077",
    "price": 59.99,
    "quantity": 2
  }'
```

### 6. Mark Invoice as Paid
```bash
curl -X PATCH "http://localhost:8080/api/invoices/INVOICE_ID_FROM_STEP_4/paid"
```

## Swagger UI Features

The Swagger UI provides:

1. **Interactive Documentation**: All endpoints are documented with examples
2. **Try It Out**: Test endpoints directly from the browser
3. **Schema Definitions**: Complete data models and validation rules
4. **Response Examples**: See what responses look like
5. **Authentication**: Ready for future auth implementation

## Available Endpoints

- **Health**: `/health`
- **Customers**: `/api/customers`
- **Products**: `/api/products`
- **Invoices**: `/api/invoices`
- **Documentation**: `/swagger/` or `/docs`
