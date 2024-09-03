package responses

type ErrorResponseWithFields struct {
	Message string            `json:"message"`
	Fields  map[string]string `json:"fields"`
}

type ErrorResponse struct {
	Message string `json:"message"`
}
