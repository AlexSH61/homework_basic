package logger

import "go.uber.org/zap"

func Newlogger() (*zap.SugaredLogger, error) {
	LogConfig := zap.NewProductionConfig()
	log, err := LogConfig.Build()
	if err != nil {
		return nil, err
	}
	return log.Sugar(), nil
}
