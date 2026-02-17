func champagneTower(poured, I, J int) float64 {
	tower := [101][101]float64{}
	tower[0][0] = float64(poured)
	for i := 0; i <= I; i++ {
		for j := 0; j < i + 1; j++ {
			if tower[i][j] >= 1 {
				pass := tower[i][j] - 1
				tower[i][j] = 1
				tower[i + 1][j] += pass / 2
				tower[i + 1][j + 1] += pass / 2
			}
		}
	}
	return tower[I][J]
}
