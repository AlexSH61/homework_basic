package counter_test

import (
	"testing"

	wordCounter "github.com/AlexSH61/homework_basic/hw07_word_counter/word_counter"
	"github.com/stretchr/testify/assert"
)

func TestCountWords(t *testing.T) {
	testCase := []struct {
		name        string
		inputString string
		expected    map[string]int
	}{
		{
			name:        "valid case",
			inputString: "The night is darkest just before the dawn. And I promise you, the dawn is coming.",
			expected: map[string]int{
				"and":     1,
				"before":  1,
				"coming":  1,
				"darkest": 1,
				"dawn":    2,
				"i":       1,
				"is":      2,
				"just":    1,
				"night":   1,
				"promise": 1,
				"the":     3,
				"you":     1,
			},
		},
		{
			name:        "invalid case",
			inputString: "",
			expected:    map[string]int{},
		},
	}
	for _, test := range testCase {
		t.Run(test.name, func(t *testing.T) {
			result := wordCounter.CountWords(test.inputString)
			assert.Equal(t, test.expected, result, test.name)
		})
	}
}
