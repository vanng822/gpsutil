package gpsutil

import (
	"testing"
)

func TestGetDistance(t *testing.T) {
	expected := 13.899604253423052
	result := GetDistance(17.661922238767147, 59.19305333867669, 17.662122901529074, 59.192982176318765)
	if result != expected {
		t.Errorf("Expected '%v' but got '%v'", expected, result)
	}
}

func TestGetTotalDistance(t *testing.T) {
	expected := 60.29237674121895
	points := make([]LatLng, 4)
	points[0] = LatLng{lat: 59.19305333867669, lng: 17.661922238767147}
	points[1] = LatLng{lat:59.192982176318765 , lng: 17.662122901529074}
	points[2] = LatLng{lat: 59.19288511388004, lng: 17.66255029477179}
	points[3] = LatLng{lat: 59.19290036894381, lng: 17.662896132096648}
	result := GetTotalDistance(points)
	if result != expected {
		t.Errorf("Expected '%v' but got '%v'", expected, result)
	}
}

func TestGetTotalDistanceZero(t *testing.T) {
	expected := 0.0
	points := make([]LatLng, 0)
	result := GetTotalDistance(points)
	if result != expected {
		t.Errorf("Expected '%v' but got '%v'", expected, result)
	}
}