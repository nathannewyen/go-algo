package greedy

// CanReachEnd determines if you can reach the last index from the first
// Each element represents maximum jump length from that position
func CanReachEnd(jumpLengths []int) bool {
	farthestReachable := 0

	for currentPosition := 0; currentPosition < len(jumpLengths); currentPosition++ {
		// If current position is beyond farthest reachable, cannot proceed
		if currentPosition > farthestReachable {
			return false
		}
		// Update farthest reachable position
		potentialReach := currentPosition + jumpLengths[currentPosition]
		if potentialReach > farthestReachable {
			farthestReachable = potentialReach
		}
	}
	return true
}

// MinJumpsToEnd returns minimum number of jumps to reach the last index
func MinJumpsToEnd(jumpLengths []int) int {
	jumpCount := 0
	currentRangeEnd := 0
	farthestReachable := 0

	for currentPosition := 0; currentPosition < len(jumpLengths)-1; currentPosition++ {
		potentialReach := currentPosition + jumpLengths[currentPosition]
		if potentialReach > farthestReachable {
			farthestReachable = potentialReach
		}
		// When we reach end of current jump range, must take another jump
		if currentPosition == currentRangeEnd {
			jumpCount++
			currentRangeEnd = farthestReachable
		}
	}
	return jumpCount
}
