package readenv

import "os"

type LogAnalyzerEnv struct {
	File   string
	Level  string
	Output string
}

func ReadEnv() *LogAnalyzerEnv {
	return &LogAnalyzerEnv{
		File:   os.Getenv("LOG_ANALYZER_FILE"),
		Level:  os.Getenv("LOG_ANALYZER_LEVEL"),
		Output: os.Getenv("LOG_ANALYZER_OUTPUT"),
	}
}
