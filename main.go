package main

import (
	"AoC24/dayOne"
	"fmt"
	"log"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <input_file>")
		return
	}

	filename := os.Args[1]
	left, right, err := dayOne.ParseInputFromFile(filename)
	if err != nil {
		log.Fatalf("Error reading file: %v\n", err)
	}
	//Day one Part 1
	totalDistance := dayOne.CalculateTotalDistance(left, right)
	fmt.Printf("Total Distance: %d\n", totalDistance)

	//Day one Part 2
	similarityScore := dayOne.CalculateSimilarityScore(left, right)
	fmt.Printf("Similarity Score: %d\n", similarityScore)

}
