package main

import (
	"bufio"
	"fmt"
	"math"
	"os"

	"github.com/cheggaaa/pb"
	internal "github.com/hariank/trace/internal"
	"github.com/hariank/trace/pkg/camera"
	"github.com/hariank/trace/pkg/color"
	geo "github.com/hariank/trace/pkg/geometry"
)

const (
	ASPECT_RATIO float64 = 16. / 9.
	IMAGE_WIDTH  int     = 400
	IMAGE_HEIGHT int     = int(float64(IMAGE_WIDTH) / ASPECT_RATIO)

	SAMPLES_PER_PIXEL int = 100
	MAX_RAY_BOUNCES   int = 50
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func ray_color(ray *geo.Ray, world geo.HittableList, depth int) color.Color {
	if depth == 0 {
		return color.Color{0, 0, 0}
	}
	if hit, hitRecord := world.Hit(ray, .001, math.Inf(1)); hit {
		diffuseVec := color.LambertianDiffuse(hitRecord)
		// diffuseVec := color.HemisphericalDiffuse(hitRecord)
		c := ray_color(geo.NewRay(hitRecord.Loc, diffuseVec), world, depth-1).Scale(.5)
		// return color.VisualizeNormal(hitRecord.Norm)
	}
	unitDir := ray.Dir.Unit()
	t := (unitDir.Y + 1.) * .5 // scale (-1,1) to (0,1)
	return color.Color{1., 1., 1.}.Scale(1. - t).Plus(color.Color{.5, .7, 1.}.Scale(t))
}

func main() {
	f, err := os.Create("../../images/image.ppm")
	check(err)
	defer f.Close()
	w := bufio.NewWriter(f)

	objects := make([]geo.Hittable, 2)
	objects[0] = geo.NewSphere(geo.Point{0, 0, -1}, 0.5)
	objects[1] = geo.NewSphere(geo.Point{0, -100.5, -1}, 100)
	world := geo.HittableList{objects}

	camera := camera.NewPinholeCamera(ASPECT_RATIO)

	progress := pb.StartNew(IMAGE_WIDTH * IMAGE_HEIGHT)
	_, err = fmt.Fprintf(w, "P3\n%d %d\n255\n", IMAGE_WIDTH, IMAGE_HEIGHT)
	check(err)
	for i := 0; i < IMAGE_HEIGHT; i++ {
		for j := 0; j < IMAGE_WIDTH; j++ {
			accumulatedColor := color.Color{0, 0, 0}
			for s := 0; s < SAMPLES_PER_PIXEL; s++ {
				u := (float64(j) + internal.RandomFloat()) / (float64(IMAGE_WIDTH) - 1)
				v := (float64(IMAGE_HEIGHT-i) + internal.RandomFloat()) / (float64(IMAGE_HEIGHT) - 1)
				ray := camera.GenerateRay(u, v)
				accumulatedColor = accumulatedColor.Plus(ray_color(ray, world, MAX_RAY_BOUNCES))
			}
			_, err = fmt.Fprintf(w, color.SampledPPMStr(accumulatedColor, SAMPLES_PER_PIXEL))
			check(err)
			progress.Increment()
		}
	}
	w.Flush()
	progress.Finish()
}
