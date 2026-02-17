package handler

import (
    "database/sql"
    "net/http"
    "github.com/go-chi/chi/v5"
    "sap-mock/model"
)

type CreditHandler struct { db *sql.DB }

func NewCreditHandler(db *sql.DB) *CreditHandler {
    return &CreditHandler{db: db}
}

func (h *CreditHandler) Check(w http.ResponseWriter, r *http.Request) {
    id := chi.URLParam(r, "id")
    var c model.Customer
    err := h.db.QueryRow(
        `SELECT id, name, credit_limit, credit_used, currency FROM customers WHERE id = $1`, id,
    ).Scan(&c.ID, &c.Name, &c.CreditLimit, &c.CreditUsed, &c.Currency)

    if err == sql.ErrNoRows {
        writeJSON(w, http.StatusNotFound,
            model.ErrorResponse{Error: "not_found", Message: "Customer not found", Code: 404})
        return
    }
    if err != nil {
        writeJSON(w, http.StatusInternalServerError,
            model.ErrorResponse{Error: "db_error", Message: err.Error(), Code: 500})
        return
    }

    balance := c.CreditLimit - c.CreditUsed
    status := "APPROVED"
    if balance <= 0 { status = "BLOCKED" }

    writeJSON(w, http.StatusOK, model.CreditCheckResult{
        CustomerID:    c.ID,
        CustomerName:  c.Name,
        CreditLimit:   c.CreditLimit,
        CreditUsed:    c.CreditUsed,
        CreditBalance: balance,
        Currency:      c.Currency,
        Status:        status,
    })
}