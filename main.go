package main

import (
	"bufio"
	"fmt"
	"math"
	"math/rand"
	"os"
)

var MAX_FLOAT float32 = 10000.0

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func color(r *Ray, world Hittable) *Vector3 {
	hitRecord := NewHitRecord()
	if world.Hit(r, 0.001, MAX_FLOAT, hitRecord) {
		target := hitRecord.p.Add(hitRecord.normal).Add(randomInUnitSphere())
		return color(NewRay(hitRecord.p, target.Sub(hitRecord.p)), world).MultByConstant(.5)
	} else {
		t := .5 * (r.Direction().Normalized().Y() + 1.0)
		return Vector3One.MultByConstant(1.0 - t).Add(NewVector3(.5, .7, 1.0).MultByConstant(t))
	}
}

func randomInUnitSphere() *Vector3 {
	var p Vector3
	for {
		p = *NewVector3(rand.Float32(), rand.Float32(), rand.Float32()).MultByConstant(2).Sub(Vector3One)
		if p.Length() < 1.0 {
			break
		}
	}
	return &p
}

func main() {

	f, err := os.Create("out.ppm")
	check(err)

	defer f.Close()

	w := bufio.NewWriter(f)
	defer w.Flush()

	nx := 1000
	ny := 500
	ns := 100

	w.WriteString(fmt.Sprintf("P3\n%d %d\n255\n", nx, ny))

	lowerLeftCorner := NewVector3(-2.0, -1.0, -1.0)
	horizontal := NewVector3(4.0, 0, 0)
	vertical := NewVector3(0, 2.0, 0)
	origin := NewVector3(0, 0, 0)

	camera := NewCamera(lowerLeftCorner, horizontal, vertical, origin)

	var world HitList = make([]Hittable, 2)
	world[0] = NewSphere(NewVector3(0, 0, -1), 0.5)
	world[1] = NewSphere(NewVector3(0, -100.5, -1), 100)

	for j := ny - 1; j >= 0; j-- {
		for i := 0; i < nx; i++ {

			col := Vector3Zero

			for s := 0; s < ns; s++ {
				u := (float32(i) + rand.Float32()) / float32(nx)
				v := (float32(j) + rand.Float32()) / float32(ny)
				col = col.Add(color(camera.GetRay(u, v), &world))
			}

			col = col.DivByConstant(float32(ns))

			col = NewVector3(
				float32(math.Sqrt(float64(col.X()))),
				float32(math.Sqrt(float64(col.Y()))),
				float32(math.Sqrt(float64(col.Z()))),
			)

			_, err := w.WriteString(fmt.Sprintf("%d %d %d\n", int(col.R()*255.99), int(col.G()*255.99), int(col.B()*255.99)))
			check(err)
		}
	}

}
