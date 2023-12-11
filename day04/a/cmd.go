package main

import (
	"advent-of-code/day04/puzzle"
	"bufio"
	"fmt"
	"math"
	"os"
)

func main() {
	fileName := "./day04/input.txt"
	sum := 0
	file, err := os.Open(fileName)
	defer func() { _ = file.Close() }()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := scanner.Text()
		winners, myNumbers := puzzle.GetWinnersAndMyNumbers(text)
		count := puzzle.MatchCount(winners, myNumbers)
		score := int(math.Pow(2, float64(count-1)))
		sum += score
		fmt.Println(score)
	}
	fmt.Printf("sum: %d", sum)
}
