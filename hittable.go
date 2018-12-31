package main

type Hittable interface {
	Hit(r *Ray, min float32, max float32, hitRecord *HitRecord) bool
}
