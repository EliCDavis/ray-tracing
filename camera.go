package main

type Camera struct {
	lowerLeftCorner *Vector3
	horizontal      *Vector3
	vertical        *Vector3
	origin          *Vector3
}

func NewCamera(lowerLeftCorner *Vector3, horizontal *Vector3, vertical *Vector3, origin *Vector3) *Camera {
	return &Camera{lowerLeftCorner, horizontal, vertical, origin}
}

func (c Camera) GetRay(u float32, v float32) *Ray {
	return NewRay(c.origin, c.lowerLeftCorner.Add(c.horizontal.MultByConstant(u)).Add(c.vertical.MultByConstant(v)))
}
