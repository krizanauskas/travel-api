package dayplanservice

import (
	"fmt"
	"time"
	"travelapp/pkg/services/attractionservice"
)

const timeLayout = "15:04"

type Location struct {
	Lat float64
	Lng float64
}

type Attraction struct {
	Lat                float64
	Lng                float64
	Id                 int
	Name               string
	Duration           time.Duration
	DistanceToPrevious float64
}

type DayPlan struct {
	attractions  []attractionservice.Place
	startsAt     time.Time
	endsAt       time.Time
	HomeLocation Location
}

func New(startsAt string, endsAt string, homeLocation Location) (DayPlan, error) {
	startsAtTime, err := time.Parse(timeLayout, startsAt)
	if err != nil {
		return DayPlan{}, fmt.Errorf("failed to parse dayplan starts at time: %w", err)
	}

	endsAtTime, err := time.Parse(timeLayout, endsAt)
	if err != nil {
		return DayPlan{}, fmt.Errorf("failed to parse dayplan ends at time: %w", err)
	}

	return DayPlan{
		startsAt:     startsAtTime,
		endsAt:       endsAtTime,
		HomeLocation: homeLocation,
	}, nil
}

func (d DayPlan) GetLastAttractionLocation() Location {
	lastAttraction := d.attractions[len(d.attractions)-1]

	return Location{
		lastAttraction.Lat,
		lastAttraction.Lng,
	}
}

func (d DayPlan) IsActivityFull() bool {
	activityTime := d.startsAt

	for _, attraction := range d.attractions {
		activityTime = activityTime.Add(attraction.Duration)
	}

	return activityTime.After(d.endsAt)
}

func (d *DayPlan) AddAttraction(attraction attractionservice.Place) {
	d.attractions = append(d.attractions, attraction)
}

func (d DayPlan) GetAttractions() []attractionservice.Place {
	return d.attractions
}
