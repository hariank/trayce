package main

import (
	"bufio"
	"fmt"
	"math"
	"os"

	"github.com/cheggaaa/pb"
	"github.com/hariank/trace/pkg/camera"
	"github.com/hariank/trace/pkg/color"
	geo "github.com/hariank/trace/pkg/geometry"
)

const (
	ASPECT_RATIO float64 = 16. / 9.
	IMAGE_WIDTH  int     = 400
	IMAGE_HEIGHT int     = int(float64(IMAGE_WIDTH) / ASPECT_RATIO)
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func ray_color(ray *geo.Ray, world geo.HittableList) color.Color {
	if hit, hitRecord := world.Hit(ray, 0, math.Inf(1)); hit {
		return color.VisualizeNormal(hitRecord.Norm)
	}
	unitDir := ray.Dir.Unit()
	t := (unitDir.Y + 1.) * .5 // scale (-1,1) to (0,1)
	return color.Color{1., 1., 1.}.Scale(1. - t).Plus(color.Color{.5, .7, 1.}.Scale(t))
}

func main() {
	f, err := os.Create("../../images/a.ppm")
	check(err)
	defer f.Close()
	w := bufio.NewWriter(f)

	// world := geo.NewHittableList()
	// world.AddHittables(geo.NewSphere(geo.Point{0, 0, -1}, 0.5), geo.NewSphere(geo.Point{0, -100.5, -1}, 100))
	// world.AddHittable(geo.NewSphere(geo.Point{0, 0, -1}, 0.5))
	// world.AddHittable(geo.NewSphere(geo.Point{0, -100.5, -1}, 100))
	// world.AddHittable(geo.NewSphere(geo.Point{0, -100.5, -1}, 100))
	objects := make([]geo.Hittable, 2)
	objects[0] = geo.NewSphere(geo.Point{0, 0, -1}, 0.5)
	objects[1] = geo.NewSphere(geo.Point{0, -100.5, -1}, 100)
	world := geo.HittableList{objects}

	camera := camera.NewPinholeCamera(ASPECT_RATIO)

	bar := pb.StartNew(IMAGE_WIDTH * IMAGE_HEIGHT)
	_, err = fmt.Fprintf(w, "P3\n%d %d\n255\n", IMAGE_WIDTH, IMAGE_HEIGHT)
	check(err)
	for i := 0; i < IMAGE_HEIGHT; i++ {
		for j := 0; j < IMAGE_WIDTH; j++ {
			r := float64(j) / (float64(IMAGE_WIDTH) - 1)
			g := float64(IMAGE_HEIGHT-i) / (float64(IMAGE_HEIGHT) - 1)
			ray := camera.GenerateRay(r, g)
			c := ray_color(ray, world)
			_, err = fmt.Fprintf(w, color.PPMStr(c))
			check(err)
			bar.Increment()
		}
	}
	w.Flush()
	bar.Finish()
}