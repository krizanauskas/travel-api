package routes

import (
	"net/http"
	"travelapp/internal/config"
	"travelapp/internal/handlers/tripplanhandler"
	"travelapp/internal/responses"
	"travelapp/pkg/services/attractionservice"

	"github.com/labstack/echo/v4"
)

func InitRoutes(e *echo.Echo, config *config.Config) {
	attractionService := attractionservice.New(config.Server)

	tripPlanHandler := tripplanhandler.New(attractionService)

	e.POST("/trips/plan", tripPlanHandler.Do)
	e.GET("/health", func(c echo.Context) error {
		return c.JSON(http.StatusOK, responses.DataResponse{
			Message: http.StatusText(http.StatusOK),
		})
	})
	e.Static("/uploads", "uploads")
}
