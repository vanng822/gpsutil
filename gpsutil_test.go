package gpsutil

import (
	"testing"
	"math"
)


func TestToRad(t *testing.T) {
	expected := math.Pi/2
	result := toRad(90)
	if result != expected {
		t.Errorf("Expected '%v' but got '%v'", expected, result)
	}
	
	expected = math.Pi/4
	result = toRad(45)
	if result != expected {
		t.Errorf("Expected '%v' but got '%v'", expected, result)
	}
	
	expected = math.Pi * 2
	result = toRad(360)
	if result != expected {
		t.Errorf("Expected '%v' but got '%v'", expected, result)
	}
}