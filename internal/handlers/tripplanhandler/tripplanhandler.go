package tripplanhandler

import (
	// "encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"time"
	"travelapp/internal/requests"
	"travelapp/internal/responses"
	"travelapp/pkg/services/attractionservice"
	"travelapp/pkg/services/dayplanservice"
	"travelapp/pkg/services/tripplaner"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type TripPlanHandler struct {
	attractionService attractionservice.AttractionService
}

func New(attractionservice attractionservice.AttractionService) TripPlanHandler {
	return TripPlanHandler{
		attractionService: attractionservice,
	}
}

func (h TripPlanHandler) Do(c echo.Context) error {
	var tripPlanRequest requests.TripPlan

	if err := c.Bind(&tripPlanRequest); err != nil {
		return c.JSON(http.StatusBadRequest, responses.ErrorResponse{})
	}

	validate := validator.New(validator.WithRequiredStructEnabled())
	if err := validate.Struct(tripPlanRequest); err != nil {
		validationErrors := err.(validator.ValidationErrors)
		errorResponse := responses.ErrorResponseWithFields{
			Fields: make(map[string]string),
		}

		typ := reflect.TypeOf(tripPlanRequest)

		for _, err := range validationErrors {
			fieldName := err.Field()
			structField, _ := typ.FieldByName(fieldName)
			jsonTag := structField.Tag.Get("json")

			if jsonTag == "" {
				jsonTag = fieldName // Fallback to field name if json tag is not set
			}

			errorResponse.Fields[jsonTag] = fmt.Sprintf("failed validation: %s", err.Tag())
		}

		return c.JSON(http.StatusBadRequest, errorResponse)
	}

	t := tripplaner.New(h.attractionService)

	from := time.Time{}

	startingLocation := dayplanservice.Location{
		Lat: 54.72412012442995,
		Lng: 25.287710216279535,
	}

	dayPlan, err := t.PlanDay(nil, from, startingLocation)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.ErrorResponse{
			Message: err.Error(),
		})
	}

	// attractionJSON, err := json.MarshalIndent(dayPlan.GetAttractions(), "", "  ")
	// if err != nil {
	// 	return c.JSON(http.StatusInternalServerError, responses.ErrorResponse{
	// 		StatusText: err.Error(),
	// 	})
	// }

	dayplanData := responses.GetTripPlanDataResponse(dayPlan)

	return c.JSON(http.StatusOK, responses.DataResponse{
		Message: "dayplan response",
		Data:    dayplanData,
	})
}
