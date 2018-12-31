package main

type HitRecord struct {
	t      float32
	p      *Vector3
	normal *Vector3
}

func NewHitRecord() *HitRecord {
	return &HitRecord{
		0,
		Vector3Zero,
		Vector3Zero,
	}
}

func (h *HitRecord) Normal() *Vector3 {
	return h.normal
}
