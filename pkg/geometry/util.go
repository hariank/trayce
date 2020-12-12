package geometry

import (
	"math"
	"math/rand"
)

func degreesToRadians(degrees float64) float64 {
	return degrees * math.Pi / 180.
}

func randomFloat() float64 {
	return rand.Float64()
}

func randomFloatIn(lowerBound, upperBound float64) float64 {
	return randomFloat()*(upperBound-lowerBound) + lowerBound
}
