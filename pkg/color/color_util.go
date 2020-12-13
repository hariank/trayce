package color

import (
	"fmt"
	"math"

	internal "github.com/hariank/trace/internal"
	geo "github.com/hariank/trace/pkg/geometry"
)

type Color = geo.Vec

func PPMStr(color Color) string {
	ir := int(255.999 * internal.Clamp(color.X, 0., 1.))
	ig := int(255.999 * internal.Clamp(color.Y, 0., 1.))
	ib := int(255.999 * internal.Clamp(color.Z, 0., 1.))

	return fmt.Sprintf("%d %d %d\n", ir, ig, ib)
}

func SampledPPMStr(accumulatedColor Color, samplesPerPixel int) string {
	color := accumulatedColor.Scale(1. / float64(samplesPerPixel))
	color = gamma2Correct(color)
	return PPMStr(color)
}

func gamma2Correct(color Color) Color {
	return Color{math.Sqrt(color.X), math.Sqrt(color.Y), math.Sqrt(color.Z)}
}

// assuming n is unit length, map components from [-1,1] to [0,1]
func VisualizeNormal(n geo.Vec) Color {
	return n.Plus(geo.Vec{1, 1, 1}).Scale(0.5)
}
