package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"travelapp/internal/config"
	"travelapp/internal/routes"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	env := flag.String("env", "local", "Application environment")

	if err := config.Initialize(*env); err != nil {
		log.Fatalf(fmt.Sprintf("failed to initialize config: %s", err.Error()))
	}

	e := echo.New()

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodDelete, http.MethodOptions},
	}))

	routes.InitRoutes(e, config.Conf)

	e.Start(":8080")
}

// 	fmt.Println("init attractions")

// 	homeLocation := dayplan.Location{
// 		Lat: 54.72412012442995,
// 		Lng: 25.287710216279535,
// 	}

// 	attractions := []dayplan.Attraction{
// 		{
// 			Id:       1,
// 			Lat:      54.68537697300427,
// 			Lng:      25.287255310614462,
// 			Duration: time.Duration(time.Minute * 40),
// 			Name:     "Vilniaus arkikatedra",
// 		},
// 		{
// 			Id:       2,
// 			Lat:      54.687372461209485,
// 			Lng:      25.301149205308647,
// 			Duration: time.Duration(time.Minute * 20),
// 			Name:     "Stalo kalnas",
// 		},
// 		{
// 			Id:       3,
// 			Lat:      54.686978419941326,
// 			Lng:      25.29066424483757,
// 			Duration: time.Duration(time.Minute * 20),
// 			Name:     "Gedimino kalnas",
// 		},
// 		{
// 			Id:       4,
// 			Lat:      54.68387133039172,
// 			Lng:      54.68387133039172,
// 			Duration: time.Duration(time.Minute * 20),
// 			Name:     "Bernardinų sodo šokantis fontanas",
// 		},
// 		{
// 			Id:       5,
// 			Lat:      54.688050864252105,
// 			Lng:      25.280855180039417,
// 			Duration: time.Duration(time.Minute * 20),
// 			Name:     "Vinco Kudirkos aikštė",
// 		},
// 		{
// 			Id:       6,
// 			Lat:      54.70088139142579,
// 			Lng:      25.283196825139097,
// 			Duration: time.Duration(time.Minute * 20),
// 			Name:     "Kalvarijų turgus",
// 		},
// 		{
// 			Id:       7,
// 			Lat:      54.71093792384421,
// 			Lng:      25.261991788272386,
// 			Duration: time.Duration(time.Minute * 20),
// 			Name:     "Akropolis",
// 		},
// 		{
// 			Id:       8,
// 			Lat:      54.6841645595128,
// 			Lng:      25.238656347351863,
// 			Duration: time.Duration(time.Minute * 20),
// 			Name:     "Vingio parkas",
// 		},
// 		{
// 			Id:       9,
// 			Lat:      54.748963427907555,
// 			Lng:      25.292848871531117,
// 			Duration: time.Duration(time.Minute * 20),
// 			Name:     "Verkių Rūmai",
// 		},
// 		{
// 			Id:       10,
// 			Lat:      54.6793267864962,
// 			Lng:      25.285199001600606,
// 			Duration: time.Duration(time.Minute * 20),
// 			Name:     "Iliuzijų muziejus",
// 		},
// 		{
// 			Id:       11,
// 			Lat:      54.67085554223728,
// 			Lng:      25.2843736073611,
// 			Duration: time.Duration(time.Minute * 20),
// 			Name:     "Traukinių muziejus",
// 		},
// 		{
// 			Id:       12,
// 			Lat:      54.68817907309971,
// 			Lng:      25.214089597158058,
// 			Duration: time.Duration(time.Minute * 20),
// 			Name:     "Vilniaus televizijos bokštas",
// 		},
// 		{
// 			Id:       13,
// 			Lat:      54.69214910659866,
// 			Lng:      25.26724660324892,
// 			Duration: time.Duration(time.Minute * 20),
// 			Name:     "Lukiškių kalėjimas 2.0",
// 		},
// 		{
// 			Id:       14,
// 			Lat:      54.692761881869835,
// 			Lng:      25.35247939674971,
// 			Duration: time.Duration(time.Minute * 20),
// 			Name:     "Pūčkorių atodanga",
// 		},
// 	}

