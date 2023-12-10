package puzzle

import (
	"fmt"
	"strconv"
	"strings"
)

var digitChars = "0123456789"
var digits = map[string]string{
	"zero":  "0",
	"one":   "1",
	"two":   "2",
	"three": "3",
	"four":  "4",
	"five":  "5",
	"six":   "6",
	"seven": "7",
	"eight": "8",
	"nine":  "9",
}

// FirstAndLastDigitAsInt returns finds all digit chars in text and returns an int
// that combines all the digit chars, e.g. "9foo8bar0" -> 980
// return 0 if none found
func FirstAndLastDigitAsInt(text string) (int, error) {
	text = replaceDigitNames(text)
	firstIndex := strings.IndexAny(text, digitChars)
	lastIndex := strings.LastIndexAny(text, digitChars)
	var firstDigit int
	var secondDigit int
	var err error

	if firstIndex == -1 {
		return 0, nil
	}
	firstDigit, err = strconv.Atoi(string(text[firstIndex]))
	if err != nil {
		return -1, err
	}
	secondDigit, err = strconv.Atoi(string(text[lastIndex]))
	if err != nil {
		return -1, err
	}

	combined := fmt.Sprintf("%d%d", firstDigit, secondDigit)
	return strconv.Atoi(combined)
}

// replaceDigitNames returns a string with digit names with the digit char
// e.g. "nineone" to "91"
func replaceDigitNames(text string) string {
	for hasDigits(text) {
		text = replaceFirstDigit(text)
	}
	return text
}

func replaceFirstDigit(text string) string {
	firstIdx := -1
	digitName := ""
	for k := range digits {
		idx := strings.Index(text, k)
		if idx == -1 {
			continue
		}
		if firstIdx == -1 || idx < firstIdx {
			firstIdx = idx
			digitName = k
		}
	}
	if firstIdx == -1 {
		return text
	}
	// add the last char back in so string like
	// twone will return 21e and not 2ne
	lastCharInName := string(digitName[len(digitName)-1])
	replaceText := fmt.Sprintf("%s%s", digits[digitName], lastCharInName)
	text = strings.Replace(text, digitName, replaceText, 1)
	return text
}

func hasDigits(text string) bool {
	for k := range digits {
		if strings.Contains(text, k) {
			return true
		}
	}
	return false
}
