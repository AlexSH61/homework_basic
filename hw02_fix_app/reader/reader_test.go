package reader

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestReadJson(t *testing.T) {
	testCases := []struct {
		json     string
		expected struct {
			UserID       int
			Age          int
			Name         string
			DepartmentID int
		}
	}{
		{
			json: `{"UserID": 0, "Age": 25, "Name": "Rob", "DepartmentID": 0}`,
			expected: struct {
				UserID       int
				Age          int
				Name         string
				DepartmentID int
			}{
				UserID:       0,
				Age:          25,
				Name:         "Rob",
				DepartmentID: 0,
			},
		},
		{
			json: `{"UserID": 0, "Age": 30, "Name": "George", "DepartmentID": 0}`,
			expected: struct {
				UserID       int
				Age          int
				Name         string
				DepartmentID int
			}{
				UserID:       0,
				Age:          30,
				Name:         "George",
				DepartmentID: 0,
			},
		},
	}

	for _, tc := range testCases {
		var result struct {
			UserID       int
			Age          int
			Name         string
			DepartmentID int
		}

		err := json.Unmarshal([]byte(tc.json), &result)
		require.NoError(t, err, "Error parsing JSON")

		assert.Equal(t, tc.expected.UserID, result.UserID)
		assert.Equal(t, tc.expected.Age, result.Age)
		assert.Equal(t, tc.expected.Name, result.Name)
		assert.Equal(t, tc.expected.DepartmentID, result.DepartmentID)

	}
}
