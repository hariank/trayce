package geometry

import (
	"math"
)

type Color = Vec

func Gamma2Correct(color Color) Color {
	return Color{math.Sqrt(color.X), math.Sqrt(color.Y), math.Sqrt(color.Z)}
}

// assuming n is unit length, map components from [-1,1] to [0,1]
func VisualizeNormal(n Vec) Color {
	return n.Plus(Vec{1, 1, 1}).Scale(0.5)
}
