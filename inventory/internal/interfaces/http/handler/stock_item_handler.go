// file: internal/interfaces/http/handler/stock_item_handler.go
package handler

import "net/http"

// StockItemUseCase defines the use case operations the handler depends on.
type StockItemUseCase interface {
	// TODO: define methods once application/usecase/ files are generated.
}

// StockItemHandler handles HTTP requests for the /api/v1/stock-items resource.
type StockItemHandler struct {
	useCase StockItemUseCase
}

// NewStockItemHandler constructs a StockItemHandler with its use case dependency.
func NewStockItemHandler(uc StockItemUseCase) *StockItemHandler {
	return &StockItemHandler{useCase: uc}
}

// Create handles POST /api/v1/stock-items
func (h *StockItemHandler) Create(w http.ResponseWriter, r *http.Request) {
	// TODO: decode dto.CreateStockItemRequest, call h.useCase.Create, encode dto.StockItemResponse
	writeNotImplemented(w)
}

// List handles GET /api/v1/stock-items
func (h *StockItemHandler) List(w http.ResponseWriter, r *http.Request) {
	// TODO: parse query params into dto.ListStockItemsRequest, call h.useCase.List, encode dto.ListStockItemsResponse
	writeNotImplemented(w)
}

// Get handles GET /api/v1/stock-items/{stockItemId}
func (h *StockItemHandler) Get(w http.ResponseWriter, r *http.Request) {
	// stockItemID := r.PathValue("stockItemId")
	// TODO: call h.useCase.GetByID, encode dto.StockItemResponse
	writeNotImplemented(w)
}

// GetAggregatedStock handles GET /api/v1/products/{productId}/stock
func (h *StockItemHandler) GetAggregatedStock(w http.ResponseWriter, r *http.Request) {
	// productID := r.PathValue("productId")
	// TODO: call h.useCase.AggregateByProduct, encode dto.AggregatedStockResponse
	writeNotImplemented(w)
}
