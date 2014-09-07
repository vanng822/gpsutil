package gpsutil

import (
	"math"
)

type LatLng struct {
	lat float64
	lng float64
}

func (latlng *LatLng) Lat() float64 {
	return latlng.lat
}

func (latlng *LatLng) Lng() float64 {
	return latlng.lng
}

type GeohashDecoded struct {
	lat    float64
	lng    float64
	latErr float64
	lgnErr float64
}

func (ghd *GeohashDecoded) Lat() float64 {
	return ghd.lat
}

func (ghd *GeohashDecoded) Lng() float64 {
	return ghd.lng
}

func (ghd *GeohashDecoded) LatErr() float64 {
	return ghd.latErr
}

func (ghd *GeohashDecoded) LngErr() float64 {
	return ghd.lgnErr
}

type BBox struct {
	southwest LatLng
	northeast LatLng
}

func (bbox *BBox) Southwest() *LatLng {
	return &bbox.southwest
}

func (bbox *BBox) Northeast() *LatLng {
	return &bbox.northeast
}

func toRad(decDegrees float64) float64 {
	return decDegrees * math.Pi / 180.0
}

func toDegrees(radians float64) float64 {
	return 180.0 * radians / math.Pi
}
