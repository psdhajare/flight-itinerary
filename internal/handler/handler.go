package handler

import (
	"encoding/json"
	"net/http"

	"flight.prof.com/internal/service"
	"github.com/labstack/echo/v4"
)

// ItineraryHandler groups dependencies for HTTP handlers.
type ItineraryHandler struct {
	svc service.ItineraryService
}

// NewItineraryHandler wires in the service.
func NewItineraryHandler(svc service.ItineraryService) *ItineraryHandler {
	return &ItineraryHandler{svc: svc}
}

// CreateItinerary accepts JSON [[src,dst],…] and returns the ordered path.
func (h *ItineraryHandler) CreateItinerary(c echo.Context) error {
	var tickets [][2]string
	if err := c.Bind(&tickets); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request payload"})
	}

	itin, err := h.svc.Reconstruct(tickets)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	b, err := json.Marshal(itin)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "failed to serialize response"})
	}
	b = append(b, '\n')
	return c.JSONBlob(http.StatusOK, b)
}
