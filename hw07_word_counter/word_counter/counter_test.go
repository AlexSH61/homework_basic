package counter_test

import (
	"errors"
	"testing"

	wordCounter "github.com/AlexSH61/homework_basic/hw07_word_counter/word_counter"
	"github.com/stretchr/testify/assert"
)

func TestCountWords(t *testing.T) {
	testCase := []struct {
		name           string
		preparedString []string
		expected       map[string]int
		expectedError  error
	}{
		{
			name: "valid case",
			preparedString: []string{
				"everything",
				"will",
				"be",
				"alright",
				"in",
				"the",
				"end",
				"alright",
			},
			expected: map[string]int{
				"everything": 1,
				"will":       1,
				"be":         1,
				"alright":    2,
				"in":         1,
				"the":        1,
				"end":        1,
			},
			expectedError: nil,
		},
		{
			name: "valid case",
			preparedString: []string{
				"the",
				"night",
				"is",
				"darkest",
				"just",
				"before",
				"the",
				"dawn",
			},
			expected: map[string]int{
				"the":     2,
				"night":   1,
				"is":      1,
				"darkest": 1,
				"just":    1,
				"before":  1,
				"dawn":    1,
			},
			expectedError: nil,
		},
		{
			preparedString: []string{},
			expected:       map[string]int{},
			expectedError:  errors.New("incorrect data"),
		},
	}
	for _, test := range testCase {
		t.Run(test.name, func(t *testing.T) {
			result, err := wordCounter.CountWords(test.preparedString)
			if test.expectedError != nil {
				assert.EqualError(t, err, test.expectedError.Error())
				return
			}
			assert.NoError(t, err, "unknown")
			assert.Equal(t, test.expected, result, "")
		})
	}
}
