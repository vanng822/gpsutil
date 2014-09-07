package gpsutil

import (
	"bytes"
)

/**
 * http://en.wikipedia.org/wiki/Geohash
 */
var (
	base32 = []byte("0123456789bcdefghjkmnpqrstuvwxyz")
	bits = []int{16, 8, 4, 2, 1}
)

func GeohashEncode(lat, lng float64, precision int) string {
	if precision < 1 {
		precision = 12
	}
	maxLat := float64(90)
	minLat := float64(-90)
	maxLng := float64(180)
	minLng := float64(-180)

	var mid float64
	var hashPos, bit int
	var geohash bytes.Buffer
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
		geohash.WriteByte(base32[hashPos])
		precision -= 1
	}
	return geohash.String()
}

func GeohashDecode(hash string) (*GeohashDecoded, error) {
	return &GeohashDecoded{}, nil
}
