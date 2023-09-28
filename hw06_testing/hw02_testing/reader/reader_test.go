package reader

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/AlexSH61/homework_basic/hw06_testing/hw02_testing/types"
)

func TestReadJSON(t *testing.T) {
	testCases := []struct {
		name        string
		filePath    string
		expected    []types.Employee
		expectError error
	}{
		{
			name:     "Valid JSON",
			filePath: "data.json",
			expected: []types.Employee{
				{
					UserID:       10,
					Age:          25,
					Name:         "Rob",
					DepartmentID: 3,
				},
				{
					UserID:       11,
					Age:          30,
					Name:         "George",
					DepartmentID: 2,
				},
			},
		}, {
			name:        "Invalid JSON",
			filePath:    "/Users/aleksandr/homework_basic/hw06_testing/hw02_testing/reader/invalid.json",
			expected:    nil,
			expectError: assert.AnError,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result, err := ReadJSON(tc.filePath)
			if tc.expectError == assert.AnError {
				assert.Error(t, err, "Expected an error")
				assert.Nil(t, result, "")
			} else {
				assert.NoError(t, err, "Expected no error")
				assert.Equal(t, tc.expected, result, "Expected values to be equal")
			}
		})
	}
}
