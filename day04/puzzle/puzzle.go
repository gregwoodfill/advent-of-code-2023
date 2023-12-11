package puzzle

import (
	"regexp"
	"strconv"
	"strings"
)

func GetWinnersAndMyNumbers(text string) ([]int, []int) {
	numbers := strings.Split(text, ": ")[1]
	split := strings.Split(numbers, " | ")
	winners, myNumbers := split[0], split[1]
	return NumberStringToSlice(winners), NumberStringToSlice(myNumbers)
}

func NumberStringToSlice(text string) []int {
	var nums []int
	re := regexp.MustCompile(`\d+`)
	matches := re.FindAllString(text, -1)
	for _, match := range matches {
		asInt, _ := strconv.Atoi(match)
		nums = append(nums, asInt)
	}
	return nums
}

func MatchCount(winners []int, myNumbers []int) int {
	count := 0
	for _, w := range winners {
		for _, m := range myNumbers {
			if w == m {
				count++
			}
		}
	}
	return count
}
