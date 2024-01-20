package utillog

import (
	"bytes"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAnalyzeLogFile(t *testing.T) {
	logs := `INFO: This is an info message
    WARNING: This is a warning message
    ERROR: This is an error message
    ERROR: Another error message`

	reader := strings.NewReader(logs)
	expectedStats := &LogStat{
		InfoCount:  0,
		WarnCount:  0,
		ErrorCount: 2,
	}

	stats, err := AnalyzeLogFile(reader, "ERROR")
	assert.NoError(t, err, "Unexpected error")

	assert.Equal(t, expectedStats, stats, "Stats do not match expected values")
}

func TestOutputStatistics(t *testing.T) {
	stats := &LogStat{
		InfoCount:  2,
		WarnCount:  3,
		ErrorCount: 1,
	}

	var buffer bytes.Buffer
	err := OutputStatistics(stats, &buffer)
	assert.NoError(t, err, "Unexpected error")

	expectedOutput := fmt.Sprintf(
		"Log Statistics:\nInfo: %d\nWarning: %d\nError: %d\n",
		stats.InfoCount, stats.WarnCount, stats.ErrorCount,
	)

	actualOutput := buffer.String()

	assert.Equal(t, expectedOutput, actualOutput, "Output does not match expected values")
}
