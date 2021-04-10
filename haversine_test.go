package haversine

import (
	"math"
	"testing"
)

func TestDistance(t *testing.T) {
	// Define test cases with expected results.
	testCases := []struct {
		name             string
		lat1             float64
		lng1             float64
		lat2             float64
		lng2             float64
		expectedDistance float64
		expectedError    error
	}{
		{
			name:             "Amsterdam to Paris",
			lat1:             52.3676,
			lng1:             4.9041,
			lat2:             48.8566,
			lng2:             2.3522,
			expectedDistance: 429.86,
			expectedError:    nil,
		},
		{
			name:             "Rome to Venice",
			lat1:             41.9028,
			lng1:             12.4964,
			lat2:             45.4408,
			lng2:             12.3155,
			expectedDistance: 393.67,
			expectedError:    nil,
		},
		{
			name:             "New York to Zurich",
			lat1:             40.7128,
			lng1:             -74.0060,
			lat2:             47.3769,
			lng2:             8.5417,
			expectedDistance: 6323.76,
			expectedError:    nil,
		},
		{
			name:             "Invalid latitude",
			lat1:             -100,
			lng1:             -118.2437,
			lat2:             40.7128,
			lng2:             -74.0060,
			expectedDistance: 0,
			expectedError:    ErrInvalidLatitude,
		},
		{
			name:             "Invalid longitude",
			lat1:             34.0522,
			lng1:             -190,
			lat2:             40.7128,
			lng2:             -74.0060,
			expectedDistance: 0,
			expectedError:    ErrInvalidLongitude,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			distance, err := Distance(tc.lat1, tc.lng1, tc.lat2, tc.lng2)

			if err != tc.expectedError {
				t.Errorf("Expected error %v, but got %v", tc.expectedError, err)
			}

			if !math.IsNaN(tc.expectedDistance) && math.Abs(distance-tc.expectedDistance) > 1.00 {
				t.Errorf("Expected distance %v, but got %v", tc.expectedDistance, distance)
			}
		})
	}
}
