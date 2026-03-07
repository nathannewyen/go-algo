package stack

// NextGreaterElement finds the next greater element for each element in the array
func NextGreaterElement(numbers []int) []int {
	elementCount := len(numbers)
	nextGreater := make([]int, elementCount)
	// Initialize all positions with -1 meaning no greater element found
	for i := range nextGreater {
		nextGreater[i] = -1
	}
	// Stack holds indices of elements waiting for their next greater element
	waitingIndices := []int{}
	for currentIndex := 0; currentIndex < elementCount; currentIndex++ {
		for len(waitingIndices) > 0 {
			topWaitingIndex := waitingIndices[len(waitingIndices)-1]
			if numbers[currentIndex] <= numbers[topWaitingIndex] {
				break
			}
			waitingIndices = waitingIndices[:len(waitingIndices)-1]
			nextGreater[topWaitingIndex] = numbers[currentIndex]
		}
		waitingIndices = append(waitingIndices, currentIndex)
	}
	return nextGreater
}
