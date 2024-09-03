package geo

import "math"

func Haversine(lat1, lng1, lat2, lng2 float64) float64 {
	const R = 6371 // Radius of the Earth in kilometers

	// Convert degrees to radians
	lat1Rad := lat1 * (math.Pi / 180)
	lng1Rad := lng1 * (math.Pi / 180)
	lat2Rad := lat2 * (math.Pi / 180)
	lng2Rad := lng2 * (math.Pi / 180)

	// Haversine formula
	dLat := lat2Rad - lat1Rad
	dLng := lng2Rad - lng1Rad

	a := math.Sin(dLat/2)*math.Sin(dLat/2) +
		math.Cos(lat1Rad)*math.Cos(lat2Rad)*
			math.Sin(dLng/2)*math.Sin(dLng/2)

	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))

	// Distance in kilometers
	return R * c
}
