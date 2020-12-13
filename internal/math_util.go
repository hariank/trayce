package util

import (
	"math"
	"math/rand"
)

func DegreesToRadians(degrees float64) float64 {
	return degrees * math.Pi / 180.
}

func RandomFloat() float64 {
	return rand.Float64()
}

func RandomFloatIn(lowerBound, upperBound float64) float64 {
	return RandomFloat()*(upperBound-lowerBound) + lowerBound
}

func Clamp(x, lowerBound, upperBound float64) float64 {
	if x < lowerBound {
		return lowerBound
	} else if x > upperBound {
		return upperBound
	} else {
		return x
	}
}
