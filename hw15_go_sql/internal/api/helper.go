package api

import "go.uber.org/zap"

func (a *APIServer) configLogger() error {
	logLevel, err := zap.ParseAtomicLevel(a.config.LoggerLevel)
	if err != nil {
		return err
	}
	cfg := zap.NewProductionConfig()
	cfg.Level = logLevel
	logger, err := cfg.Build()
	if err != nil {
		return err
	}
	a.logger = logger.Sugar()
	return nil
}
