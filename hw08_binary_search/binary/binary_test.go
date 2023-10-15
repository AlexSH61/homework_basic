package binary_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	binaryAlg "github.com/AlexSH61/homework_basic/binary_search/binary"
)

func TestBinary(t *testing.T) {
	testCase := []struct {
		name         string
		inputArray   []int
		intputNumber int
		expected     int
	}{
		{
			name:         "sample",
			inputArray:   []int{2, 3, 4, 5, 6, 7, 8, 89},
			intputNumber: 5,
			expected:     3,
		},
		{
			name:         "sample_unfounded",
			inputArray:   []int{2, 3, 4, 5, 6, 7, 8, 89},
			intputNumber: 9,
			expected:     -1,
		},
	}
	for _, tc := range testCase {
		t.Run(tc.name, func(t *testing.T) {
			result := binaryAlg.BinarySearching(tc.intputNumber, tc.inputArray)
			assert.Equal(t, tc.expected, result)
		})
	}
}
