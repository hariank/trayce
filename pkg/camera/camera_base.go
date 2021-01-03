package camera

import (
	geo "github.com/hariank/trace/pkg/geometry"
)

const ()

type Camera interface {
	GenerateRay(float64, float64) *geo.Ray
}

type PinholeCamera struct {
	AspectRatio, ViewportHeight, ViewportWidth float64
	FocalLength                                float64
	origin                                     geo.Vec
	vpHorizontal, vpVertical, vpLowerLeft      geo.Vec
}

func NewPinholeCamera(aspectRatio float64) *PinholeCamera {
	ViewportHeight := 2.
	ViewportWidth := ViewportHeight * aspectRatio
	FocalLength := 1.
	origin := geo.Origin
	vpHorizontal := geo.Vec{ViewportWidth, 0, 0}
	vpVertical := geo.Vec{0, ViewportHeight, 0}
	vpDepth := geo.Vec{0, 0, FocalLength}
	vpLowerLeft := origin.Minus(vpHorizontal.Scale(.5)).Minus(vpVertical.Scale(.5)).Minus(vpDepth)
	return &PinholeCamera{
		AspectRatio:    aspectRatio,
		ViewportHeight: ViewportHeight,
		ViewportWidth:  ViewportWidth,
		FocalLength:    FocalLength,
		origin:         origin,
		vpHorizontal:   vpHorizontal,
		vpVertical:     vpVertical,
		vpLowerLeft:    vpLowerLeft,
	}
}

func (c *PinholeCamera) GenerateRay(u, v float64) *geo.Ray {
	hPos := c.vpHorizontal.Scale(u)
	vPos := c.vpVertical.Scale(v)
	return geo.NewRay(c.origin, c.vpLowerLeft.Plus(hPos).Plus(vPos).Minus(c.origin))
}
