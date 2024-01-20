package utillog

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"strings"
)

type LogStat struct {
	InfoCount  int
	WarnCount  int
	ErrorCount int
}

func AnalyzeLogFile(reader io.Reader, logLevel string) (*LogStat, error) {
	stats := &LogStat{}
	scanner := bufio.NewScanner(reader)

	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, logLevel) {
			switch logLevel {
			case "INFO":
				stats.InfoCount++
			case "WARN":
				stats.WarnCount++
			case "ERROR":
				stats.ErrorCount++
			}
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return stats, nil
}

func OutputStatistics(stats *LogStat, output io.Writer) error {
	log.SetOutput(output)
	fmt.Fprintln(output, "Log Statistics:")
	fmt.Fprintf(output, "Info: %d\n", stats.InfoCount)
	fmt.Fprintf(output, "Warning: %d\n", stats.WarnCount)
	fmt.Fprintf(output, "Error: %d\n", stats.ErrorCount)
	return nil
}
