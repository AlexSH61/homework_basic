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

	filePFlag := pflag.StringP("file", "f", "", "путь к анализируемому лог-файлу")
	levePFlag := pflag.StringP("level", "l", "", "уровень логов")
	outputPFlag := pflag.StringP("output", "o", "", "путь к файлу для записи")
	pflag.Parse()
	fileEnv, levelEnv, outputEnv := readenv.ReadEnv()
	if *filePFlag == "" && fileEnv != "" {
		*filePFlag = fileEnv
	}
	if *levePFlag == "" && levelEnv != "" {
		*levePFlag = levelEnv
	}
	if *outputPFlag == "" && outputEnv != "" {
		*outputPFlag = outputEnv
	}
	file, err := os.Open(*filePFlag)
	if err != nil {
		log.Fatalf("Ошибка при открытии лог-файла: %v", err)
	}
	if err != nil {
		log.Fatalf("Ошибка при анализе лог-файла: %v", err)
	}

	defer func() {
		if err = file.Close(); err != nil {
			log.Printf("Ошибка при закрытии лог-файла: %v", err)
		}
	}()
	stats, err := utillog.AnalyzeLogFile(file, *levePFlag)
	log.Println("Статистика по умолчанию:")
	err = utillog.OutputStatistics(stats, *outputPFlag)
	if err != nil {
		fmt.Printf("Ошибка при выводе статистики: %v", err)
	}
	log.Printf("Инфо: %d\n", defaultStats.InfoCount)
	log.Printf("Предупреждения: %d\n", defaultStats.WarnCount)
	log.Printf("Ошибки: %d\n", defaultStats.ErrorCount)
}
