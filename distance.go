package gpsutil

import (
	"math"
)

func GetDistance(lng1, lat1, lng2, lat2 float64) float64 {
	dLat := toRad(lat2 - lat1)
	dLng := toRad(lng2 - lng1)
	a := math.Sin(dLat/2)*math.Sin(dLat/2) + math.Sin(dLng/2)*math.Sin(dLng/2)*math.Cos(toRad(lat1))*math.Cos(toRad(lat2))
	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))
	return c * EARTH_RADIUS
}

func GetTotalDistance(points []LatLng) float64 {
	max := len(points) - 1
	total := 0.0
	for i := 0; i < max; i++ {
		total += GetDistance(points[i].lng, points[i].lat, points[i+1].lng, points[i+1].lat)
	}

	return total
}

func GetBoundingBox(lat, lng, distance float64) *BBox {
	return &BBox{}
}

func GetMidPoint(points []LatLng) (*LatLng, error) {
	return nil, nil
}
