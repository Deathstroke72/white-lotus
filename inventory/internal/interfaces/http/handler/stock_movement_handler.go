// file: internal/interfaces/http/handler/stock_movement_handler.go
package handler

import "net/http"

// StockMovementUseCase defines the use case operations the handler depends on.
type StockMovementUseCase interface {
	// TODO: define methods once application/usecase/ files are generated.
}

// StockMovementHandler handles HTTP requests for stock movement resources.
type StockMovementHandler struct {
	useCase StockMovementUseCase
}

// NewStockMovementHandler constructs a StockMovementHandler with its use case dependency.
func NewStockMovementHandler(uc StockMovementUseCase) *StockMovementHandler {
	return &StockMovementHandler{useCase: uc}
}

// Replenish handles POST /api/v1/stock-movements/replenish
func (h *StockMovementHandler) Replenish(w http.ResponseWriter, r *http.Request) {
	// TODO: decode dto.ReplenishStockRequest, call h.useCase.Replenish, encode dto.StockMovementResponse
	writeNotImplemented(w)
}

// List handles GET /api/v1/stock-movements
func (h *StockMovementHandler) List(w http.ResponseWriter, r *http.Request) {
	// TODO: parse query params into dto.ListStockMovementsRequest, call h.useCase.List, encode dto.ListStockMovementsResponse
	writeNotImplemented(w)
}

// ListForStockItem handles GET /api/v1/stock-items/{stockItemId}/movements
func (h *StockMovementHandler) ListForStockItem(w http.ResponseWriter, r *http.Request) {
	// stockItemID := r.PathValue("stockItemId")
	// TODO: call h.useCase.ListByStockItem, encode dto.ListStockMovementsResponse
	writeNotImplemented(w)
}
