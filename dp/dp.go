package dp

// Fibonacci returns the nth Fibonacci number using bottom-up dynamic programming.
// Uses O(n) time and O(1) space by only tracking the two previous values.
func Fibonacci(nthPosition int) int {
	if nthPosition <= 0 {
		return 0
	}
	if nthPosition == 1 {
		return 1
	}

	previousValue := 0
	currentValue := 1

	for step := 2; step <= nthPosition; step++ {
		nextValue := previousValue + currentValue
		previousValue = currentValue
		currentValue = nextValue
	}

	return currentValue
}

// ClimbStairs returns the number of distinct ways to climb n stairs,
// where each step can be either 1 or 2 stairs at a time.
// This is equivalent to Fibonacci because f(n) = f(n-1) + f(n-2).
func ClimbStairs(totalStairs int) int {
	if totalStairs <= 0 {
		return 0
	}
	if totalStairs == 1 {
		return 1
	}

	// waysToReachPrevious represents ways to reach (step - 1)
	waysToReachPrevious := 1
	// waysToReachCurrent represents ways to reach the current step
	waysToReachCurrent := 2

	for step := 3; step <= totalStairs; step++ {
		waysToReachNext := waysToReachPrevious + waysToReachCurrent
		waysToReachPrevious = waysToReachCurrent
		waysToReachCurrent = waysToReachNext
	}

	return waysToReachCurrent
}

// LongestCommonSubsequence finds the length of the longest common subsequence
// between two strings using a 2D dynamic programming table.
func LongestCommonSubsequence(firstString string, secondString string) int {
	firstLength := len(firstString)
	secondLength := len(secondString)

	// dpTable[row][col] stores the LCS length for firstString[:row] and secondString[:col]
	dpTable := make([][]int, firstLength+1)
	for row := 0; row <= firstLength; row++ {
		dpTable[row] = make([]int, secondLength+1)
	}

	for firstIndex := 1; firstIndex <= firstLength; firstIndex++ {
		for secondIndex := 1; secondIndex <= secondLength; secondIndex++ {
			// If characters match, extend the LCS from the diagonal
			if firstString[firstIndex-1] == secondString[secondIndex-1] {
				dpTable[firstIndex][secondIndex] = dpTable[firstIndex-1][secondIndex-1] + 1
			} else {
				// Take the longer LCS from either dropping a char from firstString or secondString
				if dpTable[firstIndex-1][secondIndex] > dpTable[firstIndex][secondIndex-1] {
					dpTable[firstIndex][secondIndex] = dpTable[firstIndex-1][secondIndex]
				} else {
					dpTable[firstIndex][secondIndex] = dpTable[firstIndex][secondIndex-1]
				}
			}
		}
	}

	return dpTable[firstLength][secondLength]
}

// CoinChange returns the minimum number of coins needed to make the target amount.
// Returns -1 if the amount cannot be made with the given coin denominations.
func CoinChange(coinDenominations []int, targetAmount int) int {
	if targetAmount == 0 {
		return 0
	}

	// minCoinsForAmount[amount] stores the fewest coins needed to make that amount.
	// Initialize with targetAmount+1 as a sentinel value representing "impossible".
	impossibleSentinel := targetAmount + 1
	minCoinsForAmount := make([]int, targetAmount+1)
	for amount := 1; amount <= targetAmount; amount++ {
		minCoinsForAmount[amount] = impossibleSentinel
	}

	for amount := 1; amount <= targetAmount; amount++ {
		for _, coinValue := range coinDenominations {
			if coinValue <= amount {
				// Check if using this coin yields fewer total coins
				coinsIfUsingThisCoin := minCoinsForAmount[amount-coinValue] + 1
				if coinsIfUsingThisCoin < minCoinsForAmount[amount] {
					minCoinsForAmount[amount] = coinsIfUsingThisCoin
				}
			}
		}
	}

	if minCoinsForAmount[targetAmount] == impossibleSentinel {
		return -1
	}

	return minCoinsForAmount[targetAmount]
}

// MaxSubarraySum finds the maximum sum of any contiguous subarray
// using Kadane's algorithm. Returns 0 for empty input.
func MaxSubarraySum(numbers []int) int {
	if len(numbers) == 0 {
		return 0
	}

	maxEndingHere := numbers[0]
	globalMaxSum := numbers[0]

	for index := 1; index < len(numbers); index++ {
		// Either extend the current subarray or start a new one from this element
		if maxEndingHere+numbers[index] > numbers[index] {
			maxEndingHere = maxEndingHere + numbers[index]
		} else {
			maxEndingHere = numbers[index]
		}

		if maxEndingHere > globalMaxSum {
			globalMaxSum = maxEndingHere
		}
	}

	return globalMaxSum
}
