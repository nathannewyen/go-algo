package intervals

// MinArrowsToBurstBalloons finds minimum arrows needed to burst all balloons
// Each balloon is represented as an interval [start, end] on x-axis
func MinArrowsToBurstBalloons(balloons [][]int) int {
	if len(balloons) == 0 {
		return 0
	}
	// Sort balloons by their end position
	for i := 1; i < len(balloons); i++ {
		key := balloons[i]
		j := i - 1
		for j >= 0 && balloons[j][1] > key[1] {
			balloons[j+1] = balloons[j]
			j--
		}
		balloons[j+1] = key
	}
	arrowCount := 1
	currentArrowPosition := balloons[0][1]
	// Each arrow bursts all balloons that overlap with its position
	for i := 1; i < len(balloons); i++ {
		balloonStart := balloons[i][0]
		if balloonStart > currentArrowPosition {
			arrowCount++
			currentArrowPosition = balloons[i][1]
		}
	}
	return arrowCount
}
