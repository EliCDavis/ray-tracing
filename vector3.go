package main

import "math"

type Vector3 struct {
	x float32
	y float32
	z float32
}

func NewVector3(x float32, y float32, z float32) *Vector3 {
	return &Vector3{
		x: x,
		y: y,
		z: z,
	}
}

var Right = NewVector3(1, 0, 0)
var Vector3One = NewVector3(1, 1, 1)
var Vector3Zero = NewVector3(0, 0, 0)

func (v *Vector3) X() float32 {
	return v.x
}

func (v *Vector3) Y() float32 {
	return v.y
}

func (v *Vector3) Z() float32 {
	return v.z
}

func (v *Vector3) R() float32 {
	return v.x
}

func (v *Vector3) G() float32 {
	return v.y
}

func (v *Vector3) B() float32 {
	return v.z
}

func (v *Vector3) Add(other *Vector3) *Vector3 {
	return &Vector3{
		x: v.x + other.x,
		y: v.y + other.y,
		z: v.z + other.z,
	}
}

func (v *Vector3) Sub(other *Vector3) *Vector3 {
	return &Vector3{
		x: v.x - other.x,
		y: v.y - other.y,
		z: v.z - other.z,
	}
}

func (v *Vector3) Dot(other *Vector3) float32 {
	return (v.x * other.x) + (v.y * other.y) + (v.z * other.z)
}

func (v *Vector3) Cross(other *Vector3) *Vector3 {
	return &Vector3{
		x: (v.y * other.z) - (v.z * other.y),
		y: -((v.x * other.z) - (v.z * other.x)),
		z: (v.x * other.y) - (v.y * other.x),
	}
}

func (v *Vector3) Normalized() *Vector3 {
	return v.DivByConstant(v.Length())
}

func (v *Vector3) MultByConstant(t float32) *Vector3 {
	return &Vector3{
		x: v.x * t,
		y: v.y * t,
		z: v.z * t,
	}
}

func (v *Vector3) DivByConstant(t float32) *Vector3 {
	return v.MultByConstant(1.0 / t)
}

func (v *Vector3) Length() float32 {
	return float32(math.Sqrt(float64((v.x * v.x) + (v.y * v.y) + (v.z * v.z))))
}

func (v *Vector3) SquaredLength() float32 {
	return (v.x * v.x) + (v.y * v.y) + (v.z * v.z)
}
