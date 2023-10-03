package prepare_test

import (
	"errors"
	"testing"

	prepare "github.com/AlexSH61/homework_basic/hw07_word_counter/prepare_string"
	"github.com/stretchr/testify/assert"
)

func TestPrepare(t *testing.T) {
	testCase := []struct {
		name          string
		inpuText      string
		expected      []string
		expectedError error
	}{
		{
			name:     "valid case",
			inpuText: "Everything will be alright in the end. If it's not alright, it's not the end.",
			expected: []string{
				"everything", "will", "be", "alright", "in", "the", "end", "if", "its",
				"not", "alright", "its", "not", "the", "end",
			},
			expectedError: nil,
		},
		{
			name:          "invalid case",
			inpuText:      "",
			expected:      []string{},
			expectedError: errors.New("incorrect text"),
		},
	}
	for _, test := range testCase {
		t.Run(test.name, func(t *testing.T) {
			result, err := prepare.StringPrepare(test.inpuText)
			if test.expectedError != nil {
				assert.EqualError(t, err, test.expectedError.Error())
			} else {
				assert.NoError(t, err)
				assert.Equal(t, test.expected, result)
			}
		})
	}
}
