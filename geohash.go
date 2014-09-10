package gpsutil

import (
	"bytes"
	"fmt"
)

/**
 * http://en.wikipedia.org/wiki/Geohash
 */
var (
	base32 = []byte("0123456789bcdefghjkmnpqrstuvwxyz")
	bits   = []int{16, 8, 4, 2, 1}
)

func GeohashEncode(lat, lng float64, precision int) string {
	if precision < 1 {
		precision = 12
	}
	maxLat := 90.0
	minLat := -90.0
	maxLng := 180.0
	minLng := -180.0

	var mid float64
	var hashPos, bit int
	var hash bytes.Buffer
	even := true
	for precision > 0 {
		hashPos = 0
		for bit = 0; bit < 5; bit++ {
			if even {
				mid = (maxLng + minLng) / 2
				if lng > mid {
					hashPos |= bits[bit]
					minLng = mid
				} else {
					maxLng = mid
				}
			} else {
				mid = (maxLat + minLat) / 2
				if lat > mid {
					hashPos |= bits[bit]
					minLat = mid
				} else {
					maxLat = mid
				}
			}
			even = !even
		}
		hash.WriteByte(base32[hashPos])
		precision -= 1
	}
	return hash.String()
}

func GeohashDecode(hash string) (*GeohashDecoded, error) {
	maxLat := 90.0
	minLat := -90.0
	maxLng := 180.0
	minLng := -180.0

	var bit, hashPos int
	var lat, lng, mid float64

	length := len(hash)
	even := true
	var c string

	for i := 0; i < length; i++ {
		c = hash[i : i+1]
		hashPos = bytes.Index(base32, []byte(c))
		if hashPos == -1 {
			return nil, fmt.Errorf("Character '%s' doesn't be a part of base32", c)
		}
		for bit = 4; bit >= 0; bit-- {
			if even {
				mid = (maxLng + minLng) / 2.0
				if ((uint(hashPos) >> uint(bit)) & 1) == 1 {
					minLng = mid
				} else {
					maxLng = mid
				}
			} else {
				mid = (maxLat + minLat) / 2.0
				if ((uint(hashPos) >> uint(bit)) & 1) == 1 {
					minLat = mid
				} else {
					maxLat = mid
				}
			}
			even = !even
		}
	}
	lat = (minLat + maxLat) / 2.0
	lng = (minLng + maxLng) / 2.0
	return &GeohashDecoded{
		lat:    lat,
		lng:    lng,
		latErr: maxLat - lat,
		lngErr: maxLng - lng}, nil
}
