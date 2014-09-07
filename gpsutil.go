package gpsutil

import (
	"math"
)

type LatLng struct {
	lat float64
	lng float64
}

type GeohashDecoded struct {
	lat float64
	lng float64
	err struct {
		lat float64
		lgn float64
	}
}

type BBox struct {
	Southwest *LatLng
	Northeast *LatLng
	Center    *LatLng
}

func toRad(decDegrees float64) float64 {
	return decDegrees * math.Pi / 180.0
}

func toDegrees(radians float64) float64 {
	return 180.0 * radians / math.Pi
}
