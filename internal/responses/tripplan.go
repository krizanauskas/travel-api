package responses

import "travelapp/pkg/services/dayplanservice"

type dayplanAttractionResponse struct {
	Name         string `json:"name"`
	CoverImgPath string `json:"cover_img_url"`
}

type dayplanResponse struct {
	Attractions []dayplanAttractionResponse `json:"attractions"`
}

func GetTripPlanDataResponse(plan dayplanservice.DayPlan) dayplanResponse {
	var attractionsResponse []dayplanAttractionResponse

	for _, attraction := range plan.GetAttractions() {
		attractionsResponse = append(attractionsResponse, dayplanAttractionResponse{
			Name:         attraction.Name,
			CoverImgPath: attraction.GetCoverImgPath(),
		})
	}

	return dayplanResponse{
		Attractions: attractionsResponse,
	}
}
