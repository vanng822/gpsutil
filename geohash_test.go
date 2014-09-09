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

func TestGeohashEncodeDefaultPrecision(t *testing.T) {
	expected := "u4pruydqqvj8"
	result := GeohashEncode(57.64911,10.40744, 0)
	if expected != result {
		t.Errorf("Expected '%s' but got '%s'", expected, result)
	}
}

func TestGeohashDecode(t *testing.T) {
	result, _ := GeohashDecode("u4pruydqqvj")
	
	expected := &GeohashDecoded{lat: 57.64911063015461, lng: 10.407439693808556, latErr: 6.705522537231445e-07, lngErr: 6.705522537231445e-07}
	
	if expected.Lat() != result.Lat() {
		t.Errorf("Expected '%s' but got '%s'", expected.Lat(), result.Lat())
	}

	if expected.Lng() != result.Lng() {
		t.Errorf("Expected '%s' but got '%s'", expected.Lng(), result.Lng())
	}
	
	if expected.LngErr() != result.LngErr() {
		t.Errorf("Expected '%s' but got '%s'", expected.LngErr(), result.LngErr())
	}
	
	if expected.LatErr() != result.LatErr() {
		t.Errorf("Expected '%s' but got '%s'", expected.LatErr(), result.LatErr())
	}
	
}

func TestGeohashDecodeInvalidChar(t *testing.T) {
	result, err := GeohashDecode("u4pot")
	
	if  result != nil {
		t.Errorf("Result expected to be nil")
	}

	if err == nil {
		t.Errorf("Expected an error")
		t.FailNow()
	}
	expected := "Character 'o' doesn't be a part of base32"
	if err.Error() != expected {
		t.Errorf("Expected '%s' but got '%s'", expected, err.Error())
	}
	
}