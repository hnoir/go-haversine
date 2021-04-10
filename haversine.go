// Package haversine provides Haversine distance calculations.
package haversine

import (
	"errors"
	"math"
)

const (
	// Earth radius (km).
	earthRadius = 6371.0

	metersPerMile         = 1609.344
	metersPerNauticalMile = 1852

	degreesToRadians = math.Pi / 180.0

	minLatitude  = -90.0
	minLongitude = -180.0

	maxLatitude  = 90.0
	maxLongitude = 180.0
)

var (
	ErrInvalidLatitude  = errors.New("invalid latitude, expected value between -90 and 90 degrees")
	ErrInvalidLongitude = errors.New("invalid longitude, expected value between -180 and 180 degrees")
)

func isValidLatitude(lat float64) bool {
	return lat >= minLatitude && lat <= maxLatitude
}

func isValidLongitude(lng float64) bool {
	return lng >= minLongitude && lng <= maxLongitude
}

// Distance calculates the distance between two points on a sphere given their longitudes and latitudes.
func Distance(lat1, lng1, lat2, lng2 float64) (float64, error) {
	// Check if the latitude and longitude inputs are in valid ranges.
	if !isValidLatitude(lat1) || !isValidLatitude(lat2) {
		return 0, ErrInvalidLatitude
	}
	if !isValidLongitude(lng1) || !isValidLongitude(lng2) {
		return 0, ErrInvalidLongitude
	}

	// Convert coordinates to radians.
	latitude1Rad := lat1 * degreesToRadians
	longitude1Rad := lng1 * degreesToRadians
	latitude2Rad := lat2 * degreesToRadians
	longitude2Rad := lng2 * degreesToRadians

	// Haversine formula.
	deltaLatitude := latitude2Rad - latitude1Rad
	deltaLongitude := longitude2Rad - longitude1Rad
	squaredHalfChordLength := math.Pow(math.Sin(deltaLatitude/2), 2) + math.Cos(latitude1Rad)*math.Cos(latitude2Rad)*math.Pow(math.Sin(deltaLongitude/2), 2)

	angularDistance := 2 * math.Atan2(math.Sqrt(squaredHalfChordLength), math.Sqrt(1-squaredHalfChordLength))
	distance := earthRadius * angularDistance

	return distance, nil
}

// DistanceMi calculates the distance between two points on a sphere given their longitudes and latitudes in miles.
func DistanceMi(lat1, lng1, lat2, lng2 float64) (float64, error) {
	distance, err := Distance(lat1, lng1, lat2, lng2)
	if err != nil {
		return 0, err
	}
	return distance / metersPerMile, nil
}

// DistanceNMi calculates the distance between two points on a sphere given their longitudes and latitudes in nautical miles.
func DistanceNMi(lat1, lng1, lat2, lng2 float64) (float64, error) {
	distance, err := Distance(lat1, lng1, lat2, lng2)
	if err != nil {
		return 0, err
	}
	return distance / metersPerNauticalMile, nil
}
