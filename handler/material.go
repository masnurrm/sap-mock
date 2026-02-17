package handler

import (
    "database/sql"
    "net/http"
    "github.com/go-chi/chi/v5"
    "sap-mock/model"
)

type MaterialHandler struct { db *sql.DB }

func NewMaterialHandler(db *sql.DB) *MaterialHandler {
    return &MaterialHandler{db: db}
}

func (h *MaterialHandler) GetByID(w http.ResponseWriter, r *http.Request) {
    id := chi.URLParam(r, "id")
    var m model.Material
    err := h.db.QueryRow(
        `SELECT id, name, description, uom, price, currency, stock, plant, created_at
         FROM materials WHERE id = $1`, id,
    ).Scan(&m.ID, &m.Name, &m.Description, &m.UOM, &m.Price,
           &m.Currency, &m.Stock, &m.Plant, &m.CreatedAt)

    if err == sql.ErrNoRows {
        writeJSON(w, http.StatusNotFound,
            model.ErrorResponse{Error: "not_found", Message: "Material not found", Code: 404})
        return
    }
    if err != nil {
        writeJSON(w, http.StatusInternalServerError,
            model.ErrorResponse{Error: "db_error", Message: err.Error(), Code: 500})
        return
    }
    writeJSON(w, http.StatusOK, m)
}