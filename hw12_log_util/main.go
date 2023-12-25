package main

import (
	"log"

	"github.com/spf13/pflag"

	"github.com/AlexSH61/homework_basic/hw12_log_util/utillog/utillog"
)

func main() {
	fileFlag := pflag.StringP("file", "f", "", "путь к анализируемому лог-файлу")
	leveFLag := pflag.StringP("level", "l", "info", "уровень логов")
	outputFlag := pflag.StringP("output", "o", "", "путь к файлу для записи")
	pflag.Parse()
	if *fileFlag == "" {
		log.Fatal("Error: Log file path is not provided.")
	}

	stats, err := utillog.AnalyzeLogFile(*fileFlag, *leveFLag)
	if err != nil {
		log.Fatalf("Error analyzing log file: %v", err)
	}

	err = utillog.OutputStatistics(stats, *outputFlag)
	if err != nil {
		log.Fatalf("Error outputting statistics: %v", err)
	}
}
