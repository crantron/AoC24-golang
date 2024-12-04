package dayOne

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func ParseInputFromFileTwo(filename string) ([]int, []int, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, nil, err
	}
	defer file.Close()

	left := []int{}
	right := []int{}
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Fields(line)
		if len(parts) == 2 {
			leftNum, _ := strconv.Atoi(parts[0])
			rightNum, _ := strconv.Atoi(parts[1])
			left = append(left, leftNum)
			right = append(right, rightNum)
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, nil, err
	}
	return left, right, nil
}

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
