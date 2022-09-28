package logger

import "go.uber.org/zap"

type Logger interface {
	Infof(msg string, args ...any)
	Fatalf(msg string, args ...any)
	Warnf(msg string, args ...any)
	Errorf(msg string, args ...any)
	Panicf(msg string, args ...any)
	Debugf(msg string, args ...any)
}

type log struct {
	Logger
}

func New() *log {
	zapLogger, _ := zap.NewProduction()
	defer zapLogger.Sync()
	sugar := zapLogger.Sugar()

	return &log{
		Logger: sugar,
	}
}

func (l log) Info(msg string, args ...any) {
	l.Infof(msg, args)
}
