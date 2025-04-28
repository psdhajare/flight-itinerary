package handler

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"flight.prof.com/internal/service"
	"github.com/labstack/echo/v4"
)

func TestCreateItineraryHandler_Success(t *testing.T) {
	e := echo.New()
	svc := service.NewItineraryService()
	h := NewItineraryHandler(svc)

	payload := `[["LAX","DXB"],["JFK","LAX"],["SFO","SJC"],["DXB","SFO"]]`
	req := httptest.NewRequest(http.MethodPost, "/itinerary", bytes.NewBufferString(payload))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	if err := h.CreateItinerary(c); err != nil {
		t.Fatalf("handler error: %v", err)
	}
	if rec.Code != http.StatusOK {
		t.Fatalf("expected %d, got %d", http.StatusOK, rec.Code)
	}
	// Expect a trailing newline
	expected := `["JFK","LAX","DXB","SFO","SJC"]
`
	if rec.Body.String() != expected {
		t.Errorf("got %q, want %q", rec.Body.String(), expected)
	}
}

func TestCreateItineraryHandler_InvalidPayload(t *testing.T) {
	e := echo.New()
	svc := service.NewItineraryService()
	h := NewItineraryHandler(svc)

	req := httptest.NewRequest(http.MethodPost, "/itinerary", bytes.NewBufferString(`{"foo":"bar"}`))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	if err := h.CreateItinerary(c); err != nil {
		t.Fatalf("handler error: %v", err)
	}
	if rec.Code != http.StatusBadRequest {
		t.Errorf("expected %d, got %d", http.StatusBadRequest, rec.Code)
	}
}
