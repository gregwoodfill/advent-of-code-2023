package puzzle

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIntFromText(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected int
	}{
		{name: "no digits", input: "hello", expected: 0},
		{name: "single digit", input: "1", expected: 11},
		{name: "digits at start and end", input: "1and2and3", expected: 13},
		{name: "digits in middle", input: "he110", expected: 10},
		{name: "empty string", input: "", expected: 0},
		{name: "leading 0", input: "0987654323d1", expected: 1},
		{name: "with digits", input: "0nine87654323done", expected: 1},
		{name: "with digits non-zero", input: "onenine87654323done", expected: 11},
		{name: "with digits next to each other", input: "twone", expected: 21},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got, err := FirstAndLastDigitAsInt(tc.input)
			assert.Equal(t, tc.expected, got)
			assert.Nil(t, err)
		})
	}
}
