package main

type Material interface {
	Scatter(in *Ray, rec *HitRecord, attenuation *Vector3, scattered *Ray) bool
}
