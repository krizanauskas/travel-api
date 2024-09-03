package requests

type TripPlan struct {
	Activities []string `json:"activities" validate:"required,dive"`
	DateFrom   string   `json:"date_from" validate:"required"`
	DateTo     string   `json:"date_to" validate:"required"`
	Place      string   `json:"place" validate:"required"`
}
