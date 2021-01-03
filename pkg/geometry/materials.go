package geometry

type Material interface {
	ScatterRay(*Ray, *HitRecord) (Color, *Ray)
}

type Lambertian struct {
	Albedo Color
}

type Hemispherical struct {
	Albedo Color
}

func (l *Lambertian) ScatterRay(r *Ray, hitRecord *HitRecord) (Color, *Ray) {
	scatterDir := hitRecord.Norm.Plus(RandomVecOnUnitSphere())
	if scatterDir.NearZero() { // don't scatter if degenerate
		scatterDir = hitRecord.Norm
	}
	scatterRay := NewRay(hitRecord.Loc, scatterDir)
	return l.Albedo, scatterRay
}

func (h *Hemispherical) ScatterRay(r *Ray, hitRecord *HitRecord) (Color, *Ray) {
	scatterDir := hitRecord.Norm.Plus(RandomVecInHemisphere(hitRecord.Norm))
	if scatterDir.NearZero() { // don't scatter if degenerate
		scatterDir = hitRecord.Norm
	}
	scatterRay := NewRay(hitRecord.Loc, scatterDir)
	return h.Albedo, scatterRay
}
