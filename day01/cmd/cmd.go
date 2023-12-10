package main

import (
	"advent-of-code/day01/puzzle"
	"bufio"
	"fmt"
	"os"
)

func main() {
	fileName := "./puzzle_input.txt"
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
		i, err := puzzle.FirstAndLastDigitAsInt(text)
		if err != nil {
			fmt.Printf("could not read ints from line %s, %v\n", text, err)
			os.Exit(1)
		}
		sum += i
	}
	fmt.Printf("sum: %d", sum)
}
