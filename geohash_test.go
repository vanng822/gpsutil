package gpsutil


import (
	"testing"
)


func TestGeohashEncode(t *testing.T) {
	expected := "u4pruydqqvj"
	result := GeohashEncode(57.64911,10.40744, 11)
	if expected != result {
		t.Errorf("Expected '%s' but got '%s'", expected, result)
	}
}