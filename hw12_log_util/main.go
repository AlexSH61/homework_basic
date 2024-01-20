package main

import (
	"fmt"
	"log"
	"os"

	"github.com/AlexSH61/homework_basic/hw12_log_util/readenv"
	"github.com/AlexSH61/homework_basic/hw12_log_util/utillog"
	"github.com/spf13/pflag"
)

func main() {
	defaultStats := &utillog.LogStat{
		InfoCount:  5,
		WarnCount:  5,
		ErrorCount: 1,
	}

	filePFlag := pflag.StringP("file", "f", "", "path to the log file to analyze")
	levelPFlag := pflag.StringP("level", "l", "", "log level")
	outputPFlag := pflag.StringP("output", "o", "", "path to the file for writing")
	pflag.Parse()
	env := readenv.ReadEnv()
	if *filePFlag == "" && env.File != "" {
		*filePFlag = env.File
	}
	if *levelPFlag == "" && env.Level != "" {
		*levelPFlag = env.Level
	}
	if *outputPFlag == "" && env.Output != "" {
		*outputPFlag = env.Output
	}
	file, errOpen := os.Open(*filePFlag)
	if errOpen != nil {
		log.Printf("Error opening the log file: %v", errOpen)
	}
	defer func() {
		if errOpen = file.Close(); errOpen != nil {
			log.Printf("Error closing the log file: %v", errOpen)
		}
	}()

	stats, err := utillog.AnalyzeLogFile(file, *levelPFlag)
	if err != nil {
		log.Printf("Error analyzing the log file: %v", err)
	}

	outputWriter := os.Stdout
	if *outputPFlag != "" {
		fileOutput, errOutput := os.Create(*outputPFlag)
		if errOutput != nil {
			log.Printf("Error creating the file: %v", errOutput)
		}
		defer fileOutput.Close()
		outputWriter = fileOutput
	}

	log.Println("Default statistics:")
	err = utillog.OutputStatistics(stats, outputWriter)
	if err != nil {
		fmt.Printf("Error outputting statistics: %v", err)
	}

	log.Printf("Info: %d\n", defaultStats.InfoCount)
	log.Printf("Warnings: %d\n", defaultStats.WarnCount)
	log.Printf("Errors: %d\n", defaultStats.ErrorCount)
}
