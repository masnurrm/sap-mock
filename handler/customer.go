package handler

import (
	"database/sql"
	"net/http"

	"github.com/go-chi/chi/v5"
	"sap-mock/model"
)

type CustomerHandler struct { db *sql.DB }

func NewCustomerHandler(db *sql.DB) *CustomerHandler {
    return &CustomerHandler{db: db}
}

func (h *CustomerHandler) GetByID(w http.ResponseWriter, r *http.Request) {
    id := chi.URLParam(r, "id")
    var c model.Customer
    err := h.db.QueryRow(
        `SELECT id, name, email, phone, credit_limit, credit_used, currency, status, created_at
         FROM customers WHERE id = $1`, id,
    ).Scan(&c.ID, &c.Name, &c.Email, &c.Phone, &c.CreditLimit,
           &c.CreditUsed, &c.Currency, &c.Status, &c.CreatedAt)

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
    writeJSON(w, http.StatusOK, c)
}