package logger

import (
	slogmulti "github.com/samber/slog-multi"
	"gopkg.in/natefinch/lumberjack.v2"

	"log/slog"
	"os"
)

var logger *slog.Logger

func Set() {
	logger = slog.New(
		slog.NewTextHandler(os.Stderr, nil),
	)
	logRotate := &lumberjack.Logger{
		Filename:   "pkg/logger/app.log",
		MaxSize:    10, // MB
		MaxBackups: 3,
		MaxAge:     1, // days
		Compress:   true,
	}

	logger = slog.New(
		slogmulti.Fanout(
			slog.NewJSONHandler(logRotate, nil),
			slog.NewTextHandler(os.Stderr, nil),
		),
	)

	slog.SetDefault(logger)
}
