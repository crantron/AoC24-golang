package dayOne

func CalculateSimilarityScore(left, right []int) int {
	// Build a frequency map for the right list
	freqMap := make(map[int]int)
	for _, num := range right {
		freqMap[num]++
	}

	// Calculate similarity score
	similarityScore := 0
	for _, num := range left {
		similarityScore += num * freqMap[num]
	}
	return similarityScore
}
