package reader

import (
	"testing"

	"github.com/AlexSH61/homework_basic/hw02_fix_app/types"
	"github.com/stretchr/testify/assert"
)

func TestReadJSON(t *testing.T) {
	testCases := []struct {
		name     string
		filePath string
		expected []types.Employee
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
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result, err := ReadJSON(tc.filePath)
			assert.NoError(t, err, "Expected no error")
			assert.Equal(t, tc.expected, result, "Expected values to be equal")
		})
	}
}
