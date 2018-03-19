package latlong

import (
	"math"
)

// R is the mean radius of Earth (6,371km)
const R = 6371300

// toRadians converts a position from degrees to radians
func toRadians(pos float64) float64 {
	return (math.Pi * pos / 180)
}

// EquirectangularApproxDistance computes the distance between two points on Earth using
// EquirectangularApprox formula.
// More details: https://www.movable-type.co.uk/scripts/latlong.html
func EquirectangularApproxDistance(lat1, long1, lat2, long2 float64) float64 {
	φ1 := toRadians(lat1)
	φ2 := toRadians(lat2)

	Δφ := toRadians(lat2 - lat1)
	Δλ := toRadians(long2 - long1)

	x := Δλ * math.Cos(float64(φ1+φ2)/2.0)
	y := Δφ

	d := R * math.Sqrt((x*x)+(y*y))
	return d
}

// HaversineDistance computes the distance between two points on Earth using
// EquirectangularApprox formula.
// More details: https://www.movable-type.co.uk/scripts/latlong.html
func HaversineDistance(lat1, long1, lat2, long2 float64) float64 {
	φ1 := toRadians(lat1)
	φ2 := toRadians(lat2)

	Δφ := toRadians(math.Abs(lat2 - lat1))
	Δλ := toRadians(math.Abs(long2 - long1))

	// a = sin²(Δφ/2) + cos φ1 ⋅ cos φ2 ⋅ sin²(Δλ/2)
	a := math.Sin(Δφ/2)*math.Sin(Δφ/2) +
		math.Cos(φ1)*math.Cos(φ2)*
			math.Sin(Δλ/2)*math.Sin(Δλ/2)

	// c = 2 ⋅ asin( min(1, √a) )
	c := 2 * math.Asin(math.Min(1, math.Sqrt(a)))

	d := R * c
	return d
}
