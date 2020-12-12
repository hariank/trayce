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
