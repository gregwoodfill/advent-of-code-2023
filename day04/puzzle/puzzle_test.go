package puzzle

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNumberStringToSlice(t *testing.T) {
	input := "1 51 41"
	want := []int{1, 51, 41}
	got := NumberStringToSlice(input)

	assert.Equal(t, want, got)
}

func TestGetWinnersAndMyNumbers(t *testing.T) {
	input := "Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53"
	expectedWinners := []int{41, 48, 83, 86, 17}
	expectedMyNumbers := []int{83, 86, 6, 31, 17, 9, 48, 53}
	winners, myNumbers := GetWinnersAndMyNumbers(input)

	assert.Equal(t, expectedWinners, winners)
	assert.Equal(t, expectedMyNumbers, myNumbers)
}

func TestMatchCount(t *testing.T) {
	winners := []int{41, 48, 83, 86, 17}
	myNumbers := []int{83, 86, 6, 31, 17, 9, 48, 53}
	want := 4
	got := MatchCount(winners, myNumbers)
	assert.Equal(t, want, got)
}
