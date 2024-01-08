package readenv

import "os"

func ReadEnv() (string, string, string) {
	FileEnv := os.Getenv("LOG_ANALYZER_FILE")
	LevelEnv := os.Getenv("LOG_ANALYZER_LEVEL")
	OutputEnv := os.Getenv("LOG_ANALYZER_OUTPUT")
	return FileEnv, LevelEnv, OutputEnv
}
