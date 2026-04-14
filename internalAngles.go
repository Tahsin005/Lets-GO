func internalAngles(sides []int) []float64 {
	a, b, c := sides[0], sides[1], sides[2]
	if a + b <= c {
		return []float64{}
	}

	if b + c <= a {
		return []float64{}
	}

	if a + c <= b {
		return []float64{}
	}

	aQ, bQ, cQ := a * a, b * b, c * c
	bc2 := 2 * b * c
	ca2 := 2 * c * a
	ab2 := 2 * a * b
	aR := math.Acos(float64((bQ + cQ - aQ)) / float64(bc2)) * 180 / math.Pi
	bR := math.Acos(float64((cQ + aQ - bQ)) / float64(ca2)) * 180 / math.Pi
	cR := math.Acos(float64((aQ + bQ - cQ)) / float64(ab2)) * 180 / math.Pi
	r := []float64{aR, bR, cR}
	sort.Float64s(r)

	return r
}
