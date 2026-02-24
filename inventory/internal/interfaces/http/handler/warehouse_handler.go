// file: internal/interfaces/http/handler/warehouse_handler.go
package handler

import "net/http"

// WarehouseUseCase defines the use case operations the handler depends on.
type WarehouseUseCase interface {
	// TODO: define methods once application/usecase/ files are generated.
}

// WarehouseHandler handles HTTP requests for the /api/v1/warehouses resource.
type WarehouseHandler struct {
	useCase WarehouseUseCase
}

// NewWarehouseHandler constructs a WarehouseHandler with its use case dependency.
func NewWarehouseHandler(uc WarehouseUseCase) *WarehouseHandler {
	return &WarehouseHandler{useCase: uc}
}

// Create handles POST /api/v1/warehouses
func (h *WarehouseHandler) Create(w http.ResponseWriter, r *http.Request) {
	// TODO: decode dto.CreateWarehouseRequest, call h.useCase.Create, encode dto.WarehouseResponse
	writeNotImplemented(w)
}

// List handles GET /api/v1/warehouses
func (h *WarehouseHandler) List(w http.ResponseWriter, r *http.Request) {
	// TODO: parse query params, call h.useCase.List, encode dto.ListWarehousesResponse
	writeNotImplemented(w)
}

// Get handles GET /api/v1/warehouses/{warehouseId}
func (h *WarehouseHandler) Get(w http.ResponseWriter, r *http.Request) {
	// warehouseID := r.PathValue("warehouseId")
	// TODO: call h.useCase.GetByID, encode dto.WarehouseResponse
	writeNotImplemented(w)
}

// Update handles PUT /api/v1/warehouses/{warehouseId}
func (h *WarehouseHandler) Update(w http.ResponseWriter, r *http.Request) {
	// warehouseID := r.PathValue("warehouseId")
	// TODO: decode dto.UpdateWarehouseRequest, call h.useCase.Update, encode dto.WarehouseResponse
	writeNotImplemented(w)
}

// Delete handles DELETE /api/v1/warehouses/{warehouseId}
func (h *WarehouseHandler) Delete(w http.ResponseWriter, r *http.Request) {
	// warehouseID := r.PathValue("warehouseId")
	// TODO: call h.useCase.Delete
	w.WriteHeader(http.StatusNotImplemented)
}
