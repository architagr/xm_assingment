package jsonlogger

import (
	"io"
	"log/slog"
	"xm_logger/logger"
)

type xmJsonLogger struct {
	logger *slog.Logger
}

func (jsonLog *xmJsonLogger) Debug(msg string, args ...any) {
	jsonLog.logger.Debug(msg, args...)
}
func (jsonLog *xmJsonLogger) Info(msg string, args ...any) {
	jsonLog.logger.Info(msg, args...)
}
func (jsonLog *xmJsonLogger) Warn(msg string, args ...any) {
	jsonLog.logger.Warn(msg, args...)
}
func (jsonLog *xmJsonLogger) Error(msg string, args ...any) {
	jsonLog.logger.Error(msg, args...)
}

func InitXmJsonLogger(w io.Writer, opts *slog.HandlerOptions, serviceName string) logger.IXmLogger {
	logger := slog.New(slog.NewJSONHandler(w, opts))
	return &xmJsonLogger{
		logger: logger.With("serviceName", serviceName),
	}
}
