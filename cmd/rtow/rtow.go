package main

import (
	"flag"
	"fmt"
	"path"

	"github.com/hariank/trace/pkg/camera"
	geo "github.com/hariank/trace/pkg/geometry"
	"github.com/hariank/trace/pkg/render"
)

const (
	ASPECT_RATIO float64 = 16. / 9.
)

var (
	ImageWidth      = flag.Int("width", 400, "")
	SamplesPerPixel = flag.Int("spp", 100, "")
	MaxRayBounces   = flag.Int("bou", 50, "")
	ImageName       = flag.String("name", "image.ppm", "")
)

func main() {
	flag.Parse()
	ImageHeight := int(float64(*ImageWidth) / ASPECT_RATIO)
	fmt.Println(*ImageWidth, ImageHeight, *SamplesPerPixel, *MaxRayBounces)

	groundMaterial := &geo.Lambertian{Albedo: geo.Color{0.8, 0.8, 0.0}}
	sphereMaterial := &geo.Hemispherical{Albedo: geo.Color{0.7, 0.3, 0.3}}
	objects := make([]geo.Hittable, 2)
	objects[0] = geo.NewSphere(geo.Point{0, 0, -1}, 0.5, sphereMaterial)
	objects[1] = geo.NewSphere(geo.Point{0, -100.5, -1}, 100, groundMaterial)
	world := geo.HittableList{objects}

	renderConfig := &render.RenderConfig{*ImageWidth, ImageHeight, *SamplesPerPixel, *MaxRayBounces}
	camera := camera.NewPinholeCamera(ASPECT_RATIO)
	render.Render(world, camera, path.Join("../../images", *ImageName), renderConfig)
}
