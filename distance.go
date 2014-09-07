package gpsutil

import (
	"fmt"
	"math"
)

func GetDistance(lng1, lat1, lng2, lat2 float64) float64 {
	dLat := toRad(lat2 - lat1)
	dLng := toRad(lng2 - lng1)
	a := math.Pow(math.Sin(dLat/2), 2) + math.Pow(math.Sin(dLng/2), 2)*math.Cos(toRad(lat1))*math.Cos(toRad(lat2))
	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))
	return c * EARTH_RADIUS
}

func GetTotalDistance(points []*LatLng) float64 {
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

func GetMidPoint(points []*LatLng) (*LatLng, error) {
	length := len(points)
	x := 0.0
	y := 0.0
	z := 0.0

	if length < 1 {
		return nil, fmt.Errorf("Points must not be empty")
	} else if length == 1 {
		return &LatLng{lat: points[0].lat, lng: points[0].lng}, nil
	}
	var lat, lng float64

	for i := 0; i < length; i++ {
		lat = toRad(points[i].lat)
		lng = toRad(points[i].lng)
		x += math.Cos(lat) * math.Cos(lng)
		y += math.Cos(lat) * math.Sin(lng)
		z += math.Sin(lat)
	}

	x = x / float64(length)
	y = y / float64(length)
	z = z / float64(length)

	lng = math.Atan2(y, x)
	lat = math.Atan2(z, math.Sqrt(x*x+y*y))

	return &LatLng{lat: toDegrees(lat), lng: toDegrees(lng)}, nil
}
