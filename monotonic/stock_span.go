package monotonic

// StockSpan calculates the span of stock prices for each day
// Span is the number of consecutive days with price <= current day price
func StockSpan(dailyPrices []int) []int {
	dayCount := len(dailyPrices)
	spanValues := make([]int, dayCount)
	// Stack stores indices of days with prices in decreasing order
	priceStack := []int{}

	for currentDay := 0; currentDay < dayCount; currentDay++ {
		// Pop all days with prices less than or equal to current price
		for len(priceStack) > 0 && dailyPrices[priceStack[len(priceStack)-1]] <= dailyPrices[currentDay] {
			priceStack = priceStack[:len(priceStack)-1]
		}
		if len(priceStack) == 0 {
			spanValues[currentDay] = currentDay + 1
		} else {
			spanValues[currentDay] = currentDay - priceStack[len(priceStack)-1]
		}
		priceStack = append(priceStack, currentDay)
	}
	return spanValues
}
