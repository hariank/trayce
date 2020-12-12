package geometry

import (
	"fmt"
	"math"
)

type Vec struct {
	X, Y, Z float64
}

var Origin = Vec{0, 0, 0}

func (v Vec) String() string {
	return fmt.Sprintf("%f %f %f", v.X, v.Y, v.Z)
}

func (v Vec) Plus(u Vec) Vec {
	return Vec{v.X + u.X, v.Y + u.Y, v.Z + u.Z}
}

func (v Vec) Minus(u Vec) Vec {
	return Vec{v.X - u.X, v.Y - u.Y, v.Z - u.Z}
}

func (v Vec) Times(u Vec) Vec {
	return Vec{v.X * u.X, v.Y * u.Y, v.Z * u.Z}
}

func (v Vec) Dot(u Vec) float64 {
	return v.X*u.X + v.Y*u.Y + v.Z*u.Z
}

func (v Vec) Cross(u Vec) Vec {
	return Vec{(v.Y*u.Z - v.Z*u.Y), (v.Z*u.X - v.X*u.Z), (v.X*u.Y - v.Y*u.X)}
}

func (v Vec) Scale(c float64) Vec {
	return Vec{v.X * c, v.Y * c, v.Z * c}
}
func (v Vec) ScaleInt(c int) Vec {
	cf := float64(c)
	return Vec{v.X * cf, v.Y * cf, v.Z * cf}
}

func (v Vec) MagSquared() float64 {
	return (v.X*v.X + v.Y*v.Y + v.Z*v.Z)
}
func (v Vec) Mag() float64 {
	return math.Sqrt(v.MagSquared())
}

func (v Vec) Unit() Vec {
	return v.Scale(1. / v.Mag())
}

func (v Vec) Flip() Vec {
	return v.Scale(-1.)
}

type Point = Vec
