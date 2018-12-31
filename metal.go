package main

type Metal struct {
	color *Vector3
}

func NewMetal(color *Vector3) *Metal {
	return &Metal{color}
}

func (l Metal) reflect(v *Vector3, n *Vector3) *Vector3 {
	return v.Sub(n.MultByConstant(2 * v.Dot(n)))
}

func (l Metal) Scatter(in *Ray, rec *HitRecord, attenuation *Vector3, scattered *Ray) bool {
	reflected := l.reflect(in.Direction().Normalized(), rec.Normal())
	scattered = NewRay(rec.p, reflected)
	attenuation = l.color
	return scattered.Direction().Dot(rec.Normal()) > 0
}
