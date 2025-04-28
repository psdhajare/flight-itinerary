package main

import (
	"flight.prof.com/internal/handler"
	"flight.prof.com/internal/service"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	// wire up service → handler → route
	svc := service.NewItineraryService()
	h := handler.NewItineraryHandler(svc)

	e.POST("/itinerary", h.CreateItinerary)

	e.Logger.Fatal(e.Start(":8080"))
}
