package gpsutil

func GetDistance(lng1, lat1, lng2, lat2 float64) float64 {

	return 0.0
}

func GetTotalDistance(points []LatLng) float64 {
	max := len(points) - 1
	total := 0.0
	for i := 0; i < max; i++ {
		total += GetDistance(points[i].lng, points[i].lat, points[i+1].lng, points[i+1].lat)
	}

	return total
}

func GetBoundingBox(lat, lng, distance float64) BBox {
	return nil
}

func GetMidPoint(points []LatLng) (LatLng, error) {
	return nil
}
