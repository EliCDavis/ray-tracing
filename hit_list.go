package main

type HitList []Hittable

func (h *HitList) Hit(r *Ray, min float32, max float32, hitRecord *HitRecord) bool {
	var tempRecord *HitRecord = NewHitRecord()
	hitAnything := false
	closestSoFar := max

	for i := 0; i < len(*h); i++ {
		if (*h)[i].Hit(r, min, closestSoFar, tempRecord) {
			hitAnything = true
			closestSoFar = tempRecord.t

			hitRecord.t = tempRecord.t
			hitRecord.p = tempRecord.p
			hitRecord.normal = tempRecord.normal
		}
	}

	return hitAnything
}
