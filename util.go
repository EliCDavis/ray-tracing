package main

import "math"

func reflect(v *Vector3, n *Vector3) *Vector3 {
	return v.Sub(n.MultByConstant(2 * v.Dot(n)))
}

func refract(v *Vector3, n *Vector3, niOverNt float32, refracted *Vector3) bool {
	uv := v.Normalized()
	dt := uv.Dot(n)
	discriminant := 1.0 - (niOverNt*niOverNt)*(1-(dt*dt))
	if discriminant > 0 {
		*refracted = *uv.Sub(n.MultByConstant(dt)).MultByConstant(niOverNt).Sub(n.MultByConstant(float32(math.Sqrt(float64(discriminant)))))
		return true
	}
	return false
}
