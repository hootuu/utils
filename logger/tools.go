package logger

import (
	"go.uber.org/zap"
	"io"
)

type slogWriter struct {
	logger *zap.Logger
}

func (s *slogWriter) Write(p []byte) (int, error) {
	n := len(p)
	if n == 0 {
		return 0, nil
	}
	s.logger.Info(string(p))
	return n, nil
}

func GetLoggerWriter(logger *zap.Logger) io.Writer {
	return &slogWriter{logger: logger}
}
