package utillog

import (
	"bufio"
	"io"
	"log"
	"os"
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

func OutputStatistics(stats *LogStat, outputFilePath string) error {
	var output io.Writer

	if outputFilePath != "" {
		file, err := os.Create(outputFilePath)
		if err != nil {
			return err
		}
		defer file.Close()
		output = bufio.NewWriter(file)
	} else {
		output = os.Stdout
	}
	log.SetOutput(output)
	log.Println("Log Statistics:")
	log.Printf("Info: %d\n", stats.InfoCount)
	log.Printf("Warning: %d\n", stats.WarnCount)
	log.Printf("Error: %d\n", stats.ErrorCount)

	return nil
}
