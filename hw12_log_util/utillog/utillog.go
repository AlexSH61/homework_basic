package utillog

import (
	"bufio"
	"log"
	"os"
	"strings"
)

type LogStat struct {
	InfoCount  int
	WarnCount  int
	ErrorCount int
}

func AnalyzeLogFile(filePath string, logLevel string) (*LogStat, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	stats := &LogStat{}
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, logLevel) {
			switch {
			case strings.Contains(line, "INFO"):
				stats.InfoCount++
			case strings.Contains(line, "WARN"):
				stats.WarnCount++
			case strings.Contains(line, "ERROR"):
				stats.ErrorCount++
			}
		}
	}

	if err = scanner.Err(); err != nil {
		return nil, err
	}

	return stats, nil
}

func OutputStatistics(stats *LogStat, outputPath string) error {
	var outputWriter *os.File

	if outputPath != "" {
		outputFile, err := os.OpenFile(outputPath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0o666)
		if err != nil {
			return err
		}
		defer outputFile.Close()
		outputWriter = outputFile
	} else {
		outputWriter = os.Stdout
	}

	log.New(outputWriter, "", 0).Println("Log Statistics:")
	log.New(outputWriter, "", 0).Printf("Info: %d\n", stats.InfoCount)
	log.New(outputWriter, "", 0).Printf("Warning: %d\n", stats.WarnCount)
	log.New(outputWriter, "", 0).Printf("Error: %d\n", stats.ErrorCount)

	return nil
}
