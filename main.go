package main

import (
    "log"
    "net/http"
    "os"
    "time"

    "github.com/go-chi/chi/v5"
    "github.com/go-chi/chi/v5/middleware"
    "sap-mock/db"
    "sap-mock/handler"
)

func main() {
    conn, err := db.NewConnection()
    if err != nil {
        log.Fatal("failed to connect db:", err)
    }
    defer conn.Close()

    // Retry loop for DB readiness
    for i := 0; i < 10; i++ {
        if err = conn.Ping(); err == nil { break }
        log.Printf("waiting for db... attempt %d", i+1)
        time.Sleep(2 * time.Second)
    }

    r := chi.NewRouter()
    r.Use(middleware.Logger)
    r.Use(middleware.Recoverer)
    r.Use(middleware.SetHeader("Content-Type", "application/json"))

    // Health check
    r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
        w.WriteHeader(http.StatusOK)
        w.Write([]byte(`{"status":"ok","service":"sap-mock"}`))
    })

    // SAP APIs (consumed by CRM/D365)
    customerH := handler.NewCustomerHandler(conn)
    orderH    := handler.NewOrderHandler(conn)
    materialH := handler.NewMaterialHandler(conn)
    creditH   := handler.NewCreditHandler(conn)

    r.Route("/sap", func(r chi.Router) {
        r.Get("/customers/{id}",      customerH.GetByID)
        r.Post("/orders",             orderH.Create)
        r.Get("/orders/{id}",         orderH.GetByID)
        r.Get("/materials/{id}",      materialH.GetByID)
        r.Get("/credit-check/{id}",   creditH.Check)
    })

    port := os.Getenv("APP_PORT")
    if port == "" { port = "8081" }

    log.Printf("SAP Mock running on port %s", port)
    log.Fatal(http.ListenAndServe(":"+port, r))
}