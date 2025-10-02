package greedy

import "sort"

// KnapsackItem represents an item with weight and value
type KnapsackItem struct {
	Weight int
	Value  int
}

// FractionalKnapsack maximizes value by taking fractions of items
func FractionalKnapsack(itemList []KnapsackItem, maxCapacity int) float64 {
	// Sort items by value-to-weight ratio in descending order
	sort.Slice(itemList, func(i int, j int) bool {
		ratioFirst := float64(itemList[i].Value) / float64(itemList[i].Weight)
		ratioSecond := float64(itemList[j].Value) / float64(itemList[j].Weight)
		return ratioFirst > ratioSecond
	})

	totalValue := 0.0
	remainingCapacity := maxCapacity

	for _, currentItem := range itemList {
		if remainingCapacity == 0 {
			break
		}
		if currentItem.Weight <= remainingCapacity {
			// Take the whole item
			totalValue += float64(currentItem.Value)
			remainingCapacity -= currentItem.Weight
		} else {
			// Take a fraction of the item
			fraction := float64(remainingCapacity) / float64(currentItem.Weight)
			totalValue += float64(currentItem.Value) * fraction
			remainingCapacity = 0
		}
	}
	return totalValue
}
