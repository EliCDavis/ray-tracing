package main

type Dielectric struct {
	reflectiveIndex float32
}

func NewDielectric(reflectiveIndex float32) *Dielectric {
	return &Dielectric{reflectiveIndex}
}

func (d Dielectric) Scatter(in *Ray, rec *HitRecord, attenuation *Vector3, scattered *Ray) bool {
	*attenuation = *NewVector3(1.0, 1.0, 0)

	var outwardNormal *Vector3
	var niOverNt float32

	if in.Direction().Dot(rec.Normal()) > 0 {
		outwardNormal = rec.Normal().MultByConstant(-1)
		niOverNt = d.reflectiveIndex
	} else {
		outwardNormal = rec.Normal()
		niOverNt = 1.0 / d.reflectiveIndex
	}

	refracted := Vector3Zero()

	if refract(in.Direction(), outwardNormal, niOverNt, refracted) {
		*scattered = *NewRay(rec.p, refracted)
		return true
	} else {
		*scattered = *NewRay(rec.p, reflect(in.Direction(), rec.Normal()))
		return false
	}

}
