package main

import "math"

type Sphere struct {
	center *Vector3
	radius float32
}

func NewSphere(center *Vector3, radius float32) *Sphere {
	return &Sphere{center, radius}
}

func (s *Sphere) Hit(r *Ray, min float32, max float32, hitRecord *HitRecord) bool {
	oc := r.Origin().Sub(s.center)
	a := r.Direction().Dot(r.Direction())
	b := oc.Dot(r.Direction())
	c := oc.Dot(oc) - (s.radius * s.radius)
	discriminant := float64((b * b) - (a * c))
	if discriminant > 0 {
		temp := (-b - float32(math.Sqrt(discriminant))) / a
		if temp < max && temp > min {
			hitRecord.t = temp
			hitRecord.p = r.PointAt(hitRecord.t)
			hitRecord.normal = hitRecord.p.Sub(s.center).DivByConstant(s.radius)
			return true
		}

		temp = (-b + float32(math.Sqrt(discriminant))) / a
		if temp < max && temp > min {
			hitRecord.t = temp
			hitRecord.p = r.PointAt(hitRecord.t)
			hitRecord.normal = hitRecord.p.Sub(s.center).DivByConstant(s.radius)
			return true
		}
	}

	return false
}
