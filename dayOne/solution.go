package dayOne

import (
	"sort"
)

func CalculateTotalDistance(left, right []int) int {
	// Sort both lists
	sort.Ints(left)
	sort.Ints(right)

	// Calculate total distance
	totalDistance := 0
	for i := 0; i < len(left); i++ {
		totalDistance += abs(left[i] - right[i])
	}
	return totalDistance
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
