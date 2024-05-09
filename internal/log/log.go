package log

import (
	"context"
	"log/slog"
	"os"
)

type Logger struct {
	logger *slog.Logger
}

func NewLogger() *Logger {
	return &Logger{logger: slog.New(slog.NewJSONHandler(os.Stdout, nil))}
}

func (l *Logger) InfoContext(ctx context.Context, msg string, args ...any) {
	l.logger.InfoContext(ctx, msg, args...)
}
