package routes

import (
	"travelapp/internal/config"
	"travelapp/internal/handlers/tripplanhandler"
	"travelapp/pkg/services/attractionservice"

	"github.com/labstack/echo/v4"
)

func InitRoutes(e *echo.Echo, config *config.Config) {
	attractionService := attractionservice.New(config.Server)

	tripPlanHandler := tripplanhandler.New(attractionService)

	e.POST("/trips/plan", tripPlanHandler.Do)
	e.Static("/uploads", "uploads")
}
