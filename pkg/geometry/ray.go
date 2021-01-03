package geometry

type Ray struct {
	Orig Point
	Dir  Vec
}

func NewRay(p Point, d Vec) *Ray {
	return &Ray{Orig: p, Dir: d}
}

func (r *Ray) At(t float64) Point {
	return r.Orig.Plus(r.Dir.Scale(t))
}

type HitRecord struct {
	Loc       Point
	Norm      Vec // not necessarily unit
	T         float64
	FrontFace bool     // if hits the front face of object
	Material  Material // the material we hit
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
