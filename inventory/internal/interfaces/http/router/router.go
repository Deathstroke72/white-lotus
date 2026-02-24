// file: internal/interfaces/http/router/router.go
package router

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/inventory-service/internal/interfaces/http/handler"
	"github.com/inventory-service/internal/interfaces/http/middleware"
)

// Config holds all handler and middleware dependencies for the router.
type Config struct {
	JWT          *middleware.JWTMiddleware
	RBAC         *middleware.RBACMiddleware
	Product      *handler.ProductHandler
	Warehouse    *handler.WarehouseHandler
	StockItem    *handler.StockItemHandler
	Reservation  *handler.ReservationHandler
	StockMovement *handler.StockMovementHandler
	Alert        *handler.AlertHandler
}

// New builds and returns the fully-wired http.Handler.
// Routes are registered using Go 1.22+ enhanced ServeMux patterns (METHOD path).
// All /api/v1/* routes are protected by JWT authentication and RBAC.
func New(cfg Config) http.Handler {
	mux := http.NewServeMux()

	// Health check — unauthenticated
	mux.HandleFunc("GET /healthz", handleHealth)

	// Authenticated route chain: JWT → RBAC → handler
	auth := func(h http.HandlerFunc) http.Handler {
		return cfg.JWT.Middleware(cfg.RBAC.Middleware(http.HandlerFunc(h)))
	}

	// ── Products ─────────────────────────────────────────────────────────────
	mux.Handle("POST /api/v1/products",                      auth(cfg.Product.Create))
	mux.Handle("GET /api/v1/products",                       auth(cfg.Product.List))
	mux.Handle("GET /api/v1/products/{productId}",           auth(cfg.Product.Get))
	mux.Handle("PUT /api/v1/products/{productId}",           auth(cfg.Product.Update))
	mux.Handle("DELETE /api/v1/products/{productId}",        auth(cfg.Product.Delete))
	mux.Handle("GET /api/v1/products/{productId}/stock",     auth(cfg.StockItem.GetAggregatedStock))

	// ── Warehouses ────────────────────────────────────────────────────────────
	mux.Handle("POST /api/v1/warehouses",                    auth(cfg.Warehouse.Create))
	mux.Handle("GET /api/v1/warehouses",                     auth(cfg.Warehouse.List))
	mux.Handle("GET /api/v1/warehouses/{warehouseId}",       auth(cfg.Warehouse.Get))
	mux.Handle("PUT /api/v1/warehouses/{warehouseId}",       auth(cfg.Warehouse.Update))
	mux.Handle("DELETE /api/v1/warehouses/{warehouseId}",    auth(cfg.Warehouse.Delete))

	// ── Stock Items ───────────────────────────────────────────────────────────
	mux.Handle("POST /api/v1/stock-items",                            auth(cfg.StockItem.Create))
	mux.Handle("GET /api/v1/stock-items",                             auth(cfg.StockItem.List))
	mux.Handle("GET /api/v1/stock-items/{stockItemId}",               auth(cfg.StockItem.Get))
	mux.Handle("GET /api/v1/stock-items/{stockItemId}/movements",     auth(cfg.StockMovement.ListForStockItem))

	// ── Reservations ──────────────────────────────────────────────────────────
	mux.Handle("POST /api/v1/reservations",                                  auth(cfg.Reservation.Create))
	mux.Handle("GET /api/v1/reservations/{reservationId}",                   auth(cfg.Reservation.Get))
	mux.Handle("POST /api/v1/reservations/{reservationId}/release",          auth(cfg.Reservation.Release))
	mux.Handle("POST /api/v1/reservations/{reservationId}/fulfill",          auth(cfg.Reservation.Fulfill))
	mux.Handle("GET /api/v1/orders/{orderId}/reservations",                  auth(cfg.Reservation.ListByOrder))

	// ── Stock Movements ───────────────────────────────────────────────────────
	mux.Handle("POST /api/v1/stock-movements/replenish",     auth(cfg.StockMovement.Replenish))
	mux.Handle("GET /api/v1/stock-movements",                auth(cfg.StockMovement.List))

	// ── Alerts ────────────────────────────────────────────────────────────────
	mux.Handle("GET /api/v1/alerts/low-stock",               auth(cfg.Alert.ListLowStock))

	return mux
}

func handleHealth(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(map[string]string{
		"status": "ok",
		"time":   time.Now().UTC().Format(time.RFC3339),
	})
}
