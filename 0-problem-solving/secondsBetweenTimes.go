func secondsBetweenTimes(startTime string, endTime string) int {
	arrX := strings.Split(startTime, ":")
	arrY := strings.Split(endTime, ":")
	x, y := 0, 0
	for i := 0; i < len(arrX); i++ {
		tX, _ := strconv.Atoi(arrX[i])
		tY, _ := strconv.Atoi(arrY[i])
		switch i {
		case 0:
			x += tX * 3600
			y += tY * 3600
		case 1:
			x += tX * 60
			y += tY * 60
		case 2:
			x += tX
			y += tY
		}
	}
	return y - x
}
