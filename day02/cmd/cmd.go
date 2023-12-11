package main

import (
	"advent-of-code/day02/game"
	"bufio"
	"fmt"
	"os"
)

func main() {
	fileName := "./day02/input.txt"
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
		gamePlayed := game.NewGame(text)

		minCubes := gamePlayed.MiniumRequired()
		powers := 1
		for _, v := range minCubes {
			if v > 0 {
				powers *= v
			}
		}
		sum += powers
	}
	fmt.Printf("sum: %d", sum)
}
