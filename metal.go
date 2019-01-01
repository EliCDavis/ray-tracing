package main

type Metal struct {
	color *Vector3
	fuzz  float32
}

func NewMetal(color *Vector3) *Metal {
	return &Metal{color, 0}
}

func NewFuzzyMetal(color *Vector3, fuzz float32) *Metal {
	return &Metal{color, fuzz}
}

func (l Metal) Scatter(in *Ray, rec *HitRecord, attenuation *Vector3, scattered *Ray) bool {
	reflected := reflect(in.Direction().Normalized(), rec.Normal())
	*scattered = *NewRay(rec.p, reflected.Add(randomInUnitSphere().MultByConstant(l.fuzz)))
	*attenuation = *l.color
	return scattered.Direction().Dot(rec.Normal()) > 0
}
