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
	points := make([]*LatLng, 4)
	points[0] = &LatLng{lat: 59.19305333867669, lng: 17.661922238767147}
	points[1] = &LatLng{lat: 59.192982176318765, lng: 17.662122901529074}
	points[2] = &LatLng{lat: 59.19288511388004, lng: 17.66255029477179}
	points[3] = &LatLng{lat: 59.19290036894381, lng: 17.662896132096648}
	result := GetTotalDistance(points)
	if result != expected {
		t.Errorf("Expected '%v' but got '%v'", expected, result)
	}
}

func TestGetTotalDistanceZero(t *testing.T) {
	expected := 0.0
	points := make([]*LatLng, 0)
	result := GetTotalDistance(points)
	if result != expected {
		t.Errorf("Expected '%v' but got '%v'", expected, result)
	}
}
func TestGetMidPointNone(t *testing.T) {
	points := make([]*LatLng, 0)

	_, err := GetMidPoint(points)
	if err == nil {
		t.Errorf("Expected error but got nil")
	}

}

func TestGetMidPointOne(t *testing.T) {
	expected := &LatLng{lat: 59.19290036894381, lng: 17.662896132096648}
	points := make([]*LatLng, 1)
	points[0] = &LatLng{lat: 59.19290036894381, lng: 17.662896132096648}

	result, _ := GetMidPoint(points)
	if result.Lng() != expected.Lng() {
		t.Errorf("Expected '%v' but got '%v'", expected.Lng(), result.Lng())
	}

	if result.Lat() != expected.Lat() {
		t.Errorf("Expected '%v' but got '%v'", expected.Lat(), result.Lat())
	}
}

func TestGetMidPointSamePoint(t *testing.T) {
	expected := &LatLng{lat: 59.1929003689438, lng: 17.662896132096648}
	points := make([]*LatLng, 4)
	points[0] = &LatLng{lat: 59.1929003689438, lng: 17.662896132096648}
	points[1] = &LatLng{lat: 59.1929003689438, lng: 17.662896132096648}
	points[2] = &LatLng{lat: 59.1929003689438, lng: 17.662896132096648}
	points[3] = &LatLng{lat: 59.1929003689438, lng: 17.662896132096648}

	result, _ := GetMidPoint(points)
	if result.Lng() != expected.Lng() {
		t.Errorf("Expected '%v' but got '%v'", expected.Lng(), result.Lng())
	}

	if result.Lat() != expected.Lat() {
		t.Errorf("Expected '%v' but got '%v'", expected.Lat(), result.Lat())
	}
}

func TestGetMidPointMultiPoints(t *testing.T) {
	expected := &LatLng{lat: 59.3000383089087, lng: 17.299706052455726}
	points := make([]*LatLng, 2)
	points[0] = &LatLng{lat: 59.2, lng: 17.2}
	points[1] = &LatLng{lat: 59.4, lng: 17.4}

	result, _ := GetMidPoint(points)
	if result.Lng() != expected.Lng() {
		t.Errorf("Expected '%v' but got '%v'", expected.Lng(), result.Lng())
	}

	if result.Lat() != expected.Lat() {
		t.Errorf("Expected '%v' but got '%v'", expected.Lat(), result.Lat())
	}
}

func TestGetBoundingBox(t *testing.T) {
	// approximately 1 decimal degree is 111319.9
	// those value taken from output :-)
	// 1.0
	expected := &BBox{
		northeast: LatLng{lat: 1.001123912387125, lng: -48.33220908761288},
		southwest: LatLng{lat: -1.001123912387125, lng: -50.334456912387125}}

	result := GetBoundingBox(0, -49.333333, 111319.9)
	if result.Northeast().Lat() != expected.Northeast().Lat() {
		t.Errorf("Expcted '%v' but got '%v'", expected.Northeast().Lat(), result.Northeast().Lat())
	}

	if result.Northeast().Lng() != expected.Northeast().Lng() {
		t.Errorf("Expcted '%v' but got '%v'", expected.Northeast().Lng(), result.Northeast().Lng())
	}

	if result.Southwest().Lat() != expected.Southwest().Lat() {
		t.Errorf("Expcted '%v' but got '%v'", expected.Southwest().Lat(), result.Southwest().Lat())
	}

	if result.Southwest().Lng() != expected.Southwest().Lng() {
		t.Errorf("Expcted '%v' but got '%v'", expected.Southwest().Lng(), result.Southwest().Lng())
	}

	// 0.1 => 11132 meters
	expected = &BBox{
		northeast: LatLng{lat: 0.10011248117087308, lng: -49.23322051882913},
		southwest: LatLng{lat: -0.10011248117087308, lng: -49.43344548117087}}

	result = GetBoundingBox(0, -49.333333, 11132.0)
	if result.Northeast().Lat() != expected.Northeast().Lat() {
		t.Errorf("Expcted '%v' but got '%v'", expected.Northeast().Lat(), result.Northeast().Lat())
	}

	if result.Northeast().Lng() != expected.Northeast().Lng() {
		t.Errorf("Expcted '%v' but got '%v'", expected.Northeast().Lng(), result.Northeast().Lng())
	}

	if result.Southwest().Lat() != expected.Southwest().Lat() {
		t.Errorf("Expcted '%v' but got '%v'", expected.Southwest().Lat(), result.Southwest().Lat())
	}

	if result.Southwest().Lng() != expected.Southwest().Lng() {
		t.Errorf("Expcted '%v' but got '%v'", expected.Southwest().Lng(), result.Southwest().Lng())
	}
}
