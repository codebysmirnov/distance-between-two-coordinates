package main

import (
	"fmt"
	"math"
)

// Point represents coordinates
type Point struct {
	Latitude  float64
	Longitude float64
}

const (
	RadiansPerDegree = 3.1415927 / 180
	EarthRadius      = 6371000
)

func main() {
	from := Point{
		Latitude:  40.412279,
		Longitude: -3.676257,
	}
	to := Point{
		Latitude:  40.408335,
		Longitude: -3.673202,
	}

	phiM := (from.Latitude + to.Latitude) / 2
	deltaPhi := math.Abs(to.Latitude - from.Latitude)
	deltaLambda := math.Abs(to.Longitude - from.Longitude)

	firstTerm := math.Pow(deltaPhi, 2)
	secondTerm := math.Pow(math.Cos(phiM*RadiansPerDegree)*deltaLambda, 2)
	// result in meters
	fmt.Println(int(EarthRadius * RadiansPerDegree * math.Sqrt(firstTerm+secondTerm)))
}
