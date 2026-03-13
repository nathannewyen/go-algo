package hashing

// TopKFrequent returns the k most frequent elements in the array
func TopKFrequent(numbers []int, frequencyRank int) []int {
	// Count frequency of each number
	frequencyCount := map[int]int{}
	for _, num := range numbers {
		frequencyCount[num]++
	}
	// Use bucket sort where index represents frequency
	maxFrequency := len(numbers)
	frequencyBuckets := make([][]int, maxFrequency+1)
	for number, count := range frequencyCount {
		frequencyBuckets[count] = append(frequencyBuckets[count], number)
	}
	// Collect results starting from highest frequency bucket
	topElements := []int{}
	for bucketIndex := maxFrequency; bucketIndex >= 0 && len(topElements) < frequencyRank; bucketIndex-- {
		for _, number := range frequencyBuckets[bucketIndex] {
			topElements = append(topElements, number)
			if len(topElements) == frequencyRank {
				break
			}
		}
	}
	return topElements
}
