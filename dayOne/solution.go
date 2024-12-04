package dayOne

import (
	"bufio"
	"os"
	"sort"
	"strconv"
	"strings"
)

func ParseInputFromFile(filename string) ([]int, []int, error) {
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
