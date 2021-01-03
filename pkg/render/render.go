package render

import (
	"bufio"
	"fmt"
	"math"
	"os"

	"github.com/cheggaaa/pb"
	util "github.com/hariank/trace/internal"
	"github.com/hariank/trace/pkg/camera"
	geo "github.com/hariank/trace/pkg/geometry"
)

type Color = geo.Color

type RenderConfig struct {
	ImageWidth, ImageHeight        int
	SamplesPerPixel, MaxRayBounces int
}

func Render(world geo.HittableList, camera camera.Camera, filename string, config *RenderConfig) {
	f, err := os.Create(filename)
	check(err)
	defer f.Close()
	w := bufio.NewWriter(f)
	_, err = fmt.Fprintf(w, "P3\n%d %d\n255\n", config.ImageWidth, config.ImageHeight)
	check(err)

	progress := pb.StartNew(config.ImageWidth * config.ImageHeight)
	for i := 0; i < config.ImageHeight; i++ {
		for j := 0; j < config.ImageWidth; j++ {
			accumulatedColor := Color{0, 0, 0}
			for s := 0; s < config.SamplesPerPixel; s++ {
				u := (float64(j) + util.RandomFloat()) / (float64(config.ImageWidth) - 1)
				v := (float64(config.ImageHeight-i) + util.RandomFloat()) / (float64(config.ImageHeight) - 1)
				ray := camera.GenerateRay(u, v)
				accumulatedColor = accumulatedColor.Plus(rayColor(ray, world, config.MaxRayBounces))
			}
			_, err = fmt.Fprintf(w, sampledPPMStr(accumulatedColor, config.SamplesPerPixel))
			check(err)
			progress.Increment()
		}
	}
	w.Flush()
	progress.Finish()
}

func rayColor(ray *geo.Ray, world geo.HittableList, depth int) Color {
	if depth == 0 {
		return Color{0, 0, 0}
	}
	if hit, hitRecord := world.Hit(ray, .001, math.Inf(1)); hit {
		attenuationColor, scatteredRay := hitRecord.Material.ScatterRay(ray, hitRecord)
		return attenuationColor.Times(rayColor(scatteredRay, world, depth-1))
	}
	unitDir := ray.Dir.Unit()
	t := (unitDir.Y + 1.) * .5 // scale (-1,1) to (0,1)
	return Color{1., 1., 1.}.Scale(1. - t).Plus(Color{.5, .7, 1.}.Scale(t))
}

func PPMStr(color Color) string {
	ir := int(255.999 * util.Clamp(color.X, 0., 1.))
	ig := int(255.999 * util.Clamp(color.Y, 0., 1.))
	ib := int(255.999 * util.Clamp(color.Z, 0., 1.))

	return fmt.Sprintf("%d %d %d\n", ir, ig, ib)
}

func sampledPPMStr(accumulatedColor Color, samplesPerPixel int) string {
	color := accumulatedColor.Scale(1. / float64(samplesPerPixel))
	color = geo.Gamma2Correct(color)
	return PPMStr(color)
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
