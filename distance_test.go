package latlong

import (
	"fmt"
	"testing"
)

var lat1 = -6.2174552
var long1 = 106.8353401

var lat2 = -6.22102
var long2 = 106.8170328

func TestDistance(t *testing.T) {
	fmt.Println(EquirectangularApproxDistance(lat2, long2, lat1, long1))
	fmt.Println(HaversineDistance(lat1, long1, lat2, long2))
}

func BenchmarkEquirectangularApproxDistance(b *testing.B) {
	for i := 0; i < b.N; i++ {
		EquirectangularApproxDistance(lat1, long1, lat2, long2)
	}
}

func BenchmarkHaversineDistance(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HaversineDistance(lat1, long1, lat2, long2)
	}
}
