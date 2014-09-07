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
	radDist := distance / EARTH_RADIUS
	radLat := toRad(lat)
	radLng := toRad(lng)

	var minLat, maxLat, minLng, maxLng, deltaLng float64

	minLat = radLat - radDist
	maxLat = radLat + radDist

	if minLat > MIN_LAT && maxLat < MAX_LAT {
		deltaLng = math.Asin(math.Sin(radDist) / math.Cos(radLat))
		minLng = radLng - deltaLng
		if minLng < MIN_LNG {
			minLng += 2 * math.Pi
		}
		maxLng = radLng + deltaLng
		if maxLng > MAX_LNG {
			maxLng -= 2 * math.Pi
		}
	} else {
		minLat = math.Max(minLat, MIN_LAT)
		maxLat = math.Min(maxLat, MAX_LAT)
		minLng = MIN_LNG
		maxLng = MAX_LNG
	}

	return &BBox{
		southwest: LatLng{lat: toDegrees(minLat), lng: toDegrees(minLng)},
		northeast: LatLng{lat: toDegrees(maxLat), lng: toDegrees(maxLng)}}
}

func GetMidPoint(points []*LatLng) (*LatLng, error) {
	length := len(points)
	if length < 1 {
		return nil, fmt.Errorf("Points must not be empty")
	} else if length == 1 {
		return &LatLng{lat: points[0].lat, lng: points[0].lng}, nil
	}
	x := 0.0
	y := 0.0
	z := 0.0
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
