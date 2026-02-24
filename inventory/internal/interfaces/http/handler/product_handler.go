// file: internal/interfaces/http/handler/product_handler.go
package handler

import (
	"encoding/json"
	"net/http"
)

// ProductUseCase defines the use case operations the handler depends on.
// Implemented by the application layer (application/usecase/).
type ProductUseCase interface {
	// TODO: define methods once application/usecase/ files are generated,
	// e.g. Create, GetByID, List, Update, Delete.
}

// ProductHandler handles HTTP requests for the /api/v1/products resource.
type ProductHandler struct {
	useCase ProductUseCase
}

// NewProductHandler constructs a ProductHandler with its use case dependency.
func NewProductHandler(uc ProductUseCase) *ProductHandler {
	return &ProductHandler{useCase: uc}
}

// Create handles POST /api/v1/products
func (h *ProductHandler) Create(w http.ResponseWriter, r *http.Request) {
	// TODO: decode dto.CreateProductRequest, call h.useCase.Create, encode dto.ProductResponse
	writeNotImplemented(w)
}

// List handles GET /api/v1/products
func (h *ProductHandler) List(w http.ResponseWriter, r *http.Request) {
	// TODO: parse query params into dto.ListProductsRequest, call h.useCase.List, encode dto.ListProductsResponse
	writeNotImplemented(w)
}

// Get handles GET /api/v1/products/{productId}
func (h *ProductHandler) Get(w http.ResponseWriter, r *http.Request) {
	// productID := r.PathValue("productId")
	// TODO: call h.useCase.GetByID(r.Context(), productID), encode dto.ProductResponse
	writeNotImplemented(w)
}

// Update handles PUT /api/v1/products/{productId}
func (h *ProductHandler) Update(w http.ResponseWriter, r *http.Request) {
	// productID := r.PathValue("productId")
	// TODO: decode dto.UpdateProductRequest, call h.useCase.Update, encode dto.ProductResponse
	writeNotImplemented(w)
}

// Delete handles DELETE /api/v1/products/{productId}
func (h *ProductHandler) Delete(w http.ResponseWriter, r *http.Request) {
	// productID := r.PathValue("productId")
	// TODO: call h.useCase.Delete(r.Context(), productID)
	w.WriteHeader(http.StatusNotImplemented)
}

func writeNotImplemented(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNotImplemented)
	_ = json.NewEncoder(w).Encode(map[string]string{"error": "not implemented"})
}
