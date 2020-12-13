package color

import (
	geo "github.com/hariank/trace/pkg/geometry"
)

type Material interface {
	ScatterRay(*geo.Ray, *geo.HitRecord) (Color, *Ray)
}

func LambertianDiffuse(hitRecord *geo.HitRecord) geo.Vec {
	diffuseTarget := hitRecord.Loc.Plus(hitRecord.Norm).Plus(geo.RandomVecOnUnitSphere())
	return diffuseTarget.Minus(hitRecord.Loc)
}
func HemisphericalDiffuse(hitRecord *geo.HitRecord) geo.Vec {
	diffuseTarget := hitRecord.Loc.Plus(geo.RandomVecInHemisphere(hitRecord.Norm))
	return diffuseTarget.Minus(hitRecord.Loc)
}
