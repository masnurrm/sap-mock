package handler

import (
    "database/sql"
    "encoding/json"
    "fmt"
    "net/http"
    "time"
    "github.com/go-chi/chi/v5"
    "sap-mock/model"
)

type OrderHandler struct { db *sql.DB }

func NewOrderHandler(db *sql.DB) *OrderHandler {
    return &OrderHandler{db: db}
}

func (h *OrderHandler) Create(w http.ResponseWriter, r *http.Request) {
    var req model.CreateOrderRequest
    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
        writeJSON(w, http.StatusBadRequest,
            model.ErrorResponse{Error: "invalid_body", Message: err.Error(), Code: 400})
        return
    }

    // Calculate total
    var total float64
    for _, item := range req.Items {
        total += item.UnitPrice * float64(item.Quantity)
    }

    itemsJSON, _ := json.Marshal(req.Items)
    orderID := fmt.Sprintf("SO-%d", time.Now().UnixMilli())
    currency := req.Currency
    if currency == "" { currency = "USD" }

    _, err := h.db.Exec(
        `INSERT INTO orders (id, customer_id, total_amount, currency, status, items)
         VALUES ($1, $2, $3, $4, 'PENDING', $5)`,
        orderID, req.CustomerID, total, currency, itemsJSON,
    )
    if err != nil {
        writeJSON(w, http.StatusInternalServerError,
            model.ErrorResponse{Error: "db_error", Message: err.Error(), Code: 500})
        return
    }

    writeJSON(w, http.StatusCreated, map[string]interface{}{
        "order_id":    orderID,
        "customer_id": req.CustomerID,
        "total":       total,
        "currency":    currency,
        "status":      "PENDING",
        "created_at":  time.Now(),
    })
}

func (h *OrderHandler) GetByID(w http.ResponseWriter, r *http.Request) {
    id := chi.URLParam(r, "id")
    var o model.Order
    err := h.db.QueryRow(
        `SELECT id, customer_id, total_amount, currency, status, items, created_at
         FROM orders WHERE id = $1`, id,
    ).Scan(&o.ID, &o.CustomerID, &o.TotalAmount, &o.Currency,
           &o.Status, &o.Items, &o.CreatedAt)

    if err == sql.ErrNoRows {
        writeJSON(w, http.StatusNotFound,
            model.ErrorResponse{Error: "not_found", Message: "Order not found", Code: 404})
        return
    }
    if err != nil {
        writeJSON(w, http.StatusInternalServerError,
            model.ErrorResponse{Error: "db_error", Message: err.Error(), Code: 500})
        return
    }
    writeJSON(w, http.StatusOK, o)
}