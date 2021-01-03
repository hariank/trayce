package geometry

import (
	"math"
)

type Sphere struct {
	Center   Point
	Radius   float64
	Material Material
}

func NewSphere(c Point, ray float64, m Material) *Sphere {
	return &Sphere{Center: c, Radius: ray, Material: m}
}

func (s *Sphere) Hit(ray *Ray, tMin float64, tMax float64) (bool, *HitRecord) {
	oc := ray.Orig.Minus(s.Center)
	a := ray.Dir.MagSquared()
	hb := oc.Dot(ray.Dir)
	c := oc.MagSquared() - s.Radius*s.Radius
	discriminant := hb*hb - a*c
	if discriminant < 0 {
		return false, nil
	}
	sqrtDiscrim := math.Sqrt(discriminant)

	// find nearest root s.t. t in [tMin, tMax]
	root := (-hb - sqrtDiscrim) / a
	if root < tMin || root > tMax {
		root := (-hb + sqrtDiscrim) / a
		if root < tMin || root > tMax {
			return false, nil
		}
	}

	hitLoc := ray.At(root)
	record := &HitRecord{Loc: hitLoc, T: root, Material: s.Material}
	norm := hitLoc.Minus(s.Center).Scale(1. / s.Radius)
	record.SetFaceNormal(ray, norm)
	return true, record
}
