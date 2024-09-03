package tripplaner

import (
	"fmt"
	"time"
	"travelapp/pkg/geo"
	"travelapp/pkg/services/attractionservice"
	"travelapp/pkg/services/dayplanservice"
)

type tripPlan struct {
	attractionservice attractionservice.AttractionService
}

func New(attractionService attractionservice.AttractionService) tripPlan {
	return tripPlan{
		attractionservice: attractionService,
	}
}

func (t tripPlan) PlanDay(plannedDays []dayplanservice.DayPlan, date time.Time, startingLocation dayplanservice.Location) (dayplanservice.DayPlan, error) {
	plan, err := dayplanservice.New("10:00", "11:20", startingLocation)
	if err != nil {
		return dayplanservice.DayPlan{}, err
	}

	t.planBreakfast(&plan)
	t.planAttractions(&plan)

	return plan, nil
}

func (t tripPlan) planBreakfast(plan *dayplanservice.DayPlan) error {
	var closestBreakfastPlace attractionservice.Place
	var closestBreakfastdistance float64

	breakfastPlaces, err := t.attractionservice.GetBreakfastPlaces()
	if err != nil {
		fmt.Errorf("failed to get breakfast places %w", err)
	}

	for _, breakfastPlace := range breakfastPlaces {
		if closestBreakfastdistance < 1 {
			closestBreakfastdistance = geo.Haversine(breakfastPlace.Lat, breakfastPlace.Lng, plan.HomeLocation.Lat, plan.HomeLocation.Lng)
			closestBreakfastPlace = breakfastPlace

			continue
		}

		distance := geo.Haversine(breakfastPlace.Lat, breakfastPlace.Lng, plan.HomeLocation.Lat, plan.HomeLocation.Lng)

		if distance < closestBreakfastdistance {
			closestBreakfastdistance = distance
			closestBreakfastPlace = breakfastPlace
		}
	}

	plan.AddAttraction(closestBreakfastPlace)

	return nil
}

func (t tripPlan) planAttractions(plan *dayplanservice.DayPlan) error {
	atractionPlaces, err := t.attractionservice.GetAttractionPlaces()
	if err != nil {
		fmt.Errorf("failed to get attraction places %w", err)
	}

	t.planAttractionsRecursive(atractionPlaces, plan)

	return nil
}

func (t tripPlan) planAttractionsRecursive(attractions []attractionservice.Place, plan *dayplanservice.DayPlan) {
	if plan.IsActivityFull() || len(attractions) == 0 {
		return
	}

	var closestAttractionDistance float64
	var closestAttractionIndex int
	var closestAttraction attractionservice.Place

	lastAttraction := plan.GetLastAttractionLocation()

	for index, attraction := range attractions {
		if closestAttractionDistance == 0 {
			closestAttractionDistance = geo.Haversine(attraction.Lat, attraction.Lng, lastAttraction.Lat, lastAttraction.Lng)
			closestAttraction = attraction
			closestAttractionIndex = index

			continue
		}

		distance := geo.Haversine(attraction.Lat, attraction.Lng, lastAttraction.Lat, lastAttraction.Lng)

		if distance < closestAttractionDistance {
			closestAttractionDistance = distance
			closestAttraction = attraction
			closestAttractionIndex = index
		}
	}

	closestAttraction.DistanceToPrevious = closestAttractionDistance

	plan.AddAttraction(closestAttraction)

	// remove attraction from slice
	attractions = append(attractions[:closestAttractionIndex], attractions[closestAttractionIndex+1:]...)

	t.planAttractionsRecursive(attractions, plan)
}
