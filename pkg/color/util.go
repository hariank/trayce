package color

import (
	"fmt"

	Geo "github.com/hariank/trace/pkg/geometry"
)

type Color = Geo.Vec

func PPMStr(color Color) string {
	ir := int(255.999 * color.X)
	ig := int(255.999 * color.Y)
	ib := int(255.999 * color.Z)

	return fmt.Sprintf("%d %d %d\n", ir, ig, ib)
}

// assuming n is unit length, map components from [-1,1] to [0,1]
func VisualizeNormal(n Geo.Vec) Color {
	return n.Plus(Geo.Vec{1, 1, 1}).Scale(0.5)
}