// 	breakfastLocations := []dayplan.Attraction{
// 		{
// 			Id:       15,
// 			Lat:      54.68404083458557,
// 			Lng:      25.28922250099148,
// 			Duration: time.Duration(time.Minute * 20),
// 			Name:     "Karštos galvos",
// 		},
// 		{
// 			Id:       16,
// 			Lat:      54.67929956408939,
// 			Lng:      25.290305919414237,
// 			Duration: time.Duration(time.Minute * 20),
// 			Name:     "Pirmas Blynas",
// 		},
// 		{
// 			Id:       17,
// 			Lat:      54.68221482369828,
// 			Lng:      25.290617055640976,
// 			Duration: time.Duration(time.Minute * 20),
// 			Name:     "Beigelistai",
// 		},
// 		{
// 			Id:       18,
// 			Lat:      54.69073023367748,
// 			Lng:      25.26324064204871,
// 			Duration: time.Duration(time.Minute * 20),
// 			Name:     "Daily Poison",
// 		},
// 		{
// 			Id:       19,
// 			Lat:      54.68732380331815,
// 			Lng:      25.293748522165536,
// 			Duration: time.Duration(time.Minute * 20),
// 			Name:     "StrangeLove",
// 		},
// 		{
// 			Id:       20,
// 			Lat:      54.67235397656181,
// 			Lng:      25.27885439440505,
// 			Duration: time.Duration(time.Minute * 20),
// 			Name:     "Špricas Brunch",
// 		},
// 		{
// 			Id:       21,
// 			Lat:      54.71224449729303,
// 			Lng:      25.286007222456234,
// 			Duration: time.Duration(time.Minute * 20),
// 			Name:     "PAK FOODS",
// 		},
// 	}

// 	dayPlan, err := dayplan.New("10:00", "11:20", homeLocation)
// 	if err != nil {
// 		log.Fatal("error: %s", err.Error())
// 	}

// 	planBreakfast(breakfastLocations, &dayPlan)
// 	planAttractions(attractions, &dayPlan)

// 	attractionJSON, err := json.MarshalIndent(dayPlan.GetAttractions(), "", "  ")
// 	if err != nil {
// 		fmt.Println("Error marshaling to JSON:", err)
// 		return
// 	}

// 	fmt.Printf("dayplan attractions: %s", string(attractionJSON))
// }

// func planBreakfast(breakfastLocations []dayplan.Attraction, plan *dayplan.DayPlan) {
// 	var shortestBreakfastLocation dayplan.Attraction
// 	var shortestBreakfastdistance float64

// 	for _, breakfastLocation := range breakfastLocations {
// 		if shortestBreakfastdistance < 1 {
// 			shortestBreakfastdistance = geo.Haversine(breakfastLocation.Lat, breakfastLocation.Lng, plan.HomeLocation.Lat, plan.HomeLocation.Lng)
// 			shortestBreakfastLocation = breakfastLocation

// 			continue
// 		}

// 		distance := geo.Haversine(breakfastLocation.Lat, breakfastLocation.Lng, plan.HomeLocation.Lat, plan.HomeLocation.Lng)

// 		if distance < shortestBreakfastdistance {
// 			shortestBreakfastdistance = distance
// 			shortestBreakfastLocation = breakfastLocation
// 		}
// 	}

// 	plan.AddAttraction(shortestBreakfastLocation)
// }

// func planAttractions(attractions []dayplan.Attraction, plan *dayplan.DayPlan) {
// 	if plan.IsActivityFull() || len(attractions) == 0 {
// 		return
// 	}

// 	var closestAttractionDistance float64
// 	var closestAttractionIndex int
// 	var closestAttraction dayplan.Attraction

// 	lastAttraction := plan.GetLastAttractionLocation()

// 	for index, attraction := range attractions {
// 		if closestAttractionDistance == 0 {
// 			closestAttractionDistance = geo.Haversine(attraction.Lat, attraction.Lng, lastAttraction.Lat, lastAttraction.Lng)
// 			closestAttraction = attraction
// 			closestAttractionIndex = index

// 			continue
// 		}

// 		distance := geo.Haversine(attraction.Lat, attraction.Lng, lastAttraction.Lat, lastAttraction.Lng)

// 		if distance < closestAttractionDistance {
// 			closestAttractionDistance = distance
// 			closestAttraction = attraction
// 			closestAttractionIndex = index
// 		}
// 	}

// 	closestAttraction.DistanceToPrevious = closestAttractionDistance

// 	plan.AddAttraction(closestAttraction)

// 	// remove attraction from slice
// 	attractions = append(attractions[:closestAttractionIndex], attractions[closestAttractionIndex+1:]...)

// 	planAttractions(attractions, plan)
// }
