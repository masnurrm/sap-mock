package model

import (
    "time"
    "encoding/json"
)

type Customer struct {
    ID          string    `json:"id" db:"id"`
    Name        string    `json:"name" db:"name"`
    Email       string    `json:"email" db:"email"`
    Phone       string    `json:"phone" db:"phone"`
    CreditLimit float64   `json:"credit_limit" db:"credit_limit"`
    CreditUsed  float64   `json:"credit_used" db:"credit_used"`
    Currency    string    `json:"currency" db:"currency"`
    Status      string    `json:"status" db:"status"`
    CreatedAt   time.Time `json:"created_at" db:"created_at"`
}

type Order struct {
    ID          string          `json:"id" db:"id"`
    CustomerID  string          `json:"customer_id" db:"customer_id"`
    TotalAmount float64         `json:"total_amount" db:"total_amount"`
    Currency    string          `json:"currency" db:"currency"`
    Status      string          `json:"status" db:"status"`
    Items       json.RawMessage `json:"items" db:"items"`
    CreatedAt   time.Time       `json:"created_at" db:"created_at"`
}

type OrderItem struct {
    MaterialID string  `json:"material_id"`
    Quantity   int     `json:"quantity"`
    UnitPrice  float64 `json:"unit_price"`
}

type Material struct {
    ID          string    `json:"id" db:"id"`
    Name        string    `json:"name" db:"name"`
    Description string    `json:"description" db:"description"`
    UOM         string    `json:"uom" db:"uom"`
    Price       float64   `json:"price" db:"price"`
    Currency    string    `json:"currency" db:"currency"`
    Stock       int       `json:"stock" db:"stock"`
    Plant       string    `json:"plant" db:"plant"`
    CreatedAt   time.Time `json:"created_at" db:"created_at"`
}

type CreditCheckResult struct {
    CustomerID    string  `json:"customer_id"`
    CustomerName  string  `json:"customer_name"`
    CreditLimit   float64 `json:"credit_limit"`
    CreditUsed    float64 `json:"credit_used"`
    CreditBalance float64 `json:"credit_balance"`
    Currency      string  `json:"currency"`
    Status        string  `json:"status"` // APPROVED, BLOCKED
}

type CreateOrderRequest struct {
    CustomerID string      `json:"customer_id"`
    Items      []OrderItem `json:"items"`
    Currency   string      `json:"currency"`
}

type ErrorResponse struct {
    Error   string `json:"error"`
    Message string `json:"message"`
    Code    int    `json:"code"`
}