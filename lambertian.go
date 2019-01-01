package main

type Lambertian struct {
	color *Vector3
}

func NewLambertian(color *Vector3) *Lambertian {
	return &Lambertian{color}
}

func (l Lambertian) Scatter(in *Ray, rec *HitRecord, attenuation *Vector3, scattered *Ray) bool {
	target := rec.normal.Add(rec.p).Add(randomInUnitSphere())
	*scattered = *NewRay(rec.p, target.Sub(rec.p))
	*attenuation = *l.color
	return true
}
