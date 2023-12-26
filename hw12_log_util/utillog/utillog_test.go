package utillog

import (
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAnalyzeLogFile(t *testing.T) {
	content := `
		INFO Some information
		WARN Warning message
		ERROR An error occurred
		INFO Another info
	`
	tmpfile, err := os.CreateTemp("", "testlog_*.txt")
	if err != nil {
		t.Fatalf("Error creating temp log file: %v", err)
	}
	defer os.Remove(tmpfile.Name())
	defer tmpfile.Close()

	_, err = tmpfile.WriteString(strings.TrimSpace(content))
	if err != nil {
		t.Fatalf("Error writing to temp log file: %v", err)
	}

	tests := []struct {
		name     string
		logLevel string
		expected *LogStat
	}{
		{"TestInfoLogLevel", "INFO", &LogStat{InfoCount: 2, WarnCount: 0, ErrorCount: 0}},
		{"TestWarnLogLevel", "WARN", &LogStat{InfoCount: 0, WarnCount: 1, ErrorCount: 0}},
		{"TestErrorLogLevel", "ERROR", &LogStat{InfoCount: 0, WarnCount: 0, ErrorCount: 1}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, errAnlzr := AnalyzeLogFile(tmpfile.Name(), tt.logLevel)
			if err != nil {
				t.Fatalf("Error analyzing log file: %v", errAnlzr)
			}
			fmt.Printf("Actual Stats: %+v\n", result)
			assert.Equal(t, tt.expected.InfoCount, result.InfoCount, "InfoCount mismatch")
			assert.Equal(t, tt.expected.WarnCount, result.WarnCount, "WarnCount mismatch")
			assert.Equal(t, tt.expected.ErrorCount, result.ErrorCount, "ErrorCount mismatch")
		})
	}
}
