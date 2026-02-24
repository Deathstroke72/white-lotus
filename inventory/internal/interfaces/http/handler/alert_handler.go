// file: internal/interfaces/http/handler/alert_handler.go
package handler

import "net/http"

// AlertUseCase defines the use case operations the handler depends on.
type AlertUseCase interface {
	// TODO: define methods once application/usecase/ files are generated.
}

// AlertHandler handles HTTP requests for the /api/v1/alerts resource.
type AlertHandler struct {
	useCase AlertUseCase
}

// NewAlertHandler constructs an AlertHandler with its use case dependency.
func NewAlertHandler(uc AlertUseCase) *AlertHandler {
	return &AlertHandler{useCase: uc}
}

// ListLowStock handles GET /api/v1/alerts/low-stock
func (h *AlertHandler) ListLowStock(w http.ResponseWriter, r *http.Request) {
	// TODO: call h.useCase.GetLowStockAlerts, encode []dto.LowStockAlertResponse
	writeNotImplemented(w)
}
