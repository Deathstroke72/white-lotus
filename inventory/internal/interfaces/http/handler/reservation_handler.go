// file: internal/interfaces/http/handler/reservation_handler.go
package handler

import "net/http"

// ReservationUseCase defines the use case operations the handler depends on.
type ReservationUseCase interface {
	// TODO: define methods once application/usecase/ files are generated.
}

// ReservationHandler handles HTTP requests for the /api/v1/reservations resource.
type ReservationHandler struct {
	useCase ReservationUseCase
}

// NewReservationHandler constructs a ReservationHandler with its use case dependency.
func NewReservationHandler(uc ReservationUseCase) *ReservationHandler {
	return &ReservationHandler{useCase: uc}
}

// Create handles POST /api/v1/reservations
func (h *ReservationHandler) Create(w http.ResponseWriter, r *http.Request) {
	// TODO: decode dto.CreateReservationRequest, call h.useCase.Reserve, encode dto.ReservationResponse
	writeNotImplemented(w)
}

// Get handles GET /api/v1/reservations/{reservationId}
func (h *ReservationHandler) Get(w http.ResponseWriter, r *http.Request) {
	// reservationID := r.PathValue("reservationId")
	// TODO: call h.useCase.GetByID, encode dto.ReservationResponse
	writeNotImplemented(w)
}

// Release handles POST /api/v1/reservations/{reservationId}/release
func (h *ReservationHandler) Release(w http.ResponseWriter, r *http.Request) {
	// reservationID := r.PathValue("reservationId")
	// TODO: decode dto.ReleaseReservationRequest, call h.useCase.Release
	writeNotImplemented(w)
}

// Fulfill handles POST /api/v1/reservations/{reservationId}/fulfill
func (h *ReservationHandler) Fulfill(w http.ResponseWriter, r *http.Request) {
	// reservationID := r.PathValue("reservationId")
	// TODO: decode dto.FulfillReservationRequest, call h.useCase.Fulfill
	writeNotImplemented(w)
}

// ListByOrder handles GET /api/v1/orders/{orderId}/reservations
func (h *ReservationHandler) ListByOrder(w http.ResponseWriter, r *http.Request) {
	// orderID := r.PathValue("orderId")
	// TODO: call h.useCase.ListByOrder, encode dto.ListReservationsResponse
	writeNotImplemented(w)
}
