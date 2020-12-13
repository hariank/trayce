package geometry

import (
	"math"
)

type HitRecord struct {
	Loc       Point
	Norm      Vec // not necessarily unit
	T         float64
	FrontFace bool // if hits the front face of object
}

// ensure the Norm is always against the ray
func (hr *HitRecord) SetFaceNormal(ray *Ray, outwardNorm Vec) {
	hr.FrontFace = bool(ray.Dir.Dot(outwardNorm) < 0)
	if hr.FrontFace {
		hr.Norm = outwardNorm
	} else {
		hr.Norm = outwardNorm.Flip()
	}
}

type Hittable interface {
	Hit(*Ray, float64, float64) (bool, *HitRecord)
}

type HittableList struct {
	Objects []Hittable
}

func NewHittableList() HittableList { // TODO: debug this
	return HittableList{}
}

// TODO: fix this
func (hl HittableList) AddHittables(h ...Hittable) {
	hl.Objects = append(hl.Objects, h...)
}

func (hl HittableList) Hit(ray *Ray, tMin float64, tMax float64) (bool, *HitRecord) {
	record := &HitRecord{}
	hitAnything := false
	closestT := tMax

	for _, h := range hl.Objects {
		if hit, curRecord := h.Hit(ray, tMin, closestT); hit {
			hitAnything = true
			closestT = curRecord.T
			record = curRecord
		}
	}

	return hitAnything, record
}

type Sphere struct {
	Center Point
	Radius float64
}

func NewSphere(c Point, ray float64) *Sphere {
	return &Sphere{Center: c, Radius: ray}
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
	record := &HitRecord{Loc: hitLoc, T: root}
	norm := hitLoc.Minus(s.Center).Scale(1. / s.Radius)
	record.SetFaceNormal(ray, norm)
	return true, record
}
