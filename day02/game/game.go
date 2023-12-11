package game

import (
	"log"
	"regexp"
	"strconv"
	"strings"
)

type (
	Game struct {
		Id     int
		Rounds []map[string]int
	}
)

var gameIdRe = regexp.MustCompile(`Game (\d+):`)
var colorCountRe = regexp.MustCompile(`(\d+) (blue|red|green)`)

func NewGame(summary string) Game {
	rounds := []map[string]int{}
	idStr := gameIdRe.FindStringSubmatch(summary)[1]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		log.Fatal(err)
	}
	roundStrings := strings.Split(summary, ": ")[1]
	roundDetails := strings.Split(roundStrings, "; ")
	for _, roundDetail := range roundDetails {
		round := map[string]int{
			"red":   0,
			"green": 0,
			"blue":  0,
		}
		for _, colorsAndCounts := range strings.Split(roundDetail, ", ") {
			countAndColor := colorCountRe.FindStringSubmatch(colorsAndCounts)
			count, _ := strconv.Atoi(countAndColor[1])
			color := countAndColor[2]
			round[color] = count
		}
		rounds = append(rounds, round)
	}
	return Game{Id: id, Rounds: rounds}
}

func (g Game) Possible(blue int, green int, red int) bool {
	canPlay := true
	for _, round := range g.Rounds {
		if round["blue"] > blue || round["green"] > green || round["red"] > red {
			canPlay = false
		}
	}
	return canPlay
}

func (g Game) MiniumRequired() map[string]int {
	minRequired := map[string]int{
		"red":   0,
		"green": 0,
		"blue":  0,
	}
	for _, round := range g.Rounds {
		for color := range round {
			if minRequired[color] < round[color] {
				minRequired[color] = round[color]
			}
		}
	}
	return minRequired
}
